package osinfo

import (
	"os/exec"
)

func KernelName() string {
	output, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return "Unknown"
	}
	//remove \n from output
	output = output[:len(output)-1]
	return string(output)
}
