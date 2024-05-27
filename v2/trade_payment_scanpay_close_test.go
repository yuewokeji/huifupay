package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_TradePaymentScanPayClose(t *testing.T) {
	req := &TradePaymentScanpayCloseRequest{
		ReqDate:  testReqDate(),
		ReqSeqId: testReqSeqID(),
		HuifuId:  testHuifuID,
	}
	_, err := testClient.TradePaymentScanpayClose(context.Background(), req)
	assert.Nil(t, err)
}
