package osinfo

import (
	"os"
)

func DesktopEnvironment() string {
	desktop := os.Getenv(`DESKTOP_SESSION`)
	if desktop == "" {
		desktop = os.Getenv(`XDG_CURRENT_DESKTOP`)
		if desktop == "" {
			return "Unknown"
		}
	}
	return desktop
}
