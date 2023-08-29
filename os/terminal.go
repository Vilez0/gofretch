package osinfo

import (
	"os"
)

func Terminal() string {
	term := os.Getenv("TERM")
	if term == "" {
		return "Unknown"
	}
	return term
}
