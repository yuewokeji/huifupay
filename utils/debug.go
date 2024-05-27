package utils

import (
	"fmt"
	"os"
	"strings"
)

// NewDebug 返回一个debug对象
// 通过环境变量开启debug功能
// 变量名：HUIFU_DEBUG
// 不同模块按`,`隔开，all表示全部开启
// 支持模块：
// pay: v2包
// webhook: webhook包
func NewDebug(name string) *Debug {
	d := &Debug{
		isEnable: false,
		name:     name,
	}
	d.init()
	return d
}

type Debug struct {
	isEnable bool
	name     string
}

func (d *Debug) IsEnable() bool {
	return d.isEnable
}

func (d *Debug) init() {
	e := os.Getenv("HUIFU_DEBUG")
	e = strings.ToLower(e)
	e = strings.TrimSpace(e)
	if e == "all" {
		d.isEnable = true
		return
	}

	e = strings.ReplaceAll(e, " ", ",")
	parts := strings.Split(e, ",")
	for _, v := range parts {
		v = strings.TrimSpace(v)
		if v == d.name {
			d.isEnable = true
			return
		}
	}
}

func (d *Debug) Printf(format string, args ...interface{}) {
	if d.isEnable {
		fmt.Printf("[HuiFuPay] "+format+"\n", args...)
	}
}
