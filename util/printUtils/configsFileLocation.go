package printUtils

import (
	"strings"

	util "github.com/Vilez0/gofretch/util/other"

	osinfo "github.com/Vilez0/gofretch/os"
)

func getAsciiFileLocation() string {
	var asciiFileLocation string
	distroname := osinfo.DistroName()
	codename := strings.ToLower(strings.Split(strings.TrimSpace(distroname), " ")[0])
	ascciFiles := util.FindFileWithExtension(asciiDir, ".txt")
	for _, file := range ascciFiles {
		if file == asciiDir+codename+".txt" {
			asciiFileLocation = file
			break
		}
	}
	if asciiFileLocation == "" || util.FileExists(asciiFileLocation) == false {
		asciiFileLocation = asciiDir + "default.txt"
	}

	return asciiFileLocation
}
