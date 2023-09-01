package osinfo

import "os"

// Simply returns the hostname
func Hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "Unknown"
	}
	return hostname
}
