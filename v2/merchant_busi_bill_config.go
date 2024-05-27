package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/response"
)

// MerchantBusiBillConfig 交易结算对账文件配置
// https://paas.huifu.com/partners/api#/jyjs/api_jyjs_wjpz
// 最近更新时间：2023.5.24
func (c *Client) MerchantBusiBillConfig(ctx context.Context, req *MerchantBusiBillConfigRequest) (resp *MerchantBusiBillConfigResponse, err error) {
	request := newRequest(`/v2/merchant/busi/bill/config`, req)
	resp = &MerchantBusiBillConfigResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type MerchantBusiBillConfigRequest struct {
	ReqDate         string `json:"req_date"`
	ReqSeqId        string `json:"req_seq_id"`
	HuifuId         string `json:"huifu_id"`
	ReconSendFlag   string `json:"recon_send_flag"`
	FtpAddr         string `json:"ftp_addr"`
	FtpUser         string `json:"ftp_user"`
	FtpPwd          string `json:"ftp_pwd"`
	FileType        string `json:"file_type"`
	AsSettleFlag    string `json:"as_settle_flag"`
	NotifyUrl       string `json:"notify_url"`
	ContainUser     string `json:"contain_user"`
	ContainMerchant string `json:"contain_merchant"`
}

type MerchantBusiBillConfigResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		RespCode        string `json:"resp_code"`
		RespDesc        string `json:"resp_desc"`
		ReqSeqId        string `json:"req_seq_id"`
		ReqDate         string `json:"req_date"`
		HuifuId         string `json:"huifu_id"`
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
