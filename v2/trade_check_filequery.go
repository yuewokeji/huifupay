package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// TradeCheckFilequery 交易结算对账单查询
// https://paas.huifu.com/partners/api#/jyjs/api_jyjs_wjcx
// 最近更新时间：2023.5.24
func (c *Client) TradeCheckFilequery(ctx context.Context, req *TradeCheckFilequeryRequest) (resp *TradeCheckFilequeryResponse, err error) {
	request := newRequest(`/v2/trade/check/filequery`, req)
	resp = &TradeCheckFilequeryResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradeCheckFilequeryRequest struct {
	ReqDate       string `json:"req_date"`
	ReqSeqId      string `json:"req_seq_id"`
	HuifuId       string `json:"huifu_id"`
	FileDate      string `json:"file_date"`
	FileTypeQuery string `json:"file_type_query"`
}

type TradeCheckFilequeryFileDetail struct {
	HuifuId       string `json:"huifu_id"`
	FileType      string `json:"file_type"`
	FileId        string `json:"file_id"`
	FileName      string `json:"file_name"`
	DownloadUrl   string `json:"download_url"`
	DataDate      string `json:"data_date"`
	TaskStat      string `json:"task_stat"`
	TaskStartTime string `json:"task_start_time"`
	TaskEndTime   string `json:"task_end_time"`
}

type TradeCheckFilequeryTaskDetail struct {
	HuifuId       string `json:"huifu_id"`
	FileType      string `json:"file_type"`
	FileId        string `json:"file_id"`
	FileName      string `json:"file_name"`
	DownloadUrl   string `json:"download_url"`
	DataDate      string `json:"data_date"`
	TaskStat      string `json:"task_stat"`
	TaskStartTime string `json:"task_start_time"`
	TaskEndTime   string `json:"task_end_time"`
}

type TradeCheckFilequeryResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		RespCode    string                          `json:"resp_code"`
		RespDesc    string                          `json:"resp_desc"`
		ReqDate     string                          `json:"req_date"`
		ReqSeqId    string                          `json:"req_seq_id"`
		FileDetails []TradeCheckFilequeryFileDetail `json:"file_details"`
		TaskDetails []TradeCheckFilequeryTaskDetail `json:"task_details"`
	} `json:"data"`
}
