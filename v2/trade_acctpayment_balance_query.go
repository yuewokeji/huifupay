package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// TradeAcctpaymentBalanceQuery 账户余额信息查询接口
// https://paas.huifu.com/partners/api#/jyjs/api_jyjs_yuexxcx
// 最近更新时间：2023.4.21
func (c *Client) TradeAcctpaymentBalanceQuery(ctx context.Context, req *TradeAcctpaymentBalanceQueryRequest) (resp *TradeAcctPaymentBalanceQueryResponse, err error) {
	request := newRequest(`/v2/trade/acctpayment/balance/query`, req)
	resp = &TradeAcctPaymentBalanceQueryResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradeAcctpaymentBalanceQueryRequest struct {
	ReqDate  string `json:"req_date"`
	ReqSeqID string `json:"req_seq_id"`
	HuifuID  string `json:"huifu_id"`
}

type TradeAcctPaymentBalanceQueryResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		response.DataHeader
		ReqDate            string                                     `json:"req_date"`
		ReqSeqID           string                                     `json:"req_seq_id"`
		AcctInfoList       string                                     `json:"acctinfo_list"`
		AcctInfoListObject []TradeAcctpaymentBalanceQueryAcctInfoList `autoassign:"AcctInfoList"`
	} `json:"data"`
}

type TradeAcctpaymentBalanceQueryAcctInfoList struct {
	HuifuID    string `json:"huifu_id"`
	AcctId     string `json:"acct_id"`
	AcctType   string `json:"acct_type"`
	BalanceAmt string `json:"balance_amt"`
	AvlBal     string `json:"avl_bal"`
	FrzBal     string `json:"frz_bal"`
	LastAvlBal string `json:"last_avl_bal"`
	AcctStat   string `json:"acct_stat"`
}
