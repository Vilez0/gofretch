package mem

import (
	"log"
	"strconv"
	"strings"

	util "gofretch/util/other"
)

func available() int {
	line, err := util.ReadProcFile("/proc/meminfo", "MemAvailable:")
	if err != nil {
		log.Fatal(err)
	}
	memAvailableKB, _ := strconv.Atoi(strings.Fields(line[0])[1])
	memavailable := memAvailableKB / 1024
	return memavailable
}

func Total() int {
	line, err := util.ReadProcFile("/proc/meminfo", "MemTotal:")
	if err != nil {
		log.Fatal(err)
	}
	memTotalKB, _ := strconv.Atoi(strings.Fields(line[0])[1])
	memtotal := memTotalKB / 1024
	return memtotal
}

func UsageMB() int {
	return Total() - available()
}

func UsagePercent() int {
	return int((float64(UsageMB()) / float64(Total())) * 100)
}
