package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// MerchantBasicdataSettlementQuery 结算记录查询
// https://paas.huifu.com/partners/api#/jyjs/qx/api_jsjlfycx
// 最近更新时间：2023.4.30
func (c *Client) MerchantBasicdataSettlementQuery(ctx context.Context, req *MerchantBasicdataSettlementQueryRequest) (resp *MerchantBasicdataSettlementQueryResponse, err error) {
	request := newRequest(`/v2/merchant/basicdata/settlement/query`, req)
	resp = &MerchantBasicdataSettlementQueryResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type MerchantBasicdataSettlementQueryRequest struct {
	ReqSeqId    string `json:"req_seq_id"`
	ReqDate     string `json:"req_date"`
	HuifuId     string `json:"huifu_id"`
	BeginDate   string `json:"begin_date"`
	EndDate     string `json:"end_date"`
	PageSize    string `json:"page_size"`
	SettleCycle string `json:"settle_cycle"`
	PageNum     string `json:"page_num"`
	TransStat   string `json:"trans_stat"`
	SortColumn  string `json:"sort_column"`
	SortOrder   string `json:"sort_order"`
}

type MerchantBasicdataSettlementQueryResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		RespCode           string `json:"resp_code"`
		RespDesc           string `json:"resp_desc"`
		ResultCount        int    `json:"result_count"`
		PageSize           int    `json:"page_size"`
		PageNum            int    `json:"page_num"`
		TransLogResultList string `json:"trans_log_result_list"`
	} `json:"data"`
}
