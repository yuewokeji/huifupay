package notify

import (
	"io"
	"net/http"
)

type Notify interface {
	ParseFromHttpRequest(r *http.Request) error
	GetHttpRequest() *http.Request
	GetBodyContent() []byte
	GetBodyContentString() string
}

func NewBaseNotify() *BaseNotify {
	return &BaseNotify{}
}

type BaseNotify struct {
	req             *http.Request
	httpBodyContent []byte
}

func (b *BaseNotify) ParseFromHttpRequest(r *http.Request) (err error) {
	defer r.Body.Close()
	b.httpBodyContent, err = io.ReadAll(r.Body)
	return err
}

func (b *BaseNotify) GetHttpRequest() *http.Request {
	return b.req
}

func (b *BaseNotify) GetBodyContent() []byte {
	return b.httpBodyContent
}

func (b *BaseNotify) GetBodyContentString() string {
	return string(b.httpBodyContent)
}
