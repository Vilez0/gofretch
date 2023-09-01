package util

import "os"
// FileExists checks if a file exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	} else {
		return true
	}
}
