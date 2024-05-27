package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_TradePaymentScanpayQuery(t *testing.T) {
	req := &TradePaymentScanpayQueryRequest{
		HuifuID:     testHuifuID,
		OrgReqDate:  testReqDate(),
		OrgReqSeqID: testReqSeqID(),
		OutOrderID:  "88888888",
	}
	_, err := testClient.TradePaymentScanpayQuery(context.Background(), req)
	assert.Nil(t, err)
}
