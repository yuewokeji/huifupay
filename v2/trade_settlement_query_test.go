package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_TradeSettlementQuery(t *testing.T) {
	req := &TradeSettlementQueryRequest{
		OrgReqDate:  testReqDate(),
		OrgReqSeqId: testReqSeqID(),
		HuifuId:     testHuifuID,
	}
	_, err := testClient.TradeSettlementQuery(context.Background(), req)
	assert.Nil(t, err)
}
