package printUtils

import (
	"bufio"
	"os"
	"strings"
	"unicode/utf8"
)

func ReadAsciiAndConfig() ([]string, []string, []string) {
	//Read Ascii file
	asciiLineLength := 0
	infos, commands := readConfigFile()
	asciiFile, err := os.Open(asciiFileLocation)
	if err != nil {
		println("Error opening ascii file: ", asciiFileLocation)
	}
	defer asciiFile.Close()
	scanner := bufio.NewScanner(asciiFile)
	var asciiLines []string
	for scanner.Scan() {
		// Add the asciiFile lines to a slice named asciiLines
		asciiLines = append(asciiLines, scanner.Text())
		// Find the longest line in ascii file and set it to asciiLineLength, because every line should
		if utf8.RuneCountInString(scanner.Text()) > asciiLineLength {
			// +2 is for space after ascii
			asciiLineLength = utf8.RuneCountInString(scanner.Text()) + 2
		}
	}
	// for check length ascii and info is the same,if not make them same(for better appearance)
	if len(asciiLines) > len(infos) {
		for i := len(infos); i < len(asciiLines); i++ {
			infos = append(infos, "")
			commands = append(commands, "")
		}
	} else if len(asciiLines) < len(infos) {

		for i := len(asciiLines); i <= len(infos); i++ {
			asciiLines = append(asciiLines, "")
		}
	}
	// add 2 empty lines for printinfo.go#11
	for i := 0; i < 2; i++ {
		asciiLines = append(asciiLines, "")
	}
	for i, e := range asciiLines {
		if utf8.RuneCountInString(e) < asciiLineLength {

			for length := utf8.RuneCountInString(e); length <= asciiLineLength; length++ {

				asciiLines[i] = asciiLines[i] + " "
			}
		}
	}
	return infos, commands, asciiLines
}

func readConfigFile() ([]string, []string) {
	var infos, commands []string
	file, err := os.Open(configLocation)
	if err != nil {
		println("Error opening config file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "#") {
			continue
		} else if strings.HasPrefix(scanner.Text(), "info") {
			replacer := strings.NewReplacer("info ", "", `"`, "")
			text := replacer.Replace(scanner.Text())
			line := strings.Split(text, ":")
			info := line[0]
			command := CheckCommands(line[1])
			infos = append(infos, info)
			commands = append(commands, command)
		} else if strings.HasPrefix(scanner.Text(), "variable") {
			replacer := strings.NewReplacer("variable ", "", `"`, "")
			text := replacer.Replace(scanner.Text())
			line := strings.Split(text, "=")
			variable := line[0]
			value := line[1]
			cheackVariables(variable, value)
		}
	}
	return infos, commands
}

func cheackVariables(variable string, value string) {
	if variable == "asciiFile" {
		asciiFileLocation = configDir + "/" + value
	}
}
