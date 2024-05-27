package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// TradePaymentScanpayQuery 扫码交易查询
// https://paas.huifu.com/partners/api#/smzf/api_qrpay_cx?id=ewm
// 最近更新时间：2024.4.18
func (c *Client) TradePaymentScanpayQuery(ctx context.Context, req *TradePaymentScanpayQueryRequest) (resp *TradePaymentScanPayQueryResponse, err error) {
	request := newRequest(`/v2/trade/payment/scanpay/query`, req)
	resp = &TradePaymentScanPayQueryResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradePaymentScanpayQueryRequest struct {
	HuifuID     string `json:"huifu_id"`
	OrgReqDate  string `json:"org_req_date"`
	OutOrderID  string `json:"out_ord_id,omitempty"`
	OrgHFSeqID  string `json:"org_hf_seq_id,omitempty"`
	OrgReqSeqID string `json:"org_req_seq_id,omitempty"`
}

type TradePaymentScanPayQueryResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		response.DataHeader
		BagentId                  string                                   `json:"bagent_id"`         // 渠道商商户号
		HuifuId                   string                                   `json:"huifu_id"`          // 商户号
		OrgReqDate                string                                   `json:"org_req_date"`      // 原交易请求日期
		OrgHfSeqId                string                                   `json:"org_hf_seq_id"`     // 交易返回的全局流水号
		OrgReqSeqId               string                                   `json:"org_req_seq_id"`    // 原交易请求流水号
		OutTransId                string                                   `json:"out_trans_id"`      //用户账单上的交易订单号
		PartyOrderId              string                                   `json:"party_order_id"`    //用户账单上的商户订单号
		TransAmt                  string                                   `json:"trans_amt"`         //交易金额 单位元
		PayAmt                    string                                   `json:"pay_amt"`           //消费者实付金额 单位元
		SettlementAmt             string                                   `json:"settlement_amt"`    //结算金额 单位元
		UnconfirmAmt              string                                   `json:"unconfirm_amt"`     //待确认总金额 单位元
		ConfirmedAmt              string                                   `json:"confirmed_amt"`     //已确认总金额 单位元
		TransType                 string                                   `json:"trans_type"`        //T_JSAPI: 微信公众号 T_MINIAPP: 微信小程序 A_JSAPI: 支付宝JS A_NATIVE: 支付宝正扫 U_NATIVE: 银联正扫 U_JSAPI: 银联JS D_NATIVE: 数字人民币正扫 T_H5：微信直连H5支付 T_APP：微信APP支付 T_NATIVE：微信正扫
		TransStat                 string                                   `json:"trans_stat"`        //交易状态 P：处理中；S：成功；F：失败；I: 初始
		TransTime                 string                                   `json:"trans_time"`        //交易时间 格式：HHMMSS
		EndTime                   string                                   `json:"end_time"`          //支付完成时间 格式yyyyMMddHHMMSS
		DelayAcctFlag             string                                   `json:"delay_acct_flag"`   //延迟标识 Y：延迟 N：不延迟
		AcctId                    string                                   `json:"acct_id"`           //账户号
		AcctDate                  string                                   `json:"acct_date"`         //账务日期 格式：yyyyMMdd
		AcctStat                  string                                   `json:"acct_stat"`         //账务状态 "I：初始"; "P:处理中"; "S:成功"; "F:失败
		DebitType                 string                                   `json:"debit_type"`        //借贷标识 D：借记卡，C：信用卡，Z：借贷合一卡，O：其他；
		FeeHuifuId                string                                   `json:"fee_huifu_id"`      //手续费商户号
		FeeFormulaInfos           string                                   `json:"fee_formula_infos"` //手续费费率信息 交易成功时返回手续费费率信息
		FeeFormulaInfosObject     []TradePaymentScanpayQueryFeeFormulaInfo `autoassign:"FeeFormulaInfos"`
		FeeAmt                    string                                   `json:"fee_amt"`           //手续费金额 单位元
		FeeType                   string                                   `json:"fee_type"`          //手续费扣款标志 INNER：内扣，OUTSIDE：外扣；
		WxUserId                  string                                   `json:"wx_user_id"`        //微信用户唯一标识码
		WxResponse                string                                   `json:"wx_response"`       //微信返回的响应报文
		AlipayResponse            string                                   `json:"alipay_response"`   //支付宝返回的响应报文
		UnionpayResponse          string                                   `json:"unionpay_response"` //银联返回的响应报文
		DivFlag                   string                                   `json:"div_flag"`          //是否分账交易 Y：分账交易，N：非分账交易；
		AcctSplitBunch            string                                   `json:"acct_split_bunch"`  //分账对象
		AcctSplitBunchObject      TradePaymentScanpayQueryAcctSplitBunch   `json:"-" autoassign:"AcctSplitBunch"`
		SplitFeeInfo              string                                   `json:"split_fee_info"` //分账手续费信息
		SplitFeeInfoObject        TradePaymentScanpayQuerySplitFeeInfo     `json:"-" autoassign:"SplitFeeInfo"`
		CombinedpayData           string                                   `json:"combinedpay_data"`               //补贴支付信息 jsonArray字符串
		CombinedpayFeeAmt         string                                   `json:"combinedpay_fee_amt"`            //补贴支付手续费金额 单位元
		TransFeeAllowanceInfo     string                                   `json:"trans_fee_allowance_info"`       //手续费补贴信息 jsonObject
		Remark                    string                                   `json:"remark"`                         //备注
		DeviceType                string                                   `json:"device_type"`                    //终端类型 01-智能POS 02-扫码POS 03-云音箱 04-台牌 05-云打印 06-扫脸设备 07-收银机 08-收银助手 09-传统POS 10-一体音箱 11-虚拟终端
		MerDevLocation            string                                   `json:"mer_dev_location"`               //商户终端定位 jsonObject字符串
		MerPriv                   string                                   `json:"mer_priv"`                       //商户私有域
		AuthNo                    string                                   `json:"auth_no"`                        //授权号
		PasswordTrade             string                                   `json:"password_trade"`                 //输入密码提示 Y-等待用户输入密码状态
		MerName                   string                                   `json:"mer_name"`                       //商户名称
		ShopName                  string                                   `json:"shop_name"`                      //店铺名称
		PreAuthAmt                string                                   `json:"pre_auth_amt"`                   //预授权金额 单位元，需保留小数点后两位
		PreAuthPayAmt             string                                   `json:"pre_auth_pay_amt"`               //预授权完成金额 单位元，需保留小数点后两位
		OrgAuthNo                 string                                   `json:"org_auth_no"`                    //原授权号
		PreAuthHfSeqId            string                                   `json:"pre_auth_hf_seq_id"`             //预授权汇付全局流水号
		PreAuthPayFeeAmount       string                                   `json:"pre_auth_pay_fee_amount"`        //预授权完成手续费 单位元，需保留小数点后两位
		PreAuthPayRefundFeeAmount string                                   `json:"pre_auth_pay_refund_fee_amount"` //预授权完成退还手续费 单位元，需保留小数点后两位
		OrgFeeFlag                string                                   `json:"org_fee_flag"`                   //原手续费扣款标志 INNER：内扣， OUTSIDE：外扣；
		OrgFeeRecType             string                                   `json:"org_fee_rec_type"`               //原手续费扣取方式 1:实收,2:后收；
		OrgAllowanceType          string                                   `json:"org_allowance_type"`             //原补贴类型 0-不补贴，1-补贴,2-部分补贴；
		FqChannels                string                                   `json:"fq_channels"`                    //信用卡分期资产方式 代表优先使用资产类型；alipayfq_cc：表示信⽤卡分期
		BankCode                  string                                   `json:"bank_code"`                      //外部通道返回码
		BankDesc                  string                                   `json:"bank_desc"`                      //外部通道返回描述
		AtuSubMerId               string                                   `json:"atu_sub_mer_id"`                 //ATU真实商户号
		DcResponse                string                                   `json:"dc_response"`                    //数字货币返回的响应报文 JsonObject
	} `json:"data"`
}

