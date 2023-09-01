package cpu

import (
	"bufio"
	"os"
	"strings"
)

// Name returns the model name of the CPU.
func Name() string {
	// Open the file.
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		// Log the error.
		return "Unknown"
	}
	// Scan the file.
	scanner := bufio.NewScanner(file)
	defer file.Close()

	// Scan the file line by line.
	for scanner.Scan() {
		// Check if the line starts with "model name: ".
		if strings.HasPrefix(scanner.Text(), "model name	: ") {
			// Return the text after "model name: ".
			return strings.ReplaceAll(scanner.Text()[13:], "(R)", "")
		}
	}
	return `Unknown`
}
