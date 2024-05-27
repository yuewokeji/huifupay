package webhook

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/yuewokeji/huifupay/sign"
	"net/http"
	"testing"
)

type EventTest struct {
	EventDefineNO string `json:"event_define_no"`
	Data          string `json:"data"`
}

func Test_WebHook(t *testing.T) {
	event := EventTest{
		EventDefineNO: "pay.wx_pub",
		Data:          "data",
	}
	body := []byte(`{"event_define_no":"pay.wx_pub","data":"data"}`)

	hook := NewWithSigner([]byte("asdfkjasdkjfadjkfasdf"))
	hook.Register(EventPayWxPub, func(ctx context.Context, data []byte) error {
		e := &EventTest{}
		err := json.Unmarshal(data, e)
		assert.Nil(t, err)
		assert.Equal(t, event.EventDefineNO, e.EventDefineNO)
		assert.Equal(t, event.Data, e.Data)
		return err
	})

	req, _ := http.NewRequest("POST", "https://example.com/callback?sign=1522304a51d9b18f076e7866280edce9", bytes.NewBuffer(body))
	hr, err := hook.HandleRequest(context.Background(), req)
	assert.Nil(t, err)
	assert.Equal(t, hr.Handler, string(EventPayWxPub))
	assert.Equal(t, EventPayWxPub, hr.Event)
	assert.Equal(t, GetEventText(EventPayWxPub), hr.EventText)
	assert.Equal(t, body, hr.Data)
}

func Test_DefaultHandler(t *testing.T) {
	event := EventTest{
		EventDefineNO: "pay.wx_pub",
		Data:          "data",
	}
	body := []byte(`{"event_define_no":"pay.wx_pub","data":"data"}`)

	hook := NewWithSigner([]byte("asdfkjasdkjfadjkfasdf"))
	hook.SetDefaultHandler(func(ctx context.Context, data []byte) error {
		e := &EventTest{}
		err := json.Unmarshal(data, e)
		assert.Nil(t, err)
		assert.Equal(t, event.EventDefineNO, e.EventDefineNO)
		assert.Equal(t, event.Data, e.Data)
		return err
	})

	req, _ := http.NewRequest("POST", "https://example.com/callback?sign=1522304a51d9b18f076e7866280edce9", bytes.NewBuffer(body))
	hr, err := hook.HandleRequest(context.Background(), req)
	assert.Nil(t, err)
	assert.Equal(t, hr.Handler, "DefaultHandler")
	assert.Equal(t, EventPayWxPub, hr.Event)
	assert.Equal(t, GetEventText(EventPayWxPub), hr.EventText)
	assert.Equal(t, body, hr.Data)
}

func Test_SignVerifyFailed(t *testing.T) {
	body := []byte(`{"event_define_no":"pay.wx_pub","data":"datadata"}`)

	hook := NewWithSigner([]byte("asdfkjasdkjfadjkfasdf"))
	req, _ := http.NewRequest("POST", "https://example.com/callback?sign=1522304a51d9b18f076e7866280edce9", bytes.NewBuffer(body))
	_, err := hook.HandleRequest(context.Background(), req)
	assert.True(t, errors.Is(err, sign.ErrSignVerification))
}
