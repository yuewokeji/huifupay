package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// TradeSettlementQuery 出金交易查询接口
// https://paas.huifu.com/partners/api#/jyjs/qx/api_cjjycx
// 最近更新时间：2023.4.26
func (c *Client) TradeSettlementQuery(ctx context.Context, req *TradeSettlementQueryRequest) (resp *TradeSettlementQueryResponse, err error) {
	request := newRequest(`/v2/trade/settlement/query`, req)
	resp = &TradeSettlementQueryResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradeSettlementQueryRequest struct {
	HuifuId     string `json:"huifu_id"`
	OrgReqDate  string `json:"org_req_date"`
	OrgReqSeqId string `json:"org_req_seq_id,omitempty"`
	OrgHfSeqId  string `json:"org_hf_seq_id,omitempty"`
}

type TradeSettlementQueryResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		RespCode    string `json:"resp_code"`
		RespDesc    string `json:"resp_desc"`
		OrgHfSeqId  string `json:"org_hf_seq_id"`
		OrgReqDate  string `json:"org_req_date"`
		TransStatus string `json:"trans_status"`
		OrgReqSeqId string `json:"org_req_seq_id"`
		TransDesc   string `json:"trans_desc"`
		CashAmt     string `json:"cash_amt"`
		FeeAmt      string `json:"fee_amt"`
	} `json:"data"`
}
