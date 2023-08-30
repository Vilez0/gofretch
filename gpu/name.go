package gpu

import (
	"os/exec"
	"strings"
)

func Name() string {
	var gpus string
	cmd, err := exec.Command("lspci").Output()
	if err != nil {
		return "Unknown"
	}
	out := strings.Split(string(cmd), "\n")
	for _, e := range out {
		if strings.Contains(e, "VGA") {
			e := e[:len(e)-9][35:] + "\n"
			gpus += e
		}
	}
	gpus = strings.TrimSuffix(gpus, "\n")
	return gpus
}
