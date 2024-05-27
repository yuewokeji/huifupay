# HuifuPay

[汇付支付](https://paas.huifu.com/partners/api#/)的Golang版本SDK。

## 安装

```bash
go get -u github.com/yuewokeji/huifupay
```

## 使用

### 初始化

```go
config := huifupay.NewConfig("system_id", "product_id", "rsa_huifu_pub_key", "rsa_merch_pri_key", true)
client := v2.NewClient(config)
```

使用自定义的http Client，调用`huifupay.SetGlobalHTTPClientFunc()`（全局配置），或者`huifupay.WithHTTPClientFunc()`（只对当前实例有效）


### 同步接口

```go
package main

import (
	"context"
	"fmt"
	"github.com/yuewokeji/huifupay"
	"github.com/yuewokeji/huifupay/v2"
)

func main() {
	config := huifupay.NewConfig("system_id", "product_id", "rsa_huifu_pub_key", "rsa_merch_pri_key", true)
	client := v2.NewClient(config)
	req := &v2.TradeAcctpaymentBalanceQueryRequest{
		// ...
	}
	resp, err := client.TradeAcctpaymentBalanceQuery(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Data.RespCode)
}

```

由于时间关系，只实现了v2版本的部分api（详见[api列表](v2/)），有缘人请提交PR。

### 异步通知

```go
package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuewokeji/huifupay"
	"github.com/yuewokeji/huifupay/v2"
)

func main() {
	config := huifupay.NewConfig("system_id", "product_id", "rsa_huifu_pub_key", "rsa_merch_pri_key", true)
	client := v2.NewClient(config)
	
	r := gin.Default()
	r.POST("/refund/notify", func(c *gin.Context) {
		n, err := client.TradePaymentScanpayRefundNotify(context.Background(), c.Request)
		if err != nil {
			panic(err)
		}
		fmt.Println(n.RespCode)
	})
	r.Run()
}

```

### webhook

```go
package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuewokeji/huifupay/webhook"
)

func main() {
	hook := webhook.NewWithSigner([]byte("your_access_key"))

	// 注册默认处理方法
	hook.SetDefaultHandler(func(ctx context.Context, data []byte) error {
		fmt.Printf("web hook: %s\n", string(data))
		return nil
	})

	// 注册其它处理方法
	// hook.Register(e, fn)

	r := gin.Default()
	r.POST("/webhook", func(c *gin.Context) {
		hr, err := hook.HandleRequest(context.Background(), c.Request)
		if err != nil {
			fmt.Printf("handle webhook error: %s\n", err.Error())
		} else {
			fmt.Printf("handle webhook: %s", hr.Event)
		}
	})
	r.Run()
}

```

### json字符串与结构体互转

汇付接口中，有部分参数是json对象字符串，并不是json对象。

[autoassign](autoassign/autoassign.go)支持这两个不同类型的字段自动转换。

[交易查询接口](https://paas.huifu.com/partners/api#/smzf/api_qrpay_cx?id=ewm)中，返回参数fee_formula_infos是一个string类型的json对象。

通过tag配置可以实现自动转换：

```go
type TradePaymentScanPayQueryResponse struct {
    // ...
    FeeFormulaInfos       string                                   `json:"fee_formula_infos"` //手续费费率信息 交易成功时返回手续费费率信息
    FeeFormulaInfosObject []TradePaymentScanpayQueryFeeFormulaInfo `autoassign:"FeeFormulaInfos"`
}
```

### debug

配置环境变量，打开debug功能。

**变量名**：HUIFU_DEBUG

**变量值**：参照[debug说明](utils/debug.go)
