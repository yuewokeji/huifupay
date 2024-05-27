package v2

import (
	"bytes"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yuewokeji/huifupay/notify"
	"io"
	"net/http"
	"testing"
)

func TestClient_TradePaymentJspay(t *testing.T) {
	req := &TradePaymentJspayRequest{
		ReqDate:  testReqDate(),
		ReqSeqId: testReqSeqID(),
		HuifuId:  testHuifuID,
	}
	_, err := testClient.TradePaymentJspay(context.Background(), req)
	assert.Nil(t, err)
}

func TestClient_TradePaymentJspayNotify(t *testing.T) {
	body := []byte(`resp_desc=%E4%BA%A4%E6%98%93%E6%88%90%E5%8A%9F%5B000%5D&resp_code=00000000&sign=V489pKyVLZByfUpGq3bt8KWsUDZRvrsjPih262kTC9dq8J2CiSRwWvcztIkNISJJddoQeTGe72AnU6YlcxPjxQ%2F4b1UHMWV3QP9sY6U8eMq99VHTkT%2FoA1mraPOQp%2BHlQ%2FmqZrLO%2BUDAkefrz36F8dRYdAm30buVgLkpsQKl4YvGAaZlMRZFxxp2J%2BKuv29w5Wkn2mZYst2mKQ5%2BjWZ5MXwL6DSPmNGolJ2bszjl9XsBUjinYd0fHNDAcDfATLfSBlNonl6%2BFRtVq%2BBk4JutzB2vQ0xoQvgmWEHMSuQNKV5ylkOWyZkvOs3F%2BCWjtJUH6IAMkhNMFVZFYzAabB70sQ%3D%3D&resp_data=%7B%22acct_date%22%3A%2220240508%22%2C%22acct_id%22%3A%22A33685621%22%2C%22acct_split_bunch%22%3A%7B%22acct_infos%22%3A%5B%7B%22acct_date%22%3A%2220240508%22%2C%22acct_id%22%3A%22A33685621%22%2C%22div_amt%22%3A%220.01%22%2C%22huifu_id%22%3A%226666000150066732%22%7D%5D%2C%22fee_acct_date%22%3A%2220240508%22%2C%22fee_acct_id%22%3A%22A33750398%22%2C%22fee_amt%22%3A%220.00%22%2C%22fee_huifu_id%22%3A%226666000149982404%22%7D%2C%22acct_stat%22%3A%22S%22%2C%22alipay_response%22%3A%7B%22app_id%22%3A%22%22%2C%22buyer_id%22%3A%222088112018986187%22%2C%22buyer_logon_id%22%3A%22188****9983%22%2C%22coupon_fee%22%3A%220.00%22%2C%22fund_bill_list%22%3A%5B%7B%22amount%22%3A%220.01%22%2C%22fund_channel%22%3A%22ALIPAYACCOUNT%22%7D%5D%7D%2C%22atu_sub_mer_id%22%3A%222088440552679019%22%2C%22avoid_sms_flag%22%3A%22%22%2C%22bagent_id%22%3A%226666000149982404%22%2C%22bank_code%22%3A%22TRADE_SUCCESS%22%2C%22bank_desc%22%3A%22TRADE_SUCCESS%22%2C%22bank_message%22%3A%22TRADE_SUCCESS%22%2C%22bank_order_no%22%3A%22282024050822001486181416924955%22%2C%22bank_seq_id%22%3A%22838754%22%2C%22base_acct_id%22%3A%22A33685621%22%2C%22batch_id%22%3A%22240508%22%2C%22channel_type%22%3A%22U%22%2C%22combinedpay_data%22%3A%5B%5D%2C%22combinedpay_fee_amt%22%3A%220.00%22%2C%22debit_type%22%3A%220%22%2C%22delay_acct_flag%22%3A%22N%22%2C%22div_flag%22%3A%220%22%2C%22end_time%22%3A%2220240508161447%22%2C%22fee_amount%22%3A%220.00%22%2C%22fee_amt%22%3A%220.00%22%2C%22fee_flag%22%3A1%2C%22fee_formula_infos%22%3A%5B%7B%22fee_formula%22%3A%22MAX%280.00%2CAMT*0.0025%29%22%2C%22fee_type%22%3A%22TRANS_FEE%22%7D%5D%2C%22fee_rec_type%22%3A%222%22%2C%22fee_type%22%3A%22OUTSIDE%22%2C%22gate_id%22%3A%22Dx%22%2C%22hf_seq_id%22%3A%22002900TOP4A240508161421P773ac139c9100000%22%2C%22huifu_id%22%3A%226666000150066732%22%2C%22is_delay_acct%22%3A%220%22%2C%22is_div%22%3A%220%22%2C%22maze_resp_code%22%3A%22%22%2C%22mer_name%22%3A%22%E5%B9%BF%E5%B7%9E%E5%B8%82%E4%B9%90%E7%AA%9D%E7%8E%A9%E5%88%9B%E9%9F%B3%E4%B9%90%E6%96%87%E5%8C%96%E6%9C%89%E9%99%90%E5%85%AC%E5%8F%B8%22%2C%22mer_ord_id%22%3A%22PA20240508161421838754%22%2C%22mypaytsf_discount%22%3A%220.00%22%2C%22need_big_object%22%3Afalse%2C%22notify_type%22%3A2%2C%22org_auth_no%22%3A%22%22%2C%22org_huifu_seq_id%22%3A%22%22%2C%22org_trans_date%22%3A%22%22%2C%22out_ord_id%22%3A%22282024050822001486181416924955%22%2C%22out_trans_id%22%3A%22282024050822001486181416924955%22%2C%22party_order_id%22%3A%2203242405085846194200441%22%2C%22pay_amt%22%3A%220.01%22%2C%22pay_scene%22%3A%2202%22%2C%22posp_seq_id%22%3A%2203242405085846194200441%22%2C%22product_id%22%3A%22PAYUN%22%2C%22ref_no%22%3A%22161421838754%22%2C%22remark%22%3A%22%22%2C%22req_date%22%3A%2220240508%22%2C%22req_seq_id%22%3A%22PA20240508161421838754%22%2C%22resp_code%22%3A%2200000000%22%2C%22resp_desc%22%3A%22%E4%BA%A4%E6%98%93%E6%88%90%E5%8A%9F%22%2C%22risk_check_data%22%3A%7B%7D%2C%22risk_check_info%22%3A%7B%7D%2C%22settlement_amt%22%3A%220.01%22%2C%22sub_resp_code%22%3A%2200000000%22%2C%22sub_resp_desc%22%3A%22%E4%BA%A4%E6%98%93%E6%88%90%E5%8A%9F%22%2C%22subsidy_stat%22%3A%22I%22%2C%22sys_id%22%3A%226666000149982404%22%2C%22trade_type%22%3A%22A_NATIVE%22%2C%22trans_amt%22%3A%220.01%22%2C%22trans_date%22%3A%2220240508%22%2C%22trans_fee_allowance_info%22%3A%7B%22actual_fee_amt%22%3A%220.00%22%2C%22allowance_fee_amt%22%3A%220.00%22%2C%22allowance_type%22%3A%220%22%2C%22receivable_fee_amt%22%3A%220.00%22%7D%2C%22trans_stat%22%3A%22S%22%2C%22trans_time%22%3A%22161421%22%2C%22trans_type%22%3A%22A_NATIVE%22%2C%22ts_encash_detail%22%3A%5B%5D%7D`)
	req, _ := http.NewRequest("POST", "https://callback.musicwow.com/callback", io.NopCloser(bytes.NewBuffer(body)))

	n := &TradePaymentJspayNotify{
		Notify: notify.NewBaseNotify(),
	}
	err := testClient.DoNotifyRequest(context.Background(), req, n)
	fmt.Println(err)
	fmt.Println(n)
}
