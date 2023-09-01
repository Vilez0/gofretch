package printUtils

import (
	osinfo "gofretch/os"
	util "gofretch/util/other"
)

var (
	hostUser                    = osinfo.Hostname() + "@" + osinfo.Username()
	configDir                   = util.GetConfigDir()
	configLocation              = configDir + "/gofretch.conf"
	asciiDir                    = configDir + "/asciis/"
	asciiFileLocation           = getAsciiFileLocation()
	infos, commands, asciiLines = ReadAsciiAndConfig()
)
