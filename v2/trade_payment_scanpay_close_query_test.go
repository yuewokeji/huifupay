package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_TradePaymentScanPayCloseQuery(t *testing.T) {
	req := &TradePaymentScanPayCloseQueryRequest{
		ReqDate:  testReqDate(),
		ReqSeqId: testReqSeqID(),
		HuifuId:  testHuifuID,
	}
	_, err := testClient.TradePaymentScanPayCloseQuery(context.Background(), req)
	assert.Nil(t, err)
}
