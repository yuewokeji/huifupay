package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MerchantBasicdataQuery(t *testing.T) {
	req := &MerchantBasicdataQueryRequest{
		ReqDate:  testReqDate(),
		ReqSeqId: testReqSeqID(),
		HuifuId:  testHuifuID,
	}
	resp, err := testClient.MerchantBasicdataQuery(context.Background(), req)
	assert.Nil(t, err)
	assert.Equal(t, respCodeSuccess, resp.Data.RespCode)
}
