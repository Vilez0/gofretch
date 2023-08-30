package util

import (
	"bufio"
	"os"
	"strings"
)

func ReadProcFile(filename string, info string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var lines []string

	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), info) {
			lines = append(lines, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
