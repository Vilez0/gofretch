package sys

import (
	"bufio"
	"gofretch/util/other"
	"os"
	"strings"
)

func Resolution() string {
	resolution := "Unknown"
	if util.FileExists("/sys/class/drm") {
		entries, err := os.ReadDir("/sys/class/drm")
		if err != nil {
			return "Unknown"
		}
		for _, entry := range entries {
			if util.FileExists("/sys/class/drm/" + entry.Name() + "/modes") {
				file, _ := os.Open("/sys/class/drm/" + entry.Name() + "/modes")
				scanner := bufio.NewScanner(file)
				defer file.Close()
				for scanner.Scan() {
					resolution = strings.TrimSuffix(scanner.Text(), "\n")
				}
			}
		}
	}
	if resolution == "" {
		return "Unknown"
	}
	return resolution
}
