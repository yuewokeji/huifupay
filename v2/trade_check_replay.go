package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// TradeCheckReplay 交易结算对账文件重新生成
// https://paas.huifu.com/partners/api/#/jyjs/api_jyjs_wjbsc
// 最近更新时间：2023.5.11
func (c *Client) TradeCheckReplay(ctx context.Context, req *TradeCheckReplayRequest) (resp *TradeCheckReplayResponse, err error) {
	request := newRequest(`/v2/trade/check/replay`, req)
	resp = &TradeCheckReplayResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradeCheckReplayRequest struct {
	ReqSeqId string `json:"req_seq_id"`
	ReqDate  string `json:"req_date"`
	HuifuId  string `json:"huifu_id"`
	FileType string `json:"file_type"`
}

type TradeCheckReplayResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		RespCode    string `json:"resp_code"`
		RespDesc    string `json:"resp_desc"`
		ReqSeqId    string `json:"req_seq_id"`
		ReqDate     string `json:"req_date"`
		HuifuId     string `json:"huifu_id"`
		DownloadUrl string `json:"download_url"`
		FileType    string `json:"file_type"`
	} `json:"data"`
}
