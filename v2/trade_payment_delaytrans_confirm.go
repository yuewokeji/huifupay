package v2

import (
	"context"
	"github.com/yuewokeji/huifupay/notify"
	"github.com/yuewokeji/huifupay/response"
	"net/http"
)

// TradePaymentDelaytransConfirm 交易确认接口
// https://paas.huifu.com/partners/api#/smzf/api_jyqr
// 最近更新时间：2023.5.23
func (c *Client) TradePaymentDelaytransConfirm(ctx context.Context, req *TradePaymentDelaytransConfirmRequest) (resp *TradePaymentDelaytransConfirmResponse, err error) {
	request := newRequest(`/v2/trade/payment/delaytrans/confirm`, req)
	resp = &TradePaymentDelaytransConfirmResponse{
		BaseResponse: response.NewBaseResponse(),
	}
	err = c.DoRequest(ctx, request, resp)
	return
}

type TradePaymentDelaytransConfirmRequest struct {
	ReqDate              string                          `json:"req_date"`
	ReqSeqId             string                          `json:"req_seq_id"`
	HuifuId              string                          `json:"huifu_id"`
	OrgReqDate           string                          `json:"org_req_date"`
	OrgReqSeqId          string                          `json:"org_req_seq_id,omitempty"`
	OrgMerOrdId          string                          `json:"org_mer_ord_id,omitempty"`
	OrgHfSeqId           string                          `json:"org_hf_seq_id,omitempty"`
	AcctSplitBunch       string                          `json:"acct_split_bunch"`
	AcctSplitBunchObject TradePaymentJspayAcctSplitBunch `json:"-" autoassign:"AcctSplitBunch"`
	RiskCheckData        string                          `json:"risk_check_data"`
	PayType              string                          `json:"pay_type"`
	Remark               string                          `json:"remark"`
	HycFlag              string                          `json:"hyc_flag"`
	SalaryModleType      string                          `json:"salary_modle_type"`
	BmemberId            string                          `json:"bmember_id"`
	HycAttachId          string                          `json:"hyc_attach_id"`
	NotifyUrl            string                          `json:"notify_url"`
}

type TradePaymentDelaytransConfirmResponse struct {
	*response.BaseResponse
	response.Sign
	Data struct {
		RespCode        string `json:"resp_code"`
		RespDesc        string `json:"resp_desc"`
		HuifuId         string `json:"huifu_id"`
		ReqDate         string `json:"req_date"`
		ReqSeqId        string `json:"req_seq_id"`
		TransStat       string `json:"trans_stat"`
		HfSeqId         string `json:"hf_seq_id"`
		UnconfirmAmt    string `json:"unconfirm_amt"`
		ConfirmedAmt    string `json:"confirmed_amt"`
		AcctRespCode    string `json:"acct_resp_code"`
		AcctRespDesc    string `json:"acct_resp_desc"`
		OrgMerOrdId     string `json:"org_mer_ord_id"`
		OrgReqDate      string `json:"org_req_date"`
		OrgReqSeqId     string `json:"org_req_seq_id"`
		OrgHfSeqId      string `json:"org_hf_seq_id"`
		HycFlag         string `json:"hyc_flag"`
		HycAttachId     string `json:"hyc_attach_id"`
		SalaryModleType string `json:"salary_modle_type"`
		BmemberId       string `json:"bmember_id"`
	}
}

// TradePaymentDelaytransConfirmNotify 异步通知
func (c *Client) TradePaymentDelaytransConfirmNotify(ctx context.Context, req *http.Request) (n *TradePaymentDelaytransConfirmNotify, err error) {
	n = &TradePaymentDelaytransConfirmNotify{
		Notify: notify.NewBaseNotify(),
	}
	err = c.DoNotifyRequest(ctx, req, n)
	return
}

type TradePaymentDelaytransPayConfirmDetail struct {
	HuifuId    string `json:"huifu_id"`
	AcctId     string `json:"acct_id"`
	DivAmt     string `json:"div_amt"`
	FeeHuifuId string `json:"fee_huifu_id"`
	FeeAcctId  string `json:"fee_acct_id"`
	FeeAmt     string `json:"fee_amt"`
	TransStat  string `json:"trans_stat"`
}

// TradePaymentDelaytransConfirmNotify 异步通知
type TradePaymentDelaytransConfirmNotify struct {
	notify.Notify
	RespCode string `json:"resp_code"`
	RespDesc string `json:"resp_desc"`
	RespData struct {
		ReqSeqId          string                                   `json:"req_seq_id"`
		ReqDate           string                                   `json:"req_date"`
		HuifuId           string                                   `json:"huifu_id"`
		HfSeqId           string                                   `json:"hf_seq_id"`
		TransStat         string                                   `json:"trans_stat"`
		PayConfirmDetails []TradePaymentDelaytransPayConfirmDetail `json:"pay_confirm_details"`
	} `json:"resp_data"`
}
