package osinfo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Uptime() string {
	var uptime int
	filename := "/proc/uptime"
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	defer file.Close()
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), ".")[0]
		floatUptime, err := strconv.ParseFloat(str, 64)
		if err != nil {
			println(`Error: `, err.Error())
		}
		uptime = int(floatUptime)
	}
	days := uptime / 60 / 60 / 24
	hours := uptime / 60 / 60 % 24
	minutes := uptime / 60 % 60
	if days > 0 {
		return fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
	} else if hours > 0 {
		return fmt.Sprintf("%d hours, %d minutes", hours, minutes)
	} else if minutes > 0 {
		return fmt.Sprintf("%d minutes", minutes)
	}
	return "Unknown"
}
