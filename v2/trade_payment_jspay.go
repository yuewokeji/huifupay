package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/notify"
	"github.com/yuewokeji/huifupay/response"
	"net/http"
)

// TradePaymentJspay 聚合正扫接口
// https://paas.huifu.com/partners/api#/smzf/api_jhzs?id=h5wx
// 最近更新时间：2023.4.24
func (c *Client) TradePaymentJspay(ctx context.Context, req *TradePaymentJspayRequest) (resp *TradePaymentJspayResponse, err error) {
	request := newRequest(`/v2/trade/payment/jspay`, req)
	resp = &TradePaymentJspayResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradePaymentJspayRequest struct {
	ReqSeqId             string                          `json:"req_seq_id"`
	ReqDate              string                          `json:"req_date"`
	HuifuId              string                          `json:"huifu_id"`
	AcctId               string                          `json:"acct_id"`
	GoodsDesc            string                          `json:"goods_desc"`
	TradeType            string                          `json:"trade_type"`                    //T_JSAPI: 微信公众号 T_MINIAPP: 微信小程序 A_JSAPI: 支付宝JS A_NATIVE: 支付宝正扫 U_NATIVE: 银联正扫 U_JSAPI: 银联JS D_NATIVE: 数字人民币正扫 T_H5：微信直连H5支付 T_APP：微信APP支付 T_NATIVE：微信正扫
	TransAmt             string                          `json:"trans_amt"`                     //单位元，需保留小数点后两位
	TimeExpire           string                          `json:"time_expire"`                   //交易有效期
	WxData               string                          `json:"wx_data"`                       //微信参数集合
	WxDataObject         TradePaymentJspayWxData         `json:"-" autoassign:"WxData"`         //微信参数集合
	AlipayData           string                          `json:"alipay_data"`                   //支付宝参数集合
	UnionpayData         string                          `json:"unionpay_data"`                 //银联参数集合
	DcData               string                          `json:"dc_data"`                       //数字人民币参数集合
	DelayAcctFlag        string                          `json:"delay_acct_flag"`               //Y 为延迟 N为不延迟，不传默认N；
	FeeFlag              int8                            `json:"fee_flag"`                      //1: 外扣 2: 内扣 (默认取控台配置值)；
	AcctSplitBunch       string                          `json:"acct_split_bunch"`              //分账对象
	AcctSplitBunchObject TradePaymentJspayAcctSplitBunch `json:"-" autoassign:"AcctSplitBunch"` //分账对象
	TermDivCouponType    int8                            `json:"term_div_coupon_type"`          //传入分账遇到优惠的处理规则
	CombinedpayData      string                          `json:"combinedpay_data"`
	LimitPayType         string                          `json:"limit_pay_type"`
	FqMerDiscountFlag    string                          `json:"fq_mer_discount_flag"`
	ChannelNo            string                          `json:"channel_no"`
	PayScene             string                          `json:"pay_scene"`
	Remark               string                          `json:"remark"`
	RiskCheckData        string                          `json:"risk_check_data"`
	TerminalDeviceData   string                          `json:"terminal_device_data"`
	NotifyUrl            string                          `json:"notify_url"` //交易异步通知地址，http或https开头
}

type TradePaymentJspayAcctSplitInfo struct {
	HuifuId string `json:"huifu_id"` //分账接收方ID
	DivAmt  string `json:"div_amt"`  //分账金额	单位元，需保留小数点后两位
	AcctId  string `json:"acct_id"`  //分账账户	可指定账户号，仅支持基本户、现金户
}

type TradePaymentJspayAcctSplitBunch struct {
	AcctInfos []TradePaymentJspayAcctSplitInfo `json:"acct_infos"`
}

type TradePaymentJspayWxData struct {
	SubAppid  string `json:"sub_appid"`  //子商户公众账号id
	Openid    string `json:"openid"`     //用户标识
	SubOpenid string `json:"sub_openid"` //子商户用户标识
}

type TradePaymentAlipayResponse struct {
	AppId          string `json:"app_id"`
	BuyerId        string `json:"buyer_id"`
	BuyerLogonId   string `json:"buyer_logon_id"`
	BuyerPayAmount string `json:"buyer_pay_amount"`
}

type TradePaymentJspayPayInfo struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timestamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

type TradePaymentJspayResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		response.DataHeader
		ReqDate          string                   `json:"req_date"`
		ReqSeqId         string                   `json:"req_seq_id"`
		HfSeqId          string                   `json:"hf_seq_id"`
		TradeType        string                   `json:"trade_type"`
		TransAmt         string                   `json:"trans_amt"`
		TransStat        string                   `json:"trans_stat"`
		HuifuId          string                   `json:"huifu_id"`
		BankCode         string                   `json:"bank_code"`
		BankMessage      string                   `json:"bank_message"`
		DelayAcctFlag    string                   `json:"delay_acct_flag"`
		PayInfo          string                   `json:"pay_info"`
		PayInfoObject    TradePaymentJspayPayInfo `json:"-" autoassign:"PayInfo"`
		QrCode           string                   `json:"qr_code"`
		AlipayResponse   string                   `json:"alipay_response"`
		WxResponse       string                   `json:"wx_response"`
		WxResponseObject TradePaymentJspayWxData  `json:"-" autoassign:"WxResponse"`
		UnionpayResponse string                   `json:"unionpay_response"`
		Remark           string                   `json:"remark"`
	} `json:"data"`
}

func (c *Client) TradePaymentJspayNotify(ctx context.Context, req *http.Request) (n *TradePaymentJspayNotify, err error) {
	n = &TradePaymentJspayNotify{
		Notify: notify.NewBaseNotify(),
	}
	err = c.DoNotifyRequest(ctx, req, n)
	return
}

type TradePaymentJspayNotify struct {
	notify.Notify
	RespCode          string                  `json:"resp_code"`
	RespDesc          string                  `json:"resp_desc"`
	HuifuId           string                  `json:"huifu_id"`
	ReqSeqId          string                  `json:"req_seq_id"`
	ReqDate           string                  `json:"req_date"`
	TransType         string                  `json:"trans_type"`
	HfSeqId           string                  `json:"hf_seq_id"`
	OutTransId        string                  `json:"out_trans_id"`
	PartyOrderId      string                  `json:"party_order_id"`
	TransAmt          string                  `json:"trans_amt"`
	PayAmt            string                  `json:"pay_amt"`
	SettlementAmt     string                  `json:"settlement_amt"`
	EndTime           string                  `json:"end_time"`
	AcctDate          string                  `json:"acct_date"`
	TransStat         string                  `json:"trans_stat"`
	FeeAmount         string                  `json:"fee_amount"`
	CombinedpayFeeAmt string                  `json:"combinedpay_fee_amt"`
	DebitType         string                  `json:"debit_type"`
	IsDiv             string                  `json:"is_div"`
	IsDelayAcct       string                  `json:"is_delay_acct"`
	WxUserId          string                  `json:"wx_user_id"`
	WxResponse        TradePaymentJspayWxData `json:"wx_response"`
	DcResponse        string                  `json:"dc_response"`
	UnionpayResponse  string                  `json:"unionpay_response"`
	DeviceType        string                  `json:"device_type"`
	MerDevLocation    string                  `json:"mer_dev_location"`
	BankCode          string                  `json:"bank_code"`
	BankMessage       string                  `json:"bank_message"`
	Remark            string                  `json:"remark"`
	FqChannels        string                  `json:"fq_channels"`
	SplitFeeInfo      string                  `json:"split_fee_info"`
	AtuSubMerId       string                  `json:"atu_sub_mer_id"`
	DevsId            string                  `json:"devs_id"`
}
