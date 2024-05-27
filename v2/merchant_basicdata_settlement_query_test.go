package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_MerchantBasicdataSettlementQuery(t *testing.T) {
	req := &MerchantBasicdataSettlementQueryRequest{
		ReqSeqId: testReqSeqID(),
		ReqDate:  testReqDate(),
		HuifuId:  testHuifuID,
	}
	_, err := testClient.MerchantBasicdataSettlementQuery(context.Background(), req)
	assert.Nil(t, err)
}
