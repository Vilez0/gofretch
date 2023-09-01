package mem

import (
	"strconv"
	"strings"

	util "github.com/Vilez0/gofretch/util/other"
)

func available() int {
	// This will read `/proc/meminfo` file and return the line that contains `MemAvailable:`
	line, err := util.ReadProcFile("/proc/meminfo", "MemAvailable:")
	if err != nil {
		return 0
	}
	// the data in /proc/meminfo file is stored as kb, this convert it to MiB
	memAvailableKB, _ := strconv.Atoi(strings.Fields(line[0])[1])
	memavailable := memAvailableKB / 1024
	return memavailable
}

func Total() int {
	// This will read `/proc/meminfo` file and return the line that contains `MemTotal:`
	line, err := util.ReadProcFile("/proc/meminfo", "MemTotal:")
	if err != nil {
		return 0
	}
	// the data in /proc/meminfo file is stored as KB, this convert it to MiB
	memTotalKB, _ := strconv.Atoi(strings.Fields(line[0])[1])
	memtotal := memTotalKB / 1024
	return memtotal
}

// Memory Usage in MiB
func UsageMB() int {
	return Total() - available()
}

// Memory usage percent
func UsagePercent() int {
	return int((float64(UsageMB()) / float64(Total())) * 100)
}
