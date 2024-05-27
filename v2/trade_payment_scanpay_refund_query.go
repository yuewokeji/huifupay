package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// TradePaymentScanpayRefundQuery 扫码交易退款查询接口
// https://paas.huifu.com/partners/api#/smzf/api_qrpay_tkcx?id=appwx
// 最近更新时间：2023.4.28
func (c *Client) TradePaymentScanpayRefundQuery(ctx context.Context, req *TradePaymentScanpayRefundQueryRequest) (resp *TradePaymentScanpayRefundQueryResponse, err error) {
	request := newRequest(`/v2/trade/payment/scanpay/refundquery`, req)
	resp = &TradePaymentScanpayRefundQueryResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradePaymentScanpayRefundQueryRequest struct {
	HuifuId     string `json:"huifu_id"`
	OrgReqDate  string `json:"org_req_date"`
	OrgHfSeqId  string `json:"org_hf_seq_id,omitempty"`
	OrgReqSeqId string `json:"org_req_seq_id,omitempty"`
	MerOrdId    string `json:"mer_ord_id,omitempty"`
}

type TradePaymentScanpayRefundQueryResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		response.DataHeader
		HuifuId               string `json:"huifu_id"`
		OrgHfSeqId            string `json:"org_hf_seq_id"`
		OrgReqDate            string `json:"org_req_date"`
		OrgReqSeqId           string `json:"org_req_seq_id"`
		OrdAmt                string `json:"ord_amt"`
		ActualRefAmt          string `json:"actual_ref_amt"`
		TransDate             string `json:"trans_date"`
		TransTime             string `json:"trans_time"`
		TransType             string `json:"trans_type"`
		TransStat             string `json:"trans_stat"`
		BankCode              string `json:"bank_code"`
		BankMessage           string `json:"bank_message"`
		FeeAmt                string `json:"fee_amt"`
		AcctSplitBunch        string `json:"acct_split_bunch"`
		SplitFeeInfo          string `json:"split_fee_info"`
		CombinedpayData       string `json:"combinedpay_data"`
		CombinedpayFeeAmt     string `json:"combinedpay_fee_amt"`
		DcResponse            string `json:"dc_response"`
		OrgPartyOrderId       string `json:"org_party_order_id"`
		AuthNo                string `json:"auth_no"`
		DebitFlag             string `json:"debit_flag"`
		MerName               string `json:"mer_name"`
		MerPriv               string `json:"mer_priv"`
		OrgAuthNo             string `json:"org_auth_no"`
		OrgOutOrderId         string `json:"org_out_order_id"`
		PreAuthCanceFeeAmount string `json:"pre_auth_cance_fee_amount"`
		PreAuthCancelAmt      string `json:"pre_auth_cancel_amt"`
		PreAuthHfSeqId        string `json:"pre_auth_hf_seq_id"`
		ShopName              string `json:"shop_name"`
		FqAcqOrdAmt           string `json:"fq_acq_ord_amt"`
		FqAcqFeeAmt           string `json:"fq_acq_fee_amt"`
		OthOrdAmt             string `json:"oth_ord_amt"`
		OthFeeAmt             string `json:"oth_fee_amt"`
		WxResponse            string `json:"wx_response"`
		AlipayResponse        string `json:"alipay_response"`
		TransFinishTime       string `json:"trans_finish_time"`
		UnionpayResponse      string `json:"unionpay_response"`
	} `json:"data"`
}
