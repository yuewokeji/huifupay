package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/notify"
	"github.com/yuewokeji/huifupay/response"
	"net/http"
)

// TradePaymentMicropay 聚合反扫接口
// https://paas.huifu.com/partners/api#/smzf/api_jhfs
// 最近更新时间：2023.4.28
func (c *Client) TradePaymentMicropay(ctx context.Context, req *TradePaymentMicropayRequest) (resp *TradePaymentMicropayResponse, err error) {
	request := newRequest(`/v2/trade/payment/micropay`, req)
	resp = &TradePaymentMicropayResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradePaymentMicropayAcctSplitInfo struct {
	HuifuID string `json:"huifu_id"`
	DivAmt  string `json:"div_amt"`
	AcctID  string `json:"acct_id"`
}

type TradePaymentMicropayAcctSplitBunch struct {
	AcctInfos []TradePaymentMicropayAcctSplitInfo `json:"acct_infos"`
}

type TradePaymentMicropayNotifyAcctSplitInfo struct {
	DivAmt   string `json:"div_amt"`
	HuifuID  string `json:"huifu_id"`
	AcctDate string `json:"acct_date"`
}

type TradePaymentMicropayNotifyAcctSplitBunch struct {
	FeeHuifuID  string                                    `json:"fee_huifu_id"`
	FeeAcctDate string                                    `json:"fee_acct_date"`
	AcctInfos   []TradePaymentMicropayNotifyAcctSplitInfo `json:"acct_infos"`
}

type TradePaymentMicropayNotifyCombinedPayData struct {
	HuifuID  string `json:"huifu_id"`
	UserType string `json:"user_type"`
	AcctId   string `json:"acct_id"`
	Amount   string `json:"amount"`
}

type TradePaymentMicropayNotifyFeeFormulaInfos struct {
	FeeFormula string `json:"fee_formula"`
	FeeType    string `json:"fee_type"`
	HuifuID    string `json:"huifu_id"`
}

type TradePaymentMicropayNotifyTransFeeAllowanceInfo struct {
	ReceivableFeeAmt string `json:"receivable_fee_amt"`
	ActualFeeAmt     string `json:"actual_fee_amt"`
	AllowanceFeeAmt  string `json:"allowance_fee_amt"`
	AllowanceType    string `json:"allowance_type"`
	NoAllowanceDesc  string `json:"no_allowance_desc"`
}

type TradePaymentMicropayNotifyWxResponse struct {
	SubAppid  string `json:"sub_appid"`  //子商户公众账号id
	Openid    string `json:"openid"`     //用户标识
	SubOpenid string `json:"sub_openid"` //子商户用户标识
}

type TradePaymentMicropayNotifyAlipayResponse struct {
	AppId          string `json:"app_id"`
	BuyerId        string `json:"buyer_id"`
	BuyerLogonId   string `json:"buyer_logon_id"`
	BuyerPayAmount string `json:"buyer_pay_amount"`
}

type TradePaymentMicropayNotifyUnionpayResponse struct {
}

type TradePaymentMicropayNotifyDcResponse struct {
}

type TradePaymentMicropayNotifyMerDevLocation struct {
}

type TradePaymentMicropayRiskCheckData struct {
	IpAddr      string `json:"ip_addr"`
	BaseStation string `json:"base_station"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

type TradePaymentMicropayRequest struct {
	ReqDate              string                             `json:"req_date"`
	ReqSeqId             string                             `json:"req_seq_id"`
	HuifuId              string                             `json:"huifu_id"`
	TransAmt             string                             `json:"trans_amt"`
	GoodsDesc            string                             `json:"goods_desc"`
	AuthCode             string                             `json:"auth_code"`
	TimeExpire           string                             `json:"time_expire"`
	FeeFlag              string                             `json:"fee_flag"`
	LimitPayType         string                             `json:"limit_pay_type"`
	DelayAcctFlag        string                             `json:"delay_acct_flag"`
	ChannelNo            string                             `json:"channel_no"`
	CombinedpayData      string                             `json:"combinedpay_data"`
	PayScene             string                             `json:"pay_scene"`
	AcctSplitBunch       string                             `json:"acct_split_bunch"`
	AcctSplitBunchObject TradePaymentMicropayAcctSplitBunch `json:"-" autoassign:"AcctSplitBunch"`
	TermDivCouponType    string                             `json:"term_div_coupon_type"`
	WxData               string                             `json:"wx_data"`
	AlipayData           string                             `json:"alipay_data"`
	UnionpayData         string                             `json:"unionpay_data"`
	RiskCheckData        string                             `json:"risk_check_data"`
	RiskCheckDataObject  TradePaymentMicropayRiskCheckData  `json:"-" autoassign:"RiskCheckData"`
	TerminalDeviceInfo   string                             `json:"terminal_device_info"`
	NotifyUrl            string                             `json:"notify_url"`
	Remark               string                             `json:"remark"`
	AcctId               string                             `json:"acct_id"`
}

type TradePaymentMicropayResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		response.DataHeader
		ReqDate          string `json:"req_date"`
		ReqSeqId         string `json:"req_seq_id"`
		HfSeqId          string `json:"hf_seq_id"`
		OutTransId       string `json:"out_trans_id"`
		PartyOrderId     string `json:"party_order_id"`
		HuifuId          string `json:"huifu_id"`
		TransAmt         string `json:"trans_amt"`
		PayAmt           string `json:"pay_amt"`
		SettlementAmt    string `json:"settlement_amt"`
		TradeType        string `json:"trade_type"`
		TransStat        string `json:"trans_stat"`
		AcctStat         string `json:"acct_stat"`
		DebitType        string `json:"debit_type"`
		BankCode         string `json:"bank_code"`
		BankDesc         string `json:"bank_desc"`
		DelayAcctFlag    string `json:"delay_acct_flag"`
		EndTime          string `json:"end_time"`
		WxUserId         string `json:"wx_user_id"`
		WxResponse       string `json:"wx_response"`
		AlipayResponse   string `json:"alipay_response"`
		DcResponse       string `json:"dc_response"`
		MerDevLocation   string `json:"mer_dev_location"`
		Remark           string `json:"remark"`
		FeeAmt           string `json:"fee_amt"`
		UnionpayResponse string `json:"unionpay_response"`
		AcctId           string `json:"acct_id"`
		FeeHuifuId       string `json:"fee_huifu_id"`
		FeeFlag          int    `json:"fee_flag"`
		DeviceType       string `json:"device_type"`
		FeeFormulaInfos  string `json:"fee_formula_infos"`
		AtuSubMerId      string `json:"atu_sub_mer_id"`
	} `json:"data"`
}

// TradePaymentMicropayNotify 异步通知
func (c *Client) TradePaymentMicropayNotify(ctx context.Context, req *http.Request) (n *TradePaymentMicropayNotify, err error) {
	n = &TradePaymentMicropayNotify{
		Notify: notify.NewBaseNotify(),
	}
	err = c.DoNotifyRequest(ctx, req, n)
	return
}

// TradePaymentMicropayNotify 异步通知
type TradePaymentMicropayNotify struct {
	notify.Notify
	RespCode              string                                          `json:"resp_code"`
	RespDesc              string                                          `json:"resp_desc"`
	HfSeqId               string                                          `json:"hf_seq_id"`
	ReqSeqId              string                                          `json:"req_seq_id"`
	ReqDate               string                                          `json:"req_date"`
	AcctDate              string                                          `json:"acct_date"`
	EndTime               string                                          `json:"end_time"`
	HuifuId               string                                          `json:"huifu_id"`
	TransType             string                                          `json:"trans_type"`
	TransStat             string                                          `json:"trans_stat"`
	TransAmt              string                                          `json:"trans_amt"`
	PayAmt                string                                          `json:"pay_amt"`
	SettlementAmt         string                                          `json:"settlement_amt"`
	FeeAmount             string                                          `json:"fee_amount"`
	FeeFlag               int                                             `json:"fee_flag"`
	OutTransId            string                                          `json:"out_trans_id"`
	PartyOrderId          string                                          `json:"party_order_id"`
	DebitType             string                                          `json:"debit_type"`
	AcctSplitBunch        TradePaymentMicropayNotifyAcctSplitBunch        `json:"acct_split_bunch"`
	IsDelayAcct           string                                          `json:"is_delay_acct"`
	IsDiv                 string                                          `json:"is_div"`
	WxUserId              string                                          `json:"wx_user_id"`
	WxResponse            TradePaymentMicropayNotifyWxResponse            `json:"wx_response"`
	AlipayResponse        TradePaymentMicropayNotifyAlipayResponse        `json:"alipay_response"`
	UnionpayResponse      TradePaymentMicropayNotifyUnionpayResponse      `json:"unionpay_response"`
	DcResponse            TradePaymentMicropayNotifyDcResponse            `json:"dc_response"`
	MerDevLocation        TradePaymentMicropayNotifyMerDevLocation        `json:"mer_dev_location"`
	TransFeeAllowanceInfo TradePaymentMicropayNotifyTransFeeAllowanceInfo `json:"trans_fee_allowance_info"`
	CombinedpayData       []TradePaymentMicropayNotifyCombinedPayData     `json:"combinedpay_data"`
	CombinedpayFeeAmt     string                                          `json:"combinedpay_fee_amt"`
	BankCode              string                                          `json:"bank_code"`
	BankMessage           string                                          `json:"bank_message"`
	Remark                string                                          `json:"remark"`
	NotifyType            int                                             `json:"notify_type"`
	DeviceType            string                                          `json:"device_type"`
	FeeFormulaInfos       []TradePaymentMicropayNotifyFeeFormulaInfos     `json:"fee_formula_infos"`
	SplitFeeInfo          string                                          `json:"split_fee_info"`
	AtuSubMerId           string                                          `json:"atu_sub_mer_id"`
	DevsId                string                                          `json:"devs_id"`
}
