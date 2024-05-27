package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// TradePaymentScanpayClose 扫码交易关单
// https://paas.huifu.com/partners/api#/smzf/api_qrpay_jygd
// 最近更新时间：2023.10.26
func (c *Client) TradePaymentScanpayClose(ctx context.Context, req *TradePaymentScanpayCloseRequest) (resp *TradePaymentScanpayCloseResponse, err error) {
	request := newRequest(`/v2/trade/payment/scanpay/close`, req)
	resp = &TradePaymentScanpayCloseResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradePaymentScanpayCloseRequest struct {
	ReqDate     string `json:"req_date"`
	ReqSeqId    string `json:"req_seq_id"`
	HuifuId     string `json:"huifu_id"`
	OrgReqDate  string `json:"org_req_date"`   //原交易请求日期 格式：yyyyMMdd
	OrgReqSeqId string `json:"org_req_seq_id"` //原交易请求流水号
}

type TradePaymentScanpayCloseResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		response.DataHeader
		HuifuId      string `json:"huifu_id"`
		ReqDate      string `json:"req_date"`
		ReqSeqId     string `json:"req_seq_id"`
		OrgReqDate   string `json:"org_req_date"`
		OrgReqSeqId  string `json:"org_req_seq_id"`
		OrgHfSeqId   string `json:"org_hf_seq_id"`
		OrgTransStat string `json:"org_trans_stat"`
		TransStat    string `json:"trans_stat"`
	} `json:"data"`
}
