package cpu

import (
	"errors"
	util "gofretch/util/other"
	"strconv"
	"strings"
	"time"
)

func getUsageFromProcFile() (int, error) {
	indexNumber := 0
	var prevIdleTime, prevTotalTime int
	for i := 0; i <= 1; i++ {
		lines, err := util.ReadProcFile("/proc/stat", "cpu")
		if err != nil {
			return 0, err
		}
		var line string
		if indexNumber > len(lines)-1 {
			return 0, errors.New("index out of range")
		}
		line = (lines[indexNumber])[5:] // get rid of cpu plus 2 spaces

		split := strings.Fields(line)
		idleTime, _ := strconv.Atoi(split[3])
		totalTime := 0
		for _, s := range split {
			u, _ := strconv.Atoi(s)
			totalTime += u
		}
		if i > 0 {
			deltaIdleTime := idleTime - prevIdleTime
			deltaTotalTime := totalTime - prevTotalTime
			cpuUsage := (1.0 - float64(deltaIdleTime)/float64(deltaTotalTime)) * 100.0
			return int(cpuUsage), nil
		}
		prevIdleTime = idleTime
		prevTotalTime = totalTime
		time.Sleep(time.Millisecond * 250)
	}
	return 0, nil

}
func Usage() string {
	average, err := getUsageFromProcFile()
	if err != nil {
		return "Unknown"
	}

	usage := strconv.Itoa(average) + "%"
	return usage
}
