package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// TradeAcctpaymentAcctlogQuery 账务流水查询
// https://paas.huifu.com/partners/api#/yuer/api_acctlscx
// 最近更新时间：2023.4.29
func (c *Client) TradeAcctpaymentAcctlogQuery(ctx context.Context, req *TradeAcctpaymentAcctlogQueryRequest) (resp *TradeAcctpaymentAcctlogQueryResponse, err error) {
	request := newRequest(`/v2/trade/acctpayment/acctlog/query`, req)
	resp = &TradeAcctpaymentAcctlogQueryResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradeAcctpaymentAcctlogQueryRequest struct {
	ReqSeqId string `json:"req_seq_id"`
	HuifuId  string `json:"huifu_id"`
	AcctDate string `json:"acct_date"`
	PageSize string `json:"page_size"`
	PageNum  string `json:"page_num"`
	AcctId   string `json:"acct_id"`
}

type TradeAcctpaymentAcctlogQueryResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		response.DataHeader
		HuifuId           string                            `json:"huifu_id"`
		MerName           string                            `json:"mer_name"`
		MerShortName      string                            `json:"mer_short_name"`
		AcctType          string                            `json:"acct_type"`
		AcctId            string                            `json:"acct_id"`
		PageSize          string                            `json:"page_size"`
		PageNum           string                            `json:"page_num"`
		ResultCount       string                            `json:"result_count"`
		AcctLogList       string                            `json:"acct_log_list"`
		AcctLogListObject []TradeAcctpaymentAcctlogListItem `autoassign:"AcctLogList"`
	} `json:"data"`
}

type TradeAcctpaymentAcctlogListItem struct {
	AcctLogId           string `json:"acct_log_id"`
	AcctDate            string `json:"acct_date"`
	TransDateTime       string `json:"trans_date_time"`
	AcctTransType       string `json:"acct_trans_type"`
	AcctTransTypeCode   string `json:"acct_trans_type_code"`
	AcctProductReqSeqId string `json:"acct_product_req_seq_id"`
	DebitType           string `json:"debit_type"`
	TransAmt            string `json:"trans_amt"`
	FeeAmt              string `json:"fee_amt"`
	BalanceAmt          string `json:"balance_amt"`
	HfSeqId             string `json:"hf_seq_id"`
}
