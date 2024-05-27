package webhook

import "github.com/yuewokeji/huifupay/utils"

var debug *utils.Debug

func init() {
	debug = utils.NewDebug("webhook")
}
