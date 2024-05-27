package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_TradeAcctpaymentBalanceQuery(t *testing.T) {
	req := &TradeAcctpaymentBalanceQueryRequest{
		ReqDate:  testReqDate(),
		ReqSeqID: testReqSeqID(),
		HuifuID:  testHuifuID,
	}
	resp, err := testClient.TradeAcctpaymentBalanceQuery(context.Background(), req)
	assert.Nil(t, err)
	assert.Positive(t, len(resp.Data.AcctInfoListObject))
}
