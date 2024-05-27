package huifupay

import "github.com/yuewokeji/huifupay/utils"

var debug *utils.Debug

func init() {
	debug = utils.NewDebug("pay")
}
