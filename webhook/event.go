package webhook

type Event string

// https://paas.huifu.com/partners/devtools#/webhook/webhook_jieshao?id=webhook%e4%ba%8b%e4%bb%b6%e7%b1%bb%e5%9e%8b%e5%88%97%e8%a1%a8
// 最近更新时间：2024.4.18
const (
	EventPayWxPub                   Event = "pay.wx_pub"
	EventPayWxLite                  Event = "pay.wx_lite"
	EventPayAliJs                   Event = "pay.ali_js"
	EventPayAliQr                   Event = "pay.ali_qr"
	EventPayUnionQr                 Event = "pay.union_qr"
	EventPayUnionJs                 Event = "pay.union_js"
	EventPayUnionOnline             Event = "pay.union_online"
	EventPayDigitQr                 Event = "pay.digit_qr"
	EventRefundStandard             Event = "refund.standard"
	EventTransClose                 Event = "trans.close"
	EventMerRegisterHuifu           Event = "mer.register.huifu"
	EventConfWxMiniAPP              Event = "conf.wx.miniAPP"
	EventConfWxWoa                  Event = "conf.wx.woa"
	EventConfUnionQr                Event = "conf.union.qr"
	EventConfBankCardPay            Event = "conf.bankCard.pay"
	EventMerRegisterAli             Event = "mer.register.ali"
	EventMerRegisterWx              Event = "mer.register.wx"
	EventMerUpgradeHuifu            Event = "mer.upgrade.huifu"
	EventMerBizModifyHuifu          Event = "mer.biz.modify.huifu"
	EventPayCardConsumePos          Event = "pay.card_consume.pos"
	EventReversalPos                Event = "reversal.pos"
	EventRefundCardConsumePos       Event = "refund.card_consume.pos"
	EventPerAuthPos                 Event = "per_auth.pos"
	EventPerAuthReversalPos         Event = "per_auth.reversal.pos"
	EventPerAuthCompletePos         Event = "per_auth.complete.pos"
	EventPreAuthCompleteReversalPos Event = "pre_auth.complete_reversal.pos"
	EventPayDigitScan               Event = "pay.digit_scan"
	EventPayUnionScan               Event = "pay.union_scan"
	EventPayWxScan                  Event = "pay.wx_scan"
	EventPayAliScan                 Event = "pay.ali_scan"
	EventPayWxScaned                Event = "pay.wx_scaned"
	EventPayAliScaned               Event = "pay.ali_scaned"
	EventPayUnionScaned             Event = "pay.union_scaned"
	EventStatementDay               Event = "statement.day"
	EventStatementAuto              Event = "statement.auto"
	EventFundBankDeposit            Event = "fund.bank_deposit"
)

var eventMap = map[Event]string{
	EventPayWxPub:                   "微信公众号支付",
	EventPayWxLite:                  "微信小程序支付",
	EventPayAliJs:                   "支付宝JS支付",
	EventPayAliQr:                   "支付宝二维码支付",
	EventPayUnionQr:                 "银联二维码支付",
	EventPayUnionJs:                 "银联JS支付",
	EventPayUnionOnline:             "银联统一收银台",
	EventPayDigitQr:                 "数字货币二维码支付",
	EventRefundStandard:             "交易退款",
	EventTransClose:                 "交易关单",
	EventMerRegisterHuifu:           "商户业务入驻",
	EventConfWxMiniAPP:              "微信小程序配置",
	EventConfWxWoa:                  "微信公众号配置",
	EventConfUnionQr:                "银联二维码开通",
	EventConfBankCardPay:            "银行卡开通",
	EventMerRegisterAli:             "支付宝入驻",
	EventMerRegisterWx:              "微信入驻",
	EventMerUpgradeHuifu:            "商户升级",
	EventMerBizModifyHuifu:          "商户业务变更",
	EventPayCardConsumePos:          "终端刷卡消费",
	EventReversalPos:                "终端刷卡消费撤销",
	EventRefundCardConsumePos:       "终端刷卡消费退货",
	EventPerAuthPos:                 "终端刷卡预授权",
	EventPerAuthReversalPos:         "终端刷卡预授权撤销",
	EventPerAuthCompletePos:         "终端刷卡预授权完成",
	EventPreAuthCompleteReversalPos: "终端刷卡预授权完成撤销",
	EventPayDigitScan:               "终端正扫-数字货币",
	EventPayUnionScan:               "终端正扫-银联",
	EventPayWxScan:                  "终端正扫-微信",
	EventPayAliScan:                 "终端正扫-支付宝",
	EventPayWxScaned:                "终端反扫-微信",
	EventPayAliScaned:               "终端反扫-支付宝",
	EventPayUnionScaned:             "终端反扫-银联",
	EventStatementDay:               "结算-日结算通知",
	EventStatementAuto:              "结算-结算通知",
	EventFundBankDeposit:            "银行入账通知",
}

func GetEventText(e Event) string {
	if s, ok := eventMap[e]; ok {
		return s
	}
	return ""
}
