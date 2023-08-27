package util

import (
	"bufio"
	"os"
)

func ScanFile(filename string) []string {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	defer file.Close()
	var data []string 
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
