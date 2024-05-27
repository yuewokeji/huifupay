package response

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/yuewokeji/huifupay/autoassign"
	"io"
	"net/http"
)

type Response interface {
	ParseFromHttpRequest(r *http.Response) error
	GetHttpResponse() *http.Response
	GetHttpContent() []byte
	GetHttpContentString() string
}

func NewBaseResponse() *BaseResponse {
	return &BaseResponse{}
}

type BaseResponse struct {
	resp        *http.Response
	httpContent []byte
}

func (b *BaseResponse) ParseFromHttpRequest(r *http.Response) (err error) {
	defer r.Body.Close()
	b.httpContent, err = io.ReadAll(r.Body)
	return err
}

func (b *BaseResponse) GetHttpResponse() *http.Response {
	return b.resp
}

func (b *BaseResponse) GetHttpContent() []byte {
	return b.httpContent
}

func (b *BaseResponse) GetHttpContentString() string {
	return string(b.httpContent)
}

type DataHeader struct {
	RespCode string `json:"resp_code"`
	RespDesc string `json:"resp_desc"`
}

type Sign struct {
	Sign string `json:"sign"`
}

func Unmarshal(v interface{}, data []byte) error {
	return json.Unmarshal(data, v)
}

func UnmarshalAutoAssign(v interface{}, data []byte) error {
	err := Unmarshal(v, data)
	if err != nil {
		return err
	}
	err = autoassign.JSONStringToObject(v)
	if err != nil {
		return errors.Wrap(err, "unmarshal auto assign")
	}
	return nil
}
