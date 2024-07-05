package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// TradePaymentDelaytransConfirmquery 交易确认查询接口
// https://paas.huifu.com/partners/api#/smzf/api_jyqrcx
// 最近更新时间：2023.5.23
func (c *Client) TradePaymentDelaytransConfirmquery(ctx context.Context, req *TradePaymentDelaytransConfirmqueryRequest) (resp *TradePaymentDelaytransConfirmqueryResponse, err error) {
	request := newRequest(`/v2/trade/payment/delaytrans/confirmquery`, req)
	resp = &TradePaymentDelaytransConfirmqueryResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradePaymentDelaytransConfirmqueryRequest struct {
	OrgReqDate  string `json:"org_req_date"`
	OrgReqSeqId string `json:"org_req_seq_id"`
	HuifuId     string `json:"huifu_id"`
}

type TradePaymentDelaytransConfirmqueryAcctSplitBunch struct {
	HuifuId   string `json:"huifu_id"`
	AcctId    string `json:"acct_id"`
	DivAmt    string `json:"div_amt"`
	TransStat string `json:"trans_stat"`
}

type TradePaymentDelaytransConfirmqueryResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		RespCode             string                                             `json:"resp_code"`
		RespDesc             string                                             `json:"resp_desc"`
		HuifuId              string                                             `json:"huifu_id"`
		OrgReqSeqId          string                                             `json:"org_req_seq_id"`
		TransStat            string                                             `json:"trans_stat"`
		UnconfirmAmt         string                                             `json:"unconfirm_amt"`
		ConfirmedAmt         string                                             `json:"confirmed_amt"`
		HfSeqId              string                                             `json:"hf_seq_id"`
		AcctSplitBunch       string                                             `json:"acct_split_bunch"`
		AcctSplitBunchObject []TradePaymentDelaytransConfirmqueryAcctSplitBunch `json:"-" autoassign:"AcctSplitBunch"`
		HycFlag              string                                             `json:"hyc_flag"`
		HycAttachId          string                                             `json:"hyc_attach_id"`
	} `json:"data"`
}
