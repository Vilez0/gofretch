package osinfo

import (
	"os"
)

func DesktopEnvironment() string {
	// reads the DESKTOP_SESSION env, then if it empty it will read XDG_CURRENT_DESKTOP env and if it also empty it will return `Unknown`
	desktop := os.Getenv(`DESKTOP_SESSION`)
	if desktop == "" {
		desktop = os.Getenv(`XDG_CURRENT_DESKTOP`)
		if desktop == "" {
			return "Unknown"
		}
	}
	return desktop
}
