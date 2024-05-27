package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_TradePaymentMicropay(t *testing.T) {
	req := &TradePaymentMicropayRequest{
		ReqDate:  testReqDate(),
		ReqSeqId: testReqSeqID(),
		HuifuId:  testHuifuID,
	}
	_, err := testClient.TradePaymentMicropay(context.Background(), req)
	assert.Nil(t, err)
}
