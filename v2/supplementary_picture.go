package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/request"
	"github.com/yuewokeji/huifupay/response"
)

// SupplementaryPicture 图片上传
// https://paas.huifu.com/partners/api#/shgl/shjj/api_shjj_shtpsc?id=%e5%9b%be%e7%89%87%e4%b8%8a%e4%bc%a0
// 最近更新时间：2024.3.1
func (c *Client) SupplementaryPicture(ctx context.Context, filePath string, req *SupplementaryPictureRequest) (resp *SupplementaryPictureResponse, err error) {
	// 官方暂不支持通过file_url上传
	req.FileURL = ""

	request := request.NewFileRequest(requestURL(`/v2/supplementary/picture`), req)
	resp = &SupplementaryPictureResponse{
		BaseResponse: &response.BaseResponse{},
	}
	err = c.DoUploadFile(ctx, filePath, request, resp)
	return
}

type SupplementaryPictureRequest struct {
	ReqSeqID string `json:"req_seq_id"` // 业务请求流水号，业务请求流水号；示例值：2022012614120615001
	ReqDate  string `json:"req_date"`   // 业务请求日期，日期格式:yyyyMMdd;示例值：20220915

	// 图片类型，具体取值请参考各接口文档中的字段说明。枚举取值表；示例值：F01
	// https://paas.huifu.com/partners/api/#/csfl/api_csfl_wjlx
	FileType string `json:"file_type"`

	HuifuID string `json:"huifu_id"` // 商户号，渠道与一级代理商的直属商户ID；示例值：6666000123123123 如果商户未开户没有商户号，可以为空。不支持"企业用户基本信息开户"和"个人用户基本信息开户"接口所返回的用户号；
	FileURL string `json:"file_url"` // 文件url链接，文件url链接与file文件流不能同时上传；文件支持类型：JPG,BMP,PNG；单个图片最大支持2M；除去图片外其他类型文件支持10M 示例值：https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/imgs/guide/jinjian/%E8%90%A5%E4%B8%9A%E6%89%A7%E7%85%A71-2.png
}

type SupplementaryPictureResponse struct {
	*response.BaseResponse
	Data struct {
		RespCode string `json:"resp_code"` // resp_code 字段
		RespDesc string `json:"resp_desc"` // resp_desc 字段
		FileID   string `json:"file_id"`   // file_id 字段
	} `json:"data"`
}
