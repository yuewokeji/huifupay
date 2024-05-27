package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_TradeAcctpaymentAcctlogQuery(t *testing.T) {
	req := &TradeAcctpaymentAcctlogQueryRequest{
		ReqSeqId: testReqSeqID(),
		HuifuId:  testHuifuID,
		AcctDate: testReqDate(),
	}
	_, err := testClient.TradeAcctpaymentAcctlogQuery(context.Background(), req)
	assert.Nil(t, err)
}
