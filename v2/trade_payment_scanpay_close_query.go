package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// TradePaymentScanPayCloseQuery 扫码交易关单查询
// https://paas.huifu.com/partners/api#/smzf/api_jygdcx
// 最近更新时间：2023.3.31
func (c *Client) TradePaymentScanPayCloseQuery(ctx context.Context, req *TradePaymentScanPayCloseQueryRequest) (resp *TradePaymentScanPayCloseQueryResponse, err error) {
	request := newRequest(`/v2/trade/payment/scanpay/closequery`, req)
	resp = &TradePaymentScanPayCloseQueryResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradePaymentScanPayCloseQueryRequest struct {
	ReqDate     string `json:"req_date"`
	ReqSeqId    string `json:"req_seq_id"`
	HuifuId     string `json:"huifu_id"`
	OrgReqDate  string `json:"org_req_date"`
	OrgReqSeqId string `json:"org_req_seq_id"`
}

type TradePaymentScanPayCloseQueryResponse struct {
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
