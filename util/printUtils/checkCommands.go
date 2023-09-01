package printUtils

import (
	"fmt"
	"gofretch/cpu"
	"gofretch/gpu"
	"gofretch/mem"
	"gofretch/sys"
	osinfo "gofretch/os"
	"strings"
)

func CheckCommands(info string) string {
	info_without_spaces := strings.ReplaceAll(info, " ", "")
	switch info_without_spaces {
	case "distro":
		return osinfo.DistroName()
	case "kernel":
		return osinfo.KernelName()
	case "uptime":
		return osinfo.Uptime()
	case "cpu":
		return cpu.Name()
	case "gpu":
		return gpu.Name()
	case "memory":
		return fmt.Sprintf("%vMiB / %vMiB  (%v%%)", mem.UsageMB(), mem.Total(), mem.UsagePercent())
	case "cpu_usage":
		return cpu.Usage()
	case "model":
		return sys.ModelName()
	case "desktop":
		return osinfo.DesktopEnvironment()
	case "packages":
		return fmt.Sprintf("%v", osinfo.PackageCount())
	case "shell":
		return osinfo.Shell()
	case "resolution":
		return sys.Resolution()
	case "terminal":
		return osinfo.Terminal()
	}
	return info
}
