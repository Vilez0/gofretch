package osinfo

import (
	"encoding/json"
	"fmt"
	util "gofretch/util/other"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	// the database where is package managers relatead config stored, its generally under ~/.config/gofretch/packageManagers.json
	packageManagersDatabase       = util.GetConfigDir() + "/packageManagers.json"
	packageManagers               PackageManagers
	packages                      string
	installedPackageManagersCount int
)

type PackageManagers struct {
	PackageManagers []PackageManager `json:"packageManagers"`
}

type PackageManager struct {
	Name   string   `json:"name"`
	Path   string   `json:"path"`
	Params []string `json:"params"`
}

// this will calculate the number of lines in the output of the command
func calcLength(s []byte) string {
	return strconv.Itoa(len(strings.Split(string(s), "\n")) - 1)
}

// PackageCount returns the number of packages installed in the system
func PackageCount() string {
	packageManagersData, err := os.Open(packageManagersDatabase)
	if err != nil {
		return "Unknown"
	}
	defer packageManagersData.Close()
	// read our opened json as a byte array.
	byteValue, err := io.ReadAll(packageManagersData)
	if err != nil {
		return "Unknown"
	}
	// we unmarshal our byteArray to `packageManagers` which we defined above
	json.Unmarshal(byteValue, &packageManagers)
	// we iterate through every package manager and check if its installed
	for _, element := range packageManagers.PackageManagers {
		packageManager := element
		// if the package manager is installed, we run the command and get the number of packages installed
		if util.FileExists(packageManager.Path) {
			output, _ := exec.Command(packageManager.Name, packageManager.Params...).Output()
			packagesCount := calcLength(output)
			if packagesCount == "0" {
				continue
			}
			// if the package manager is installed, we add it to the packages string
			if installedPackageManagersCount == 0 {
				// if its the first package manager, we dont add a comma
				packages += fmt.Sprintf("%v (%v)", packagesCount, packageManager.Name)
			} else {
				packages += fmt.Sprintf(", %v (%v)", packagesCount, packageManager.Name)
			}
			installedPackageManagersCount++
		}
	}
	// if no package manager is installed, we return unknown
	if packages == "" {
		return "Unknown"
	}
	return packages
}
