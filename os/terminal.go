package osinfo

import (
	"os"
)

func Terminal() string {
	// get the terminal name from the environment variable
	term := os.Getenv("TERM")
	if term == "" {
		return "Unknown"
	}
	return term
}
