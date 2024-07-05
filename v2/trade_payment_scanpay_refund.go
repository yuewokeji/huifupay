package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/notify"
	"github.com/yuewokeji/huifupay/response"
	"net/http"
)

// TradePaymentScanpayRefund 扫码交易退款
// https://paas.huifu.com/partners/api#/smzf/api_qrpay_tk?id=appwx
// 最近更新时间：2023.4.28
func (c *Client) TradePaymentScanpayRefund(ctx context.Context, req *TradePaymentScanpayRefundRequest) (resp *TradePaymentScanpayRefundResponse, err error) {
	request := newRequest(`/v2/trade/payment/scanpay/refund`, req)
	resp = &TradePaymentScanpayRefundResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradePaymentScanpayRefundNotifyAcctInfo struct {
	DivAmt  string `json:"div_amt"`
	HuifuID string `json:"huifu_id"`
}
type TradePaymentScanpayRefundNotifyAcctSplitBunch struct {
	FeeAmt    string                                    `json:"fee_amt"`
	AcctInfos []TradePaymentScanpayRefundNotifyAcctInfo `json:"acct_infos"`
}

type TradePaymentScanpayRefundNotifyWxResponse struct {
	SubAppid  string `json:"sub_appid"`  //子商户公众账号id
	Openid    string `json:"openid"`     //用户标识
	SubOpenid string `json:"sub_openid"` //子商户用户标识
}

type TradePaymentScanpayRefundNotifyDcResponse struct {
}

type TradePaymentScanpayRefundNotifyCombinedPayData struct {
}

type TradePaymentScanpayRefundNotifyUnionpayResponse struct {
}

type TradePaymentScanpayRefundAcctInfo struct {
	DivAmt      string `json:"div_amt"`
	HuifuId     string `json:"huifu_id"`
	PartLoanAmt string `json:"part_loan_amt,omitempty"`
}

type TradePaymentScanpayRefundAcctSplitBunch struct {
	AcctInfos []TradePaymentScanpayRefundAcctInfo `json:"acct_infos,omitempty"`
}

func (t TradePaymentScanpayRefundAcctSplitBunch) IsEmpty() bool {
	return len(t.AcctInfos) == 0
}

type TradePaymentScanpayRefundRequest struct {
	ReqDate              string                                  `json:"req_date"`
	ReqSeqId             string                                  `json:"req_seq_id"`
	HuifuId              string                                  `json:"huifu_id"`
	OrdAmt               string                                  `json:"ord_amt"`
	OrgReqDate           string                                  `json:"org_req_date"`
	OrgHfSeqId           string                                  `json:"org_hf_seq_id"`
	AcctSplitBunch       string                                  `json:"acct_split_bunch,omitempty"`
	AcctSplitBunchObject TradePaymentScanpayRefundAcctSplitBunch `json:"-" autoassign:"AcctSplitBunch"`
	OrgPartyOrderId      string                                  `json:"org_party_order_id"`
	OrgReqSeqId          string                                  `json:"org_req_seq_id"`
	WxData               string                                  `json:"wx_data"`
	DigitalCurrencyData  string                                  `json:"digital_currency_data"`
	CombinedpayData      string                                  `json:"combinedpay_data"`
	UnionpayData         string                                  `json:"unionpay_data"`
	Remark               string                                  `json:"remark"`
	LoanFlag             string                                  `json:"loan_flag"`
	LoanUndertaker       string                                  `json:"loan_undertaker"`
	LoanAcctType         string                                  `json:"loan_acct_type"`
	RiskCheckData        string                                  `json:"risk_check_data"`
	ErminalDeviceData    string                                  `json:"erminal_device_data"`
	NotifyUrl            string                                  `json:"notify_url"`
}

type TradePaymentScanpayRefundResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		response.DataHeader
		ProductId        string `json:"product_id"`
		HuifuId          string `json:"huifu_id"`
		ReqDate          string `json:"req_date"`
		ReqSeqId         string `json:"req_seq_id"`
		HfSeqId          string `json:"hf_seq_id"`
		OrgReqDate       string `json:"org_req_date"`
		OrgReqSeqId      string `json:"org_req_seq_id"`
		TransDate        string `json:"trans_date"`
		TransTime        string `json:"trans_time"`
		TransFinishTime  string `json:"trans_finish_time"`
		TransStat        string `json:"trans_stat"`
		OrdAmt           string `json:"ord_amt"`
		ActualRefAmt     string `json:"actual_ref_amt"`
		AcctSplitBunch   string `json:"acct_split_bunch"`
		WxResponse       string `json:"wx_response"`
		AlipayResponse   string `json:"alipay_response"`
		CombinedpayData  string `json:"combinedpay_data"`
		Remark           string `json:"remark"`
		LoanFlag         string `json:"loan_flag"`
		LoanUndertaker   string `json:"loan_undertaker"`
		LoanAcctType     string `json:"loan_acct_type"`
		BankCode         string `json:"bank_code"`
		BankMessage      string `json:"bank_message"`
		UnionpayResponse string `json:"unionpay_response"`
		DcResponse       string `json:"dc_response"`
	} `json:"data"`
}

func (c *Client) TradePaymentScanpayRefundNotify(ctx context.Context, req *http.Request) (n *TradePaymentScanpayRefundNotify, err error) {
	n = &TradePaymentScanpayRefundNotify{
		Notify: notify.NewBaseNotify(),
	}
	err = c.DoNotifyRequest(ctx, req, n)
	return
}

// TradePaymentScanpayRefundNotify 异步通知
type TradePaymentScanpayRefundNotify struct {
	notify.Notify
	RespCode         string                                           `json:"resp_code"`
	RespDesc         string                                           `json:"resp_desc"`
	HuifuId          string                                           `json:"huifu_id"`
	ReqDate          string                                           `json:"req_date"`
	ReqSeqId         string                                           `json:"req_seq_id"`
	HfSeqId          string                                           `json:"hf_seq_id"`
	OrgReqDate       string                                           `json:"org_req_date"`
	OrgReqSeqId      string                                           `json:"org_req_seq_id"`
	OrgOrdAmt        string                                           `json:"org_ord_amt"`
	OrgFeeAmt        string                                           `json:"org_fee_amt"`
	TransDate        string                                           `json:"trans_date"`
	TransTime        string                                           `json:"trans_time"`
	TransFinishTime  string                                           `json:"trans_finish_time"`
	TransType        string                                           `json:"trans_type"`
	TransStat        string                                           `json:"trans_stat"`
	OrdAmt           string                                           `json:"ord_amt"`
	ActualRefAmt     string                                           `json:"actual_ref_amt"`
	TotalRefAmt      string                                           `json:"total_ref_amt"`
	TotalRefFeeAmt   string                                           `json:"total_ref_fee_amt"`
	RefCut           string                                           `json:"ref_cut"`
	AcctSplitBunch   TradePaymentScanpayRefundNotifyAcctSplitBunch    `json:"acct_split_bunch"`
	PartyOrderId     string                                           `json:"party_order_id"`
	WxResponse       TradePaymentScanpayRefundNotifyWxResponse        `json:"wx_response"`
	DcResponse       TradePaymentScanpayRefundNotifyDcResponse        `json:"dc_response"`
	CombinedpayData  []TradePaymentScanpayRefundNotifyCombinedPayData `json:"combinedpay_data"`
	Remark           string                                           `json:"remark"`
	BankCode         string                                           `json:"bank_code"`
	BankMessage      string                                           `json:"bank_message"`
	UnionpayResponse TradePaymentScanpayRefundNotifyUnionpayResponse  `json:"unionpay_response"`
}
