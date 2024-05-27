package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/yuewokeji/huifupay/sign"
	"io"
	"net/http"
	"strings"
	"sync"
)

// https://paas.huifu.com/partners/devtools#/webhook/webhook_jieshao
// 最近更新时间：2024.4.18

type Handler func(ctx context.Context, data []byte) error

func NewWithSigner(accessKey []byte, options ...HookOption) *Hook {
	signer := sign.NewMd5Signer(accessKey, sign.AppendRight)
	options = append(options, WithSigner(signer))
	return New(options...)
}

func New(options ...HookOption) *Hook {
	h := &Hook{}
	h.initWithOptions(options...)
	return h
}

type Hook struct {
	signer sign.Signer

	handles        sync.Map
	defaultHandler Handler
}

func (h *Hook) initWithOptions(options ...HookOption) {
	if len(options) > 0 {
		for _, option := range options {
			option(h)
		}
	}
}

// Register 注册一个处理方法
// 注意：不会重复注册，如果注册过，会被覆盖
func (h *Hook) Register(e Event, fn Handler) {
	h.handles.Store(e, fn)
}

// Unregister 移除一个处理方法
func (h *Hook) Unregister(e Event) bool {
	_, ok := h.handles.LoadAndDelete(e)
	return ok
}

// SetDefaultHandler 没有注册的事件，可以通过SetDefaultHandle()统一处理
func (h *Hook) SetDefaultHandler(fn Handler) {
	h.defaultHandler = fn
}

type HandleResult struct {
	Data      []byte
	Event     Event
	EventText string
	Handler   string
}

func (h *Hook) HandleRequest(ctx context.Context, req *http.Request) (hr HandleResult, err error) {
	defer req.Body.Close()
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return hr, errors.Wrap(err, "read request")
	}
	debug.Printf(">> WebHook Handle Request: %s", req.URL.RequestURI())
	debug.Printf(">> WebHook Handle Body: %s", string(data))

	// 签名验证
	if h.signer != nil {
		q := req.URL.Query()
		signature := q.Get("sign")
		// 汇付返回的签名是大写的
		signature = strings.ToLower(signature)

		err = h.signer.Verify(data, []byte(signature))
		if err != nil {
			debug.Printf(">> WebHook Verify Sign Error: %s", err.Error())
			return
		}
	}
	return h.Handle(ctx, data)
}

// Handle 处理回调事件
// 注意：该方法并不检测签名
func (h *Hook) Handle(ctx context.Context, data []byte) (hr HandleResult, err error) {
	defer func() {
		debug.Printf(">> WebHook Event: %s %s %s", hr.Event, hr.EventText, hr.Handler)
		if err != nil {
			debug.Printf(">> WebHook Error: %s", err.Error())
		}
	}()
	debug.Printf(">> WebHook Handle Data: %s", string(data))

	hr.Data = data

	// 事件类型
	obj := &struct {
		EventDefineNO string `json:"event_define_no"`
	}{}
	err = json.Unmarshal(hr.Data, obj)
	if err != nil {
		return hr, errors.Wrap(err, "unmarshal request")
	}

	hr.Event = Event(obj.EventDefineNO)
	hr.EventText = GetEventText(hr.Event)

	if hr.EventText == "" {
		return hr, fmt.Errorf("event not found: %s", hr.Event)
	}

	handler, isDefault := h.getHandler(hr.Event)
	if handler == nil {
		return hr, fmt.Errorf("handler not configured: %s", hr.Event)
	}

	if isDefault {
		hr.Handler = "DefaultHandler"
	} else {
		hr.Handler = string(hr.Event)
	}
	err = handler(ctx, hr.Data)
	return
}

func (h *Hook) getHandler(e Event) (fn Handler, isDefault bool) {
	v, ok := h.handles.Load(e)
	if !ok {
		if h.defaultHandler != nil {
			return h.defaultHandler, true
		}
		return nil, false
	}
	fn, ok = v.(Handler)
	if !ok {
		return nil, false
	}
	return fn, false
}

type HookOption func(hook *Hook)

func WithSigner(singer sign.Signer) HookOption {
	return func(hook *Hook) {
		hook.signer = singer
	}
}
