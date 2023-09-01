package osinfo

import (
	"strings"

	util "github.com/Vilez0/gofretch/util/other"
)

var (
	lines []string
)

// reads /etc/os-release file then return its lines in a slice
func parseOsrelease() []string {
	filename := "/etc/os-release"
	if !util.FileExists("/etc/os-release") {
		filename = "/usr/lib/os-release"
	}

	lines = util.ScanFile(filename)

	return lines
}

func DistroName() string {
	osRelease := parseOsrelease()
	// find the line is starts with `NAME`, parse it and return the distro name
	for _, e := range osRelease {
		if strings.HasPrefix(e, "NAME") {
			name := strings.Split(e, "=")[1]
			name = strings.ReplaceAll(name, `"`, ``)
			return name
		}
	}
	return ""
}
