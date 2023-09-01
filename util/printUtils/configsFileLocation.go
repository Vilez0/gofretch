package printUtils

import (
	osinfo "gofretch/os"
	util "gofretch/util/other"
	"strings"
)

func getAsciiFileLocation() string {
	var asciiFileLocation string
	distroname := osinfo.DistroName()
	codename := strings.ToLower(strings.Split(strings.TrimSpace(distroname), " ")[0])
	ascciFiles := util.FindFileWithExtension(asciiDir, ".txt")
	for _, file := range ascciFiles {
		if strings.Contains(file, codename) {
			asciiFileLocation = file
			break
		}
	}
	if asciiFileLocation == "" || util.FileExists(asciiFileLocation) == false {
		asciiFileLocation = asciiDir + "default.txt"
	}

	return asciiFileLocation
}
