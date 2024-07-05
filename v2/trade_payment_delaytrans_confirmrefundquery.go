package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// TradePaymentDelaytransConfirmrefundquery 交易确认退款查询接口
// https://paas.huifu.com/partners/api#/smzf/api_jyqrtkcq
// 最近更新时间：2023.5.23
func (c *Client) TradePaymentDelaytransConfirmrefundquery(ctx context.Context, req *TradePaymentDelaytransConfirmrefundqueryRequest) (resp *TradePaymentDelaytransConfirmrefundqueryResponse, err error) {
	request := newRequest(`/v2/trade/payment/delaytrans/confirmrefundquery`, req)
	resp = &TradePaymentDelaytransConfirmrefundqueryResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradePaymentDelaytransConfirmrefundqueryRequest struct {
	HuifuId     string `json:"huifu_id"`
	OrgReqDate  string `json:"org_req_date"`
	OrgReqSeqId string `json:"org_req_seq_id,omitempty"`
	OrgHfSeqId  string `json:"org_hf_seq_id,omitempty"`
}

type TradePaymentDelaytransConfirmrefundqueryPayConfirmAcctDetail struct {
	HuifuId   string `json:"huifu_id"`
	AcctId    string `json:"acct_id"`
	DivAmt    string `json:"div_amt"`
	TransStat string `json:"trans_stat"`
}

type TradePaymentDelaytransConfirmrefundqueryResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		RespCode                   string                                                         `json:"resp_code"`
		RespDesc                   string                                                         `json:"resp_desc"`
		HuifuId                    string                                                         `json:"huifu_id"`
		OrgReqDate                 string                                                         `json:"org_req_date"`
		OrgHfSeqId                 string                                                         `json:"org_hf_seq_id"`
		OrgReqSeqId                string                                                         `json:"org_req_seq_id"`
		TransStat                  string                                                         `json:"trans_stat"`
		PayConfirmAcctDetails      string                                                         `json:"pay_confirm_acct_details"`
		PayConfirmAcctDetailObject []TradePaymentDelaytransConfirmrefundqueryPayConfirmAcctDetail `json:"-" autoassign:"PayConfirmAcctDetails"`
	} `json:"data"`
}