type TradePaymentScanpayQueryFeeFormulaInfo struct {
	FeeFormula string `json:"fee_formula"` //手续费公式 示例值：AMT*0.003
	FeeType    string `json:"fee_type"`    //手续费类型 TRANS_FEE：交易手续费；ACCT_FEE：组合支付账户补贴手续费
	HuifuId    string `json:"huifu_id"`    //商户号 组合支付账户补贴时，补贴账户的huifuId
}

type TradePaymentScanpayQueryAcctSplitInfo struct {
	DivAmt  string `json:"div_amt"`  //分账金额 单位元，需保留小数点后两位
	HuifuId string `json:"huifu_id"` //分账接收方ID
	AcctId  string `json:"acct_id"`  //分账接收方子账户
}

type TradePaymentScanpayQueryAcctSplitBunch struct {
	AcctInfos []TradePaymentScanpayQueryAcctSplitInfo `json:"acct_infos"`
}

type TradePaymentScanpayQuerySplitFeeDetails struct {
	SplitFeeAmt     string `json:"split_fee_amt"`      //分账手续费金额(元)
	SplitFeeHuifuId string `json:"split_fee_huifu_id"` //分账手续费商户号
	SplitFeeAcctId  string `json:"split_fee_acct_id"`  //分账手续费承担方账号
}

type TradePaymentScanpayQuerySplitFeeInfo struct {
	TotalSplitFeeAmt string                                    `json:"total_split_fee_amt"` //分账手续费总金额(元)
	SplitFeeFlag     int                                       `json:"split_fee_flag"`      //分账手续费扣款标志 1: 外扣 2: 内扣；
	SplitFeeDetails  []TradePaymentScanpayQuerySplitFeeDetails `json:"split_fee_details"`   //分账手续费明细
}
