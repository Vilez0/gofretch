package osinfo

import (
	"os/exec"
)

func KernelName() string {
	//returns kernel name according to `uname -r` command output
	output, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return "Unknown"
	}
	//remove \n from output
	output = output[:len(output)-1]
	return string(output)
}
