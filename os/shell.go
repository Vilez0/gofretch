package osinfo

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func Shell() string {
	reShell := regexp.MustCompile(`bash|zsh|fish|fish|csh|tcsh|ksh`)
	reShellVersion := regexp.MustCompile(`\d+(?:\.\d+){1,}`)
	shellPath := strings.Split(os.Getenv("SHELL"), "/")
	shellName := shellPath[len(shellPath)-1]
	output, err := exec.Command(shellName, "--version").CombinedOutput()
	if err != nil {
		panic(err)
	}
	strOutput := strings.TrimSuffix(string(output), "\n")

	var shell, version string
	if reShell.MatchString(strOutput) {
		shell = reShell.FindString(strOutput)
	}
	if reShellVersion.MatchString(strOutput) {
		version = reShellVersion.FindString(strOutput)
	}
	if shell == "" && version == "" {
		return "Unknown"
	}
	return shell + " " + version
}
