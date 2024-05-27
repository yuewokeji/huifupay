package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// TradePaymentDelaytransConfirmrefund 交易确认退款接口
// https://paas.huifu.com/partners/api#/smzf/api_jyqrtk
// 最近更新时间：2023.5.23
func (c *Client) TradePaymentDelaytransConfirmrefund(ctx context.Context, req *TradePaymentDelaytransConfirmrefundRequest) (resp *TradePaymentDelaytransConfirmrefundResponse, err error) {
	request := newRequest(`/v2/trade/payment/delaytrans/confirmrefund`, req)
	resp = &TradePaymentDelaytransConfirmrefundResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradePaymentDelaytransConfirmrefundRequest struct {
	ReqDate              string                                  `json:"req_date"`
	ReqSeqId             string                                  `json:"req_seq_id"`
	HuifuId              string                                  `json:"huifu_id"`
	OrgReqDate           string                                  `json:"org_req_date"`
	OrgReqSeqId          string                                  `json:"org_req_seq_id,omitempty"`
	AcctSplitBunch       string                                  `json:"acct_split_bunch"`
	AcctSplitBunchObject TradePaymentScanpayRefundAcctSplitBunch `json:"-" autoassign:"AcctSplitBunch"`
	LoanFlag             string                                  `json:"loan_flag"`
	LoanUndertaker       string                                  `json:"loan_undertaker"`
	LoanAcctType         string                                  `json:"loan_acct_type"`
}

type TradePaymentDelaytransConfirmrefundResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		RespCode       string `json:"resp_code"`
		RespDesc       string `json:"resp_desc"`
		TransStat      string `json:"trans_stat"`
		HfSeqId        string `json:"hf_seq_id"`
		ReqDate        string `json:"req_date"`
		ReqSeqId       string `json:"req_seq_id"`
		HuifuId        string `json:"huifu_id"`
		OrgReqSeqId    string `json:"org_req_seq_id"`
		LoanFlag       string `json:"loan_flag"`
		LoanUndertaker string `json:"loan_undertaker"`
		LoanAcctType   string `json:"loan_acct_type"`
		UnconfirmAmt   string `json:"unconfirm_amt"`
		ConfirmedAmt   string `json:"confirmed_amt"`
		OrgReqDate     string `json:"org_req_date"`
		OrgHfSeqId     string `json:"org_hf_seq_id"`
	} `json:"data"`
}
