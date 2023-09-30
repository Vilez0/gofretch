package printUtils

import "fmt"

func PrintInfo() {

	for i := 0; i < len(commands); i++ {
		if i == 0 {
			fmt.Printf("%s %s\n", asciiLines[0], hostUser)
			fmt.Printf("%s %s\n", asciiLines[1], titleSeperator())
			asciiLines = asciiLines[2:]
		}
		if infos[i] == "" && commands[i] == "" {
			fmt.Printf("%s\n", asciiLines[i])
			continue
		}
		fmt.Printf("%s %s: %s\n", asciiLines[i], infos[i], commands[i])

	}
}
