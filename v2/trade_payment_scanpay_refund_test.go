package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_TradePaymentScanpayRefund(t *testing.T) {
	req := &TradePaymentScanpayRefundRequest{
		ReqDate:  testReqDate(),
		ReqSeqId: "RF20240524105250312968",
		HuifuId:  "6666000151139507",
	}
	_, err := testClient.TradePaymentScanpayRefund(context.Background(), req)
	assert.Nil(t, err)
}
