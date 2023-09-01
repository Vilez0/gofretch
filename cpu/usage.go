package cpu

import (
	"errors"
	"strconv"
	"strings"
	"time"

	util "github.com/Vilez0/gofretch/util/other"
)

var line string

func getUsageFromProcFile() (int, error) {
	indexNumber := 0
	var prevIdleTime, prevTotalTime int
	//This function will run 2 times and then return the cpu usage
	for i := 0; i <= 1; i++ {
		// return the lines in /proc/stat that start with `cpu` and assign them to the lines variable
		lines, err := util.ReadProcFile("/proc/stat", "cpu")
		if err != nil {
			return 0, err
		}
		if indexNumber > len(lines)-1 {
			return 0, errors.New("index out of range")
		}
		line = (lines[indexNumber])[5:] // get rid of `cpu` plus 2 spaces in /proc/stat file

		split := strings.Fields(line)         // split it from whitspaces,it will return something like ["4006536","61384","1068355","2318371","55870","99790","45491","0","0","0"]
		idleTime, _ := strconv.Atoi(split[3]) // gett the idle time
		totalTime := 0
		for _, s := range split {
			u, _ := strconv.Atoi(s)
			totalTime += u
		}
		// calculate the cpu usage
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
	intUsage, err := getUsageFromProcFile()
	if err != nil {
		return "Unknown"
	}

	usage := strconv.Itoa(intUsage) + "%" // for return it as something like 45%
	return usage
}
