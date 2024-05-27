package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_TradePaymentScanpayRefundQuery(t *testing.T) {
	req := &TradePaymentScanpayRefundQueryRequest{
		OrgReqDate:  testReqDate(),
		OrgReqSeqId: testReqSeqID(),
		HuifuId:     testHuifuID,
	}
	_, err := testClient.TradePaymentScanpayRefundQuery(context.Background(), req)
	assert.Nil(t, err)
}
