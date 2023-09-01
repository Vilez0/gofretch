package gpu

import (
	"os/exec"
	"strings"
)

func Name() string {
	var gpus string
	// this executes the lspci command and assign its output to cmd variable
	cmd, err := exec.Command("lspci").Output()
	if err != nil {
		// return Unknown if there an error
		return "Unknown"
	}
	// split the lspci output into lines
	lines := strings.Split(string(cmd), "\n")
	for _, e := range lines {
		if strings.Contains(e, "VGA") {
			/*
				this get the info we want, for example the line should look like this:
				00:02.0 VGA compatible controller: Intel Corporation 3rd Gen Core processor Graphics Controller (rev 09)
				first its trim the last 9 characters = ` (rev 09)`
				then its trim the first 35 characters = `00:02.0 VGA compatible controller: `
				then the output will returns `Intel Corporation 3rd Gen Core processor Graphics Controller`
			*/
			e := e[:len(e)-9][35:] + "\n"
			// append it to gpus because there can be more then one gpu in the host machine
			gpus += e
		}
	}
	// remove the `\n` in end of the `gpus`
	gpus = strings.TrimSuffix(gpus, "\n")
	return gpus
}
