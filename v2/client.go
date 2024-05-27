package v2

import (
	"github.com/yuewokeji/huifupay"
)

func NewClient(config huifupay.Config, options ...huifupay.ClientOption) *Client {
	c := &Client{}
	c.Client = huifupay.NewClient(config, options...)
	return c
}

type Client struct {
	*huifupay.Client
}
