package v2

import "github.com/yuewokeji/huifupay/request"

func newRequest(api string, v interface{}) request.Request {
	return request.New(requestURL(api), v)
}

func requestURL(path string) string {
	const domain = "https://api.huifu.com"
	return domain + path
}
