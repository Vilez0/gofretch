package sys

import (
	"gofretch/util"
	"os/exec"
)

func ModelName() string {
	var model string
	if util.FileExists("/system/app/") && util.FileExists("/system/priv-app") {
		byteBrand, _ := exec.Command("getprop", "ro.product.brand").Output()
		byteModel, _ := exec.Command("getprop", "ro.product.model").Output()
		model = string(byteBrand) + " " + string(byteModel)
		return model
	} else if util.FileExists("/sys/devices/virtual/dmi/id/board_vendor") || util.FileExists("/sys/devices/virtual/dmi/id/board_name") {
		model = util.ScanFile("/sys/devices/virtual/dmi/id/board_vendor")[0]
		model += " " + util.ScanFile("/sys/devices/virtual/dmi/id/board_name")[0]
		return model
	} else if util.FileExists("/sys/devices/virtual/dmi/id/product_name") || util.FileExists("/sys/devices/virtual/dmi/id/product_version") {
		model = util.ScanFile("/sys/devices/virtual/dmi/id/product_name")[0]
		model += " " + util.ScanFile("/sys/devices/virtual/dmi/id/product_version")[0]
		return model
	} else if util.FileExists("/sys/firmware/devicetree/base/model") {
		model = util.ScanFile("/sys/firmware/devicetree/base/model")[0]
		return model
	} else if util.FileExists("/tmp/sysinfo/model") {
		model = util.ScanFile("/tmp/sysinfo/model")[0]
		return model
	}
	return model
}
