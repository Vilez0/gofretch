package osinfo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	uptime   int
	filename string = "/proc/uptime"
	err      error
)

// Uptime returns the uptime of the system
func Uptime() string {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	defer file.Close()
	// read the file line by line
	for scanner.Scan() {
		// split the line by the dot
		str := strings.Split(scanner.Text(), ".")[0]
		// convert the uptime to int
		uptime, err = strconv.Atoi(str)
		if err != nil {
			return "Unknown"
		}
	}
	// convert the uptime to days, hours and minutes
	days := uptime / 60 / 60 / 24
	hours := uptime / 60 / 60 % 24
	minutes := uptime / 60 % 60
	// if the uptime is more than 0 days, we return the uptime in days, hours and minutes
	if days > 0 {
		return fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
	} else if hours > 0 {
		// if there is no days, we return the uptime in hours and minutes
		return fmt.Sprintf("%d hours, %d minutes", hours, minutes)
	} else if minutes > 0 {
		// if there is no days and hours, we return the uptime in minutes
		return fmt.Sprintf("%d minutes", minutes)
	}
	// if there is no days, hours and minutes, we return unknown
	return "Unknown"
}
