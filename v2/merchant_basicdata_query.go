package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// MerchantBasicdataQuery 商户详细信息查询
// https://paas.huifu.com/partners/api#/shgl/shjj/api_shjj_shxxxxcx_kyc
// 最近更新时间：2023.4.26
func (c *Client) MerchantBasicdataQuery(ctx context.Context, req *MerchantBasicdataQueryRequest) (resp *MerchantBasicdataQueryResponse, err error) {
	request := newRequest(`/v2/merchant/basicdata/query`, req)
	resp = &MerchantBasicdataQueryResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type MerchantBasicdataQueryRequest struct {
	ReqSeqId string `json:"req_seq_id"`
	ReqDate  string `json:"req_date"`
	HuifuId  string `json:"huifu_id"`
}

type MerchantBasicdataQueryResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		RespCode                   string                               `json:"resp_code"`
		RespDesc                   string                               `json:"resp_desc"`
		ProductId                  string                               `json:"product_id"`
		UpperHuifuId               string                               `json:"upper_huifu_id"`
		ExtMerId                   string                               `json:"ext_mer_id"`
		RegName                    string                               `json:"reg_name"`
		ShortName                  string                               `json:"short_name"`
		ReceiptName                string                               `json:"receipt_name"`
		Remarks                    string                               `json:"remarks"`
		CustType                   string                               `json:"cust_type"`
		EntType                    string                               `json:"ent_type"`
		BusiType                   string                               `json:"busi_type"`
		Mcc                        string                               `json:"mcc"`
		LicenseType                string                               `json:"license_type"`
		LicenseCode                string                               `json:"license_code"`
		LicenseValidityType        string                               `json:"license_validity_type"`
		LicenseBeginDate           string                               `json:"license_begin_date"`
		LicenseEndDate             string                               `json:"license_end_date"`
		RegProvId                  string                               `json:"reg_prov_id"`
		RegAreaId                  string                               `json:"reg_area_id"`
		RegDistrictId              string                               `json:"reg_district_id"`
		RegDetail                  string                               `json:"reg_detail"`
		ProvId                     string                               `json:"prov_id"`
		AreaId                     string                               `json:"area_id"`
		DistrictId                 string                               `json:"district_id"`
		DetailAddr                 string                               `json:"detail_addr"`
		LegalName                  string                               `json:"legal_name"`
		LegalCertType              string                               `json:"legal_cert_type"`
		LegalCertNo                string                               `json:"legal_cert_no"`
		LegalCertValidityType      string                               `json:"legal_cert_validity_type"`
		LegalCertBeginDate         string                               `json:"legal_cert_begin_date"`
		LegalCertEndDate           string                               `json:"legal_cert_end_date"`
		LegalMobileNo              string                               `json:"legal_mobile_no"`
		LegalAddr                  string                               `json:"legal_addr"`
		Occupation                 string                               `json:"occupation"`
		BeneficiaryInfoList        string                               `json:"beneficiary_info_list"`
		ContactName                string                               `json:"contact_name"`
		ContactMobileNo            string                               `json:"contact_mobile_no"`
		ContactEmail               string                               `json:"contact_email"`
		ServicePhone               string                               `json:"service_phone"`
		LoginName                  string                               `json:"login_name"`
		SmsSendFlag                string                               `json:"sms_send_flag"`
		MerUrl                     string                               `json:"mer_url"`
		MerIcp                     string                               `json:"mer_icp"`
		OpenLicenceNo              string                               `json:"open_licence_no"`
		QryCashCardInfoList        string                               `json:"qry_cash_card_info_list"`
		QryCashCardInfoListObject  []MerchantBasicdataQueryCashCardInfo `autoassign:"QryCashCardInfoList"`
		QryCashConfigList          string                               `json:"qry_cash_config_list"`
		QrySettleConfigList        string                               `json:"qry_settle_config_list"`
		CollectionSettleConfigList string                               `json:"collection_settle_config_list"`
		AgreementInfoList          string                               `json:"agreement_info_list"`
		SignUserInfoList           string                               `json:"sign_user_info_list"`
		OnlineBusiType             string                               `json:"online_busi_type"`
		OnlineMediaInfoList        string                               `json:"online_media_info_list"`
		QuickFlag                  string                               `json:"quick_flag"`
		OnlineFlag                 string                               `json:"online_flag"`
		WithholdFlag               string                               `json:"withhold_flag"`
		PreAuthorizationFlag       string                               `json:"pre_authorization_flag"`
		WebFlag                    string                               `json:"web_flag"`
		BalancePayFlag             string                               `json:"balance_pay_flag"`
		QryBalancePayConfig        string                               `json:"qry_balance_pay_config"`
		OnlineFeeConfList          string                               `json:"online_fee_conf_list"`
		DelayFlag                  string                               `json:"delay_flag"`
		ForcedDelayFlag            string                               `json:"forced_delay_flag"`
		OutFeeFlag                 string                               `json:"out_fee_flag"`
		OutFeeHuifuId              string                               `json:"out_fee_huifu_id"`
		OutFeeAcctType             string                               `json:"out_fee_acct_type"`
		QryWxConfList              string                               `json:"qry_wx_conf_list"`
		QryAliConfList             string                               `json:"qry_ali_conf_list"`
		QryBankCardConf            string                               `json:"qry_bank_card_conf"`
		QryUnionConf               string                               `json:"qry_union_conf"`
		BankBigAmtPayConfig        string                               `json:"bank_big_amt_pay_config"`
		OutOrderFundsMerge         string                               `json:"out_order_funds_merge"`
		CombinePayConfig           string                               `json:"combine_pay_config"`
		WxZlConf                   string                               `json:"wx_zl_conf"`
		AliZlConfList              string                               `json:"ali_zl_conf_list"`
		FileInfoList               string                               `json:"file_info_list"`
		ReconRespAddr              string                               `json:"recon_resp_addr"`
		EnterFee                   float64                              `json:"enter_fee"`
		EnterFeeFlag               string                               `json:"enter_fee_flag"`
		MerIdentity                string                               `json:"mer_Identity"`
		MerLevel                   string                               `json:"mer_level"`
		MerConfigInfo              string                               `json:"mer_config_info"`
		ElecAcctConfig             string                               `json:"elec_acct_config"`
		TaxConfig                  string                               `json:"tax_config"`
		OnlineRefund               string                               `json:"online_refund"`
		PlatformRefund             string                               `json:"platform_refund"`
		SupportRevoke              string                               `json:"support_revoke"`
		SplitBillResult            string                               `json:"split_bill_result"`
		ShareHolderInfoList        string                               `json:"share_holder_info_list"`
		HeadOfficeFlag             string                               `json:"head_office_flag"`
	} `json:"data"`
}

type MerchantBasicdataQueryCashCardInfo struct {
	CardType         string `json:"card_type"`
	CardName         string `json:"card_name"`
	CardNo           string `json:"card_no"`
	ProvId           string `json:"prov_id"`
	AreaId           string `json:"area_id"`
	BankCode         string `json:"bank_code"`
	BankName         string `json:"bank_name"`
	BranchCode       string `json:"branch_code"`
	BranchName       string `json:"branch_name"`
	CertType         string `json:"cert_type"`
	CertNo           string `json:"cert_no"`
	CertValidityType string `json:"cert_validity_type"`
	CertBeginDate    string `json:"cert_begin_date"`
	CertEndDate      string `json:"cert_end_date"`
	Status           string `json:"status"`
	TokenNo          string `json:"token_no"`
	IsSettleDefault  string `json:"is_settle_default"`
}
