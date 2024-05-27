package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_TradeSettlementEnchashment(t *testing.T) {
	req := &TradeSettlementEnchashmentRequest{
		ReqDate:  testReqDate(),
		ReqSeqId: testReqSeqID(),
		HuifuId:  testHuifuID,
	}
	_, err := testClient.TradeSettlementEnchashment(context.Background(), req)
	assert.Nil(t, err)
}
