package huifupay

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// example:
// [HuiFuPay][INFO] 2024/04/23 17:54:13 "POST /v2/trade/payment/scanpay/query HTTP/1.1" 200 "api.huifu.com" "240.121096ms" "-"
var logTemplate = `"{method} {uri} {proto}" {status} "{host}" "{cost}" "{error}"`
var logFields = []string{"{uri}", "{method}", "{proto}", "{host}", "{cost}", "{status}", "{error}"}

type Logger struct {
	*log.Logger
	enable bool
}

func (l *Logger) Enable() {
	l.enable = true
}

func (l *Logger) Disable() {
	l.enable = false
}

func defaultLogFields() map[string]string {
	m := make(map[string]string)
	for _, v := range logFields {
		m[v] = ""
	}
	return m
}
func requestToLog(fields map[string]string, req *http.Request) {
	fields["{uri}"] = req.URL.RequestURI()
	fields["{method}"] = req.Method
	fields["{proto}"] = req.Proto
	fields["{host}"] = req.URL.Host
}

func (c *Client) requestLog(fields map[string]string, err error) {
	if !c.Logger.enable || c.Logger == nil {
		return
	}
	if err != nil {
		fields["{error}"] = fmt.Sprintf(`error: %s`, err.Error())
	} else {
		fields["{error}"] = "-"
	}
	s := logTemplate
	for k, v := range fields {
		s = strings.Replace(s, k, v, -1)
	}
	c.Logger.Output(2, s)
}

func (c *Client) Println(s string) {
	c.Logger.Output(2, s)
}
