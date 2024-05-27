package request

import (
	"bytes"
	"encoding/json"
	"github.com/yuewokeji/huifupay/autoassign"
	"github.com/yuewokeji/huifupay/utils"
	"net/http"
)

type Request interface {
	SetURL(string)
	GetURL() string
	SetSysID(string)
	GetSysID() string
	SetProductID(string)
	GetProductID() string
	SetSign(string)
	SetData(interface{})
	GetData() interface{}
	Build() (*http.Request, error)
}

func New(url string, data interface{}) Request {
	req := &BaseRequest{}
	req.SetURL(url)
	req.SetData(data)

	return req
}

type BaseRequest struct {
	url string `json:"-"`

	SysID     string      `json:"sys_id"`
	ProductID string      `json:"product_id"`
	Sign      string      `json:"sign"`
	Data      interface{} `json:"data"`
}

func (b *BaseRequest) SetURL(url string) {
	b.url = url
}

func (b *BaseRequest) GetURL() string {
	return b.url
}

func (b *BaseRequest) SetSysID(sid string) {
	b.SysID = sid
}

func (b *BaseRequest) GetSysID() string {
	return b.SysID
}

func (b *BaseRequest) SetProductID(pid string) {
	b.ProductID = pid
}

func (b *BaseRequest) GetProductID() string {
	return b.ProductID
}

func (b *BaseRequest) SetSign(sign string) {
	b.Sign = sign
}

func (b *BaseRequest) SetData(data interface{}) {
	b.Data = data
}

func (b *BaseRequest) GetData() interface{} {
	return b.Data
}

func (b *BaseRequest) Build() (req *http.Request, err error) {
	body, err := Marshal(b)
	if err != nil {
		return
	}

	// https://paas.huifu.com/partners/jiekouguifan#/README?id=%e5%8d%8f%e8%ae%ae%e8%a7%84%e5%88%99
	// 只有POST接口
	req, err = http.NewRequest("POST", b.url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	// https://paas.huifu.com/partners/jiekouguifan#/api_qqxy?id=header%e6%a8%a1%e5%9e%8b
	// 最近更新时间：2023.10.9
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	return
}

func Marshal(v interface{}) (b []byte, err error) {
	ok := false
	m := make(map[string]interface{})
	if m, ok = v.(map[string]interface{}); !ok {
		m = utils.StructToMapJSONTag(v)
	}

	// https://paas.huifu.com/partners/jiekouguifan#/api_v2jqyq?id=%e5%a6%82%e4%bd%95%e5%8a%a0%e7%ad%be
	// 最近更新时间：2022.9.20
	// 加签前字段排序
	m = utils.SortByKeys(m)
	return json.Marshal(m)
}

func MarshalAutoAssign(v interface{}) (b []byte, err error) {
	err = autoassign.ObjectToJSONString(v)
	if err != nil {
		return
	}
	return Marshal(v)
}
