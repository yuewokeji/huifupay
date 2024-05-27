package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/notify"
	"github.com/yuewokeji/huifupay/response"
	"net/http"
)

// TradeSettlementEnchashment 提现接口
// https://paas.huifu.com/partners/api#/jyjs/qx/api_qx
// 最近更新时间：2023.4.26
func (c *Client) TradeSettlementEnchashment(ctx context.Context, req *TradeSettlementEnchashmentRequest) (resp *TradeSettlementEnchashmentResponse, err error) {
	request := newRequest(`/v2/trade/settlement/enchashment`, req)
	resp = &TradeSettlementEnchashmentResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradeSettlementEnchashmentRequest struct {
	ReqDate            string `json:"req_date"`
	ReqSeqId           string `json:"req_seq_id"`
	CashAmt            string `json:"cash_amt"` //单位元
	HuifuId            string `json:"huifu_id"`
	AcctId             string `json:"acct_id"`
	IntoAcctDateType   string `json:"into_acct_date_type"`
	TokenNo            string `json:"token_no"`
	EnchashmentChannel string `json:"enchashment_channel"`
	Remark             string `json:"remark"`
	NotifyUrl          string `json:"notify_url"`
}

type TradeSettlementEnchashmentResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		response.DataHeader
		ReqDate   string `json:"req_date"`
		ReqSeqId  string `json:"req_seq_id"`
		HfSeqId   string `json:"hf_seq_id"`
		TransStat string `json:"trans_stat"`
		HuifuId   string `json:"huifu_id"`
		AcctId    string `json:"acct_id"`
	} `json:"data"`
}

// TradeSettlementEnchashmentNotify 异步通知
func (c *Client) TradeSettlementEnchashmentNotify(ctx context.Context, req *http.Request) (n *TradeSettlementEnchashmentNotify, err error) {
	n = &TradeSettlementEnchashmentNotify{
		Notify: notify.NewBaseNotify(),
	}
	err = c.DoNotifyRequest(ctx, req, n)
	return
}

// TradeSettlementEnchashmentNotify 异步通知
type TradeSettlementEnchashmentNotify struct {
	notify.Notify
	SubRespCode   string `json:"sub_resp_code"`
	SubRespDesc   string `json:"sub_resp_desc"`
	ReqSeqId      string `json:"req_seq_id"`
	ReqDate       string `json:"req_date"`
	HfSeqId       string `json:"hf_seq_id"`
	TransStatus   string `json:"trans_status"`
	AcctStatus    string `json:"acct_status"`
	ChannelStatus string `json:"channel_status"`
	FeeAmt        string `json:"fee_amt"`
	CashAmt       string `json:"cash_amt"`
	MsgType       string `json:"msg_type"`
}
