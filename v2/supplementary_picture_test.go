package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/yuewokeji/huifupay/response"
	"testing"
)

func TestClient_SupplementaryPicture(t *testing.T) {
	req := &SupplementaryPictureRequest{
		ReqSeqID: testReqSeqID(),
		ReqDate:  testReqDate(),
		FileType: "F25",
		HuifuID:  testHuifuID,
		FileURL:  "",
	}
	resp := &SupplementaryPictureResponse{
		BaseResponse: &response.BaseResponse{},
	}
	resp, err := testClient.SupplementaryPicture(context.Background(), "supplementary_picture_test.png", req)
	assert.Nil(t, err)
	assert.Equal(t, respCodeSuccess, resp.Data.RespCode)
	assert.NotEmpty(t, resp.Data.FileID)
}
