package huifupay

import (
	"fmt"
	"net/http"
)

// HTTPClientFunc 用来自定义http.Client
type HTTPClientFunc func(r *http.Request) *http.Client

var globalHTTPClientFunc HTTPClientFunc

func SetGlobalHTTPClientFunc(fn HTTPClientFunc) {
	globalHTTPClientFunc = fn
}

// https://paas.huifu.com/partners/jiekouguifan#/api_bzzd?id=%e8%bf%94%e5%9b%9e%e7%a0%81
// 最近更新时间：2022.7.11
var httpStatusErrorText = map[int]string{
	404: "服务不存在",
	412: "请求体包含非法字符",
	500: "后端服务器错误",
	502: "后端服务调用超时或连接失败",
	512: "API网关SHUTDOWN",
	911: "按ip、商户id、api_key等维度限流的返回码",
	914: "按接口总的访问数限流拒绝返回码",
	912: "安全封禁返回码",
	921: "网关鉴权不通过",
	922: "验签不通过",
	923: "返回结果加签失败",
}

func httpError(code int) error {
	str := "Unknown error!"
	if v, ok := httpStatusErrorText[code]; ok {
		str = v
	}
	return fmt.Errorf("http response code: %d, desc: %s", code, str)
}
