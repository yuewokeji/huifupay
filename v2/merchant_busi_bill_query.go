package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// MerchantBusiBillQuery 交易结算对账单配置查询
// https://paas.huifu.com/partners/api#/jyjs/api_jyjs_wjpzcx
// 最近更新时间：2023.5.24
func (c *Client) MerchantBusiBillQuery(ctx context.Context, req *MerchantBusiBillQueryRequest) (resp *MerchantBusiBillQueryResponse, err error) {
	request := newRequest(`/v2/merchant/busi/bill/query`, req)
	resp = &MerchantBusiBillQueryResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type MerchantBusiBillQueryRequest struct {
	ReqDate  string `json:"req_date"`
	ReqSeqId string `json:"req_seq_id"`
	HuifuId  string `json:"huifu_id"`
}

type MerchantBusiBillQueryResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		RespCode        string `json:"resp_code"`
		RespDesc        string `json:"resp_desc"`
		HuifuId         string `json:"huifu_id"`
		Nature          string `json:"nature"`
		ReconSendFlag   string `json:"recon_send_flag"`
		FtpAddr         string `json:"ftp_addr"`
		FtpUser         string `json:"ftp_user"`
		FileType        string `json:"file_type"`
		AsSettleFlag    string `json:"as_settle_flag"`
		NotifyUrl       string `json:"notify_url"`
		ContainUser     string `json:"contain_user"`
		ContainMerchant string `json:"contain_merchant"`
	} `json:"data"`
}
