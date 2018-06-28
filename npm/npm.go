package npm

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/riggerthegeek/license-disco/common"
	"path/filepath"
	"github.com/spf13/cobra"
	"net/http"
	"time"
	"github.com/kyoh86/go-spdx/spdx"
)

const npmUrl = "https://registry.npmjs.org"

type NPM struct {}

func (obj *NPM) Enabled(cmd *cobra.Command) bool  {
	flag := cmd.Flags().Lookup("npm.detect")

	if flag != nil {
		return flag.Value.String() == "true"
	}

	return false
}

func (obj *NPM) Flags() []common.Flag {
	return []common.Flag {
		{
			Type: "Bool",
			Name: "detect",
			Value: false,
			Usage: "Search for npm dependencies",
		},
	}
}

func (*NPM) Name () string {
	return "NPM"
}

func (obj *NPM) Scan(path string) (error, []common.Dependency) {
	initialPackage := filepath.Join(path, "package.json")

	err, pkg := loadPackage(initialPackage)
	if err != nil {
		return err, nil
	}

	dep := []common.Dependency{}

	for dep, ver := range pkg.Dependencies {
		fmt.Println(dep)
		fmt.Println(ver)
		//getLicenseInfo(dep, ver)
	}

	return nil, dep
}

type npmPackage struct {
	Name string                 `json:"name"`
	Description string          `json:"description"`
	Versions map[string]Version `json:"versions"`
	License string              `json:"license"`
}

func getLicenseInfo (pkg string, version string) (error, *Version) {
	url := getUrl(pkg)

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err, nil
	}

	//req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := client.Do(req)
	if getErr != nil {
		return getErr, nil
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return readErr, nil
	}

	data := npmPackage{}
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		return jsonErr, nil
	}
	fmt.Println(data.License)

	license, _ := spdx.Get("MIT")
	fmt.Println(license.Text)

	return nil, nil
}

func getUrl(pkg string) string {
	/* URL encode slashes in package names */
	parsedPkg := strings.Replace(pkg, "/", "%2F", 1)

	return npmUrl + "/" + parsedPkg
}

func loadPackage (search string) (error, *Version) {
	if _ , err := os.Stat(search); os.IsNotExist(err) {
		return err, nil
	}

	raw, err := ioutil.ReadFile(search)
	if err != nil {
		return err, nil
	}

	pkg := Version{}
	json.Unmarshal(raw, &pkg)
	return nil, &pkg
}