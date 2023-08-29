package osinfo

import (
	"bufio"
	"os"
	"strings"
)

func parseOsrelease() []string {
	var lines []string
	file, err := os.Open("/etc/os-release")
	if err != nil {
		file, err = os.Open("/usr/lib/os-release")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func DistroName() string {
	osRelease := parseOsrelease()
	for e := range osRelease {
		if strings.Contains(osRelease[e], "NAME") {
			name := strings.Split(osRelease[e], "=")[1]
			name = strings.ReplaceAll(name, `"`, ``)
			return name
		}
	}
	return ""
}
