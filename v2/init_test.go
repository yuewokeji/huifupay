package v2

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/yuewokeji/huifupay"
	"strconv"
	"time"
)

//go:embed config_test.json
var testConfigJSON []byte
var testConfig huifupay.Config
var testHuifuID string
var respCodeSuccess = "00000000"
var testClient *Client

func init() {
	fmt.Printf("test config: %s\n", string(testConfigJSON))

	v := struct {
		Config  huifupay.Config `json:"config"`
		HuifuID string          `json:"huifu_id"`
	}{}

	err := json.Unmarshal(testConfigJSON, &v)
	if err != nil {
		panic("unmarshal config error: " + err.Error())
	}
	testConfig = v.Config
	testHuifuID = v.HuifuID

	testClient = NewClient(testConfig)
}

func testReqDate() string {
	return time.Now().Format("20060102")
}

func testReqSeqID() string {
	i := time.Now().UnixMicro()
	return strconv.Itoa(int(i))
}
