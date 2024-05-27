package huifupay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/yuewokeji/huifupay/notify"
	"github.com/yuewokeji/huifupay/request"
	"github.com/yuewokeji/huifupay/response"
	"github.com/yuewokeji/huifupay/sign"
	"golang.org/x/net/context/ctxhttp"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func NewClient(config Config, options ...ClientOption) *Client {
	c := &Client{
		config: config,
	}
	c.SetLogger("Info", os.Stdout)
	c.initWithOptions(options...)
	c.initSigner()

	return c
}

func (c *Client) initWithOptions(options ...ClientOption) {
	for _, o := range options {
		o(c)
	}
}

func (c *Client) initSigner() {
	c.signer = sign.NewRSASigner(c.config.RsaHuifuPublicKey, c.config.RsaMerchantPrivateKey)
}

type Client struct {
	config Config

	*Logger
	signer sign.Signer

	httpClientFunc HTTPClientFunc
}

func (c *Client) SetLogger(level string, out io.Writer) {
	if level == "" {
		level = "Info"
	}

	logger := log.New(out, "[HuiFuPay]["+strings.ToUpper(level)+"] ", log.Flags())
	c.Logger = &Logger{
		Logger: logger,
		enable: true,
	}
}

func (c *Client) DoRequest(ctx context.Context, req request.Request, resp response.Response) (err error) {
	httpReq, err := c.BuildRequestWithSign(req)
	if err != nil {
		return errors.Wrap(err, "build request")
	}
	err = c.doHttpRequest(ctx, httpReq, resp, c.config.VerifySign)
	return
}

func (c *Client) DoUploadFile(ctx context.Context, filePath string, req *request.FileRequest, resp response.Response) (err error) {
	req.SetSysID(c.config.SystemID)
	req.SetProductID(c.config.ProductId)

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	req.SetFile(file)
	defer file.Close()

	httpReq, err := req.Build()
	if err != nil {
		return errors.Wrap(err, "build file request")
	}
	go func() {
		err = req.StartWriteToBody()
		if err != nil {
			c.Println("copy file: " + err.Error())
		}
	}()
	err = c.doHttpRequest(ctx, httpReq, resp, false)
	return
}

func (c *Client) doHttpRequest(ctx context.Context, httpReq *http.Request, resp response.Response, verifySign bool) (err error) {
	lf := defaultLogFields()
	defer func() {
		c.requestLog(lf, err)
	}()
	requestToLog(lf, httpReq)

	debug.Printf(">> Request: %s %s %s %s", httpReq.Proto, httpReq.Method, httpReq.URL.RequestURI(), httpReq.URL.Host)
	debug.Printf(">> Request Headers: ")
	for k, v := range httpReq.Header {
		debug.Printf(">> %s: %s", k, strings.Join(v, " "))
	}
	if debug.IsEnable() {
		buf := &bytes.Buffer{}
		tr := io.TeeReader(httpReq.Body, buf)
		httpReq.Body = io.NopCloser(buf)

		data, err := io.ReadAll(tr)
		if err != nil {
			return err
		}
		debug.Printf(">> Request Body: %s", string(data))
	}

	startTime := time.Now()
	rawResp, err := ctxhttp.Do(ctx, c.getHttpClient(httpReq), httpReq)
	lf["{cost}"] = fmt.Sprintf("%s", time.Now().Sub(startTime))
	if err != nil {
		return errors.Wrap(err, "do http request")
	}

	debug.Printf("<< Response: %s %d", rawResp.Proto, rawResp.StatusCode)
	debug.Printf("<< Response Headers: ")
	for k, v := range rawResp.Header {
		debug.Printf("<< %s: %s", k, strings.Join(v, " "))
	}

	lf["{status}"] = strconv.Itoa(rawResp.StatusCode)
	if rawResp.StatusCode != 200 {
		err = httpError(rawResp.StatusCode)
		return err
	}

	err = resp.ParseFromHttpRequest(rawResp)
	if err != nil {
		return errors.Wrap(err, "parse http request")
	}
	debug.Printf("<< Response Body: %s", resp.GetHttpContentString())

	if verifySign {
		err = c.verifySign(resp.GetHttpContent())
		if err != nil {
			return err
		}
	}

	err = response.UnmarshalAutoAssign(resp, resp.GetHttpContent())
	if err != nil {
		return errors.Wrap(err, "unmarshal body")
	}
	return
}

func (c *Client) getHttpClient(req *http.Request) *http.Client {
	if nil != c.httpClientFunc {
		return c.httpClientFunc(req)
	} else if globalHTTPClientFunc != nil {
		return globalHTTPClientFunc(req)
	}
	return http.DefaultClient
}

// DoNotifyRequest 异步回调验证
// https://paas.huifu.com/partners/jiekouguifan#/ybxx/jiekouguifan_ybxx
func (c *Client) DoNotifyRequest(ctx context.Context, req *http.Request, n notify.Notify) (err error) {
	debug.Printf(">> Verify Request: %s %s %s %s", req.Proto, req.Method, req.URL.RequestURI(), req.URL.Host)
	defer func() {
		if err != nil {
			debug.Printf(">> Verify Request error: %s ", err.Error())
		}
	}()

	err = n.ParseFromHttpRequest(req)
	if err != nil {
		return errors.Wrap(err, "parse http request")
	}
	debug.Printf(">> Verify Request Body: %s ", n.GetBodyContentString())

	value, err := url.ParseQuery(n.GetBodyContentString())
	if err != nil {
		return errors.Wrap(err, "parse url query")
	}
	if debug.IsEnable() {
		debug.Printf(">> Verify Request Values: %s ", value.Encode())
	}

	// 支付交易类接口异步通知返回的参数名为 resp_data
	// 商户进件配置类接口的异步通知返回参数名为 data
	data := value.Get("resp_data")
	if data == "" {
		data = value.Get("data")
	}
	sign := value.Get("sign")
	debug.Printf(">> Verify Sign: %s %s ", data, sign)

	plainText := []byte(data)
	err = c.signer.Verify(plainText, []byte(sign))
	if err != nil {
		return err
	}

	return response.UnmarshalAutoAssign(n, plainText)
}

// 验证签名
func (c *Client) verifySign(body []byte) (err error) {
	s := &struct {
		Data json.RawMessage `json:"data"`
		Sign string          `json:"sign"`
	}{}

	err = response.Unmarshal(s, body)
	if err != nil {
		return err
	}
	debug.Printf(">> Verify Sign: %s %s ", string(s.Data), s.Sign)

	err = c.signer.Verify(s.Data, []byte(s.Sign))
	return
}

func (c *Client) BuildRequestWithSign(req request.Request) (httpReq *http.Request, err error) {
	req.SetSysID(c.config.SystemID)
	req.SetProductID(c.config.ProductId)

	text, err := request.MarshalAutoAssign(req.GetData())
	if err != nil {
		return nil, errors.Wrap(err, "marshal `data`")
	}

	signature, err := c.signer.Encrypt(text)
	if err != nil {
		return nil, errors.Wrap(err, "encrypt request")
	}

	req.SetSign(string(signature))
	httpReq, err = req.Build()
	return
}

func (c *Client) GetSigner() sign.Signer {
	return c.signer
}

type ClientOption func(client *Client)

// WithHTTPClientFunc Client会优先使用WithHttpClientFunc()，如果没有设置再从SetGlobalHttpClientFunc()中获取
func WithHTTPClientFunc(fn HTTPClientFunc) ClientOption {
	return func(client *Client) {
		client.httpClientFunc = fn
	}
}
