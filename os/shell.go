package osinfo

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	shell   string
	version string
)

// Shell returns the current shell name and version
func Shell() string {
	// regex to match the shell name
	reShell := regexp.MustCompile(`bash|zsh|fish|fish|csh|tcsh|ksh`)
	// regex to match the shell version
	reShellVersion := regexp.MustCompile(`\d+(?:\.\d+){1,}`)
	// get the shell name
	shellPath := strings.Split(os.Getenv("SHELL"), "/")
	// get the shell version
	shellName := shellPath[len(shellPath)-1]
	output, err := exec.Command(shellName, "--version").CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	// remove the trailing newline
	strOutput := strings.TrimSuffix(string(output), "\n")

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
