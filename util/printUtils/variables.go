package printUtils

import (
	util "github.com/Vilez0/gofretch/util/other"

	osinfo "github.com/Vilez0/gofretch/os"
)

var (
	hostUser                    = osinfo.Hostname() + "@" + osinfo.Username()
	configDir                   = util.GetConfigDir()
	configLocation              = configDir + "/gofretch.conf"
	asciiDir                    = configDir + "/asciis/"
	asciiFileLocation           = getAsciiFileLocation()
	infos, commands, asciiLines = ReadAsciiAndConfig()
)
