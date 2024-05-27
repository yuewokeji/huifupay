package webhook

// TradeEvent 交易事件
type TradeEvent struct {
	Namespace          string                       `json:"namespace"`            // 斗拱webhook事件域名，统一为“opps-webhook”
	EventDefineNo      string                       `json:"event_define_no"`      // webhook事件类型编号，参见【webhook事件类型表】
	HuifuID            string                       `json:"huifu_id"`             // 客户号
	MerName            string                       `json:"mer_name"`             // 商户名称
	MerOrdID           string                       `json:"mer_ord_id"`           // 商户订单号
	MerPriv            string                       `json:"mer_priv"`             // 商户私有域
	Remark             string                       `json:"remark"`               // 备注
	AcctID             string                       `json:"acct_id"`              // 斗拱系统虚账户号
	TransType          string                       `json:"trans_type"`           // 交易类型
	TransAmt           string                       `json:"trans_amt"`            // 交易金额
	DiscountAmt        string                       `json:"discount_amt"`         // 优惠金额
	SettlementAmt      string                       `json:"settlement_amt"`       // 结算金额
	FeeAmount          string                       `json:"fee_amount"`           // 手续费金额
	FeeFlag            string                       `json:"fee_flag"`             // 手续费扣款标志
	TransDate          string                       `json:"trans_date"`           // 交易日期
	TransTime          string                       `json:"trans_time"`           // 交易时间
	TransStat          string                       `json:"trans_stat"`           // 交易状态
	IsDelayAcct        string                       `json:"is_delay_acct"`        // 是否延时交易
	IsDiv              string                       `json:"is_div"`               // 是否分账交易
	CouponInfos        []TradeEventCouponInfo       `json:"coupon_infos"`         // 优惠券信息，json数组
	AcctSplitBunch     []TradeEventAcctSplitBunch   `json:"acct_split_bunch"`     // 分账串
	OutTransID         string                       `json:"out_trans_id"`         // 微信、支付宝、银联扫码订单号
	PartyOrderID       string                       `json:"party_order_id"`       // 支付通道凭证号
	WxResponse         TradeEventWxResponse         `json:"wx_response"`          // 微信返回的报文
	TerminalDeviceInfo TradeEventTerminalDeviceInfo `json:"terminal_device_info"` // 设备信息
	AlipayResponse     TradeEventAlipayResponse     `json:"alipay_response"`      // 支付宝返回报文
	UnionpayResponse   TradeEventUnionpayResponse   `json:"unionpay_response"`    // 银联返回报文
	DcResponse         TradeEventDcResponse         `json:"dc_response"`          // 数字货币返回报文
	HfSeqID            string                       `json:"hf_seq_id"`            // 全局流水号
	BankSeqID          string                       `json:"bank_seq_id"`          // 凭证号
	BatchID            string                       `json:"batch_id"`             // 批次号
	RefNum             string                       `json:"ref_num"`              // 参考号
}

type TradeEventCouponInfo struct {
	CouponId           string `json:"couponId"`           // 券ID，券或者立减优惠id
	CouponName         string `json:"couponName"`         // 优惠名称，优惠名称
	CouponRange        string `json:"couponRange"`        // 优惠范围，GLOBAL-全场代金券, SINGLE-单品优惠, 折扣券
	CouponType         string `json:"couponType"`         // 优惠类型，COUPON-代金券，DISCOUNT-优惠券
	CouponAmt          string `json:"couponAmt"`          // 优惠券面额，用户享受优惠的金额
	ActiveId           string `json:"activeId"`           // 活动ID，在微信商户后台配置的批次ID
	MerchantContribute string `json:"merchantContribute"` // 商户出资，特指商户自己创建的优惠，出资金额等于本项优惠总金额，单位为元
	OtherContribute    string `json:"otherContribute"`    // 其他出资，其他出资方金额，可能是通道方，可能是品牌商，或者其他方，也可能是他们的一起出资
	GoodsInfo          string `json:"goodsInfo"`          // 单品信息，使用Json格式
	AddnInfo           string `json:"addnInfo"`           // 银联字段，内容自定义
}

type TradeEventAcctSplitBunch struct {
}

type TradeEventWxResponse struct {
}

type TradeEventTerminalDeviceInfo struct {
}

type TradeEventAlipayResponse struct {
}

type TradeEventUnionpayResponse struct {
}

type TradeEventDcResponse struct {
}
