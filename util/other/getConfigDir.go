package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetConfigDir() string {
	// get the config directory from the environment variable
	configDir := os.Getenv("XDG_CONFIG_HOME")

	// if the environment variable is not set, use the default config directory
	if configDir == "" {
		configDir = os.Getenv("HOME") + "/.config"

		// if the `HOME` env is not set:
		if configDir == "/.config" {
			output, _ := exec.Command("whoami").Output()
			username := strings.TrimSuffix(string(output), "\n")
			configDir = "/home/" + username + "/.config"
		}
	}
	// append the program name to the config directory
	configDir += "/gofretch"
	
	// if the config directory does not exist, copy the default config files
	if FileExists(configDir) == false {
		rootConfigDirs := []string{"/etc/gofretch", "/usr/local/etc/gofretch", "/usr/share/gofretch", "/usr/local/share/gofretch"}
		var rootConfigDir string
		for _, e := range rootConfigDirs {
			if FileExists(e) {
				rootConfigDir = e
				break
			}
		}

		if rootConfigDir == "" {
			fmt.Printf("Cannot find config directory in: %v", rootConfigDirs)
		}

		err := exec.Command("cp", "-r", rootConfigDir, configDir).Run()
		if err != nil {
			println("Cannot copy config files from "+rootConfigDir+" to "+configDir+"/gofretch: ", err)
			println("Using default config files")
			configDir = rootConfigDir
		}
	}
	return configDir
}
