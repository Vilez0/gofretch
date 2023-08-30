package osinfo

import (
	"encoding/json"
	"fmt"
	"gofretch/util/other"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type PackageManagers struct {
	PackageManagers []PackageManager `json:"packageManagers"`
}

type PackageManager struct {
	Name   string   `json:"name"`
	Path   string   `json:"path"`
	Params []string `json:"params"`
}

func calcLength(s []byte) string {
	return ` ` + strconv.Itoa(len(strings.Split(string(s), "\n"))-1)
}

func PackageCount() string {
	var packageManagers PackageManagers
	packageManagersData, err := os.Open("./os/packageManagers.json")
	var packages string
	if err != nil {
		panic(err)
	}
	defer packageManagersData.Close()
	byteValue, err := io.ReadAll(packageManagersData)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(byteValue, &packageManagers)
	var installedPackageManagersCount int
	for _, element := range packageManagers.PackageManagers {
		packageManager := element
		if util.FileExists(packageManager.Path) {
			output, _ := exec.Command(packageManager.Name, packageManager.Params...).Output()
			if installedPackageManagersCount == 0 {
				packages += fmt.Sprintf("%v (%v)", calcLength(output), packageManager.Name)
			} else {
				packages += fmt.Sprintf(", %v (%v)", calcLength(output), packageManager.Name)
			}
			installedPackageManagersCount++
		}
	}
	if packages == "" {
		return "Unknown"
	}
	return packages
}
