// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"fmt"
	"strings"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"github.com/YuriyLisovskiy/lfp/src/static"
)

func processPaths(cfg Config) error {
	if cfg.AddLicenseNotice {

		// parse all paths and its children if specified
		var paths []string
		for _, path := range cfg.Paths {
			subPaths, err := parsePath(cfg.ProjectRoot + "/" + path)
			if err != nil {
				return err
			}
			paths = append(paths, subPaths...)
		}
		licenseNotice, err := prepareLicenseNotice(cfg)
		if err != nil {
			return err
		}
		for _, path := range paths {
			fileData, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			resultFileData := append(licenseNotice, fileData...)
			err = ioutil.WriteFile(path, resultFileData, 0644)
			if err != nil {
				return err
			}
		}
	}
	if cfg.AddLicenseFile {
		data, err := createLicenseFile(cfg)
		if err != nil {
			return err
		}

		// write to file generated license
		err = ioutil.WriteFile(cfg.ProjectRoot + "/LICENSE", data, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

// parseConfig parses given configuration data
func parseConfig(data []byte) (cfg Config, err error) {
	err = yaml.Unmarshal(data, &cfg)
	return
}

// parsePath parses given path and its children recursively, if specified.
func parsePath(path string) ([]string, error) {

	// check if path includes all children - /...
	includeDirs := false
	if strings.HasSuffix(path, "/...") {
		includeDirs = true
		path = path[:len(path)-3]
	}

	if isFile(path) {
		return []string{path}, nil
	}

	// read path data
	readResult, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var ret []string

	// for each sub path if sub path is file, append it to the result,
	// otherwise if path includes all children, append all its children recursively
	for _, pathData := range readResult {
		newPath := path + pathData.Name()
		if includeDirs && isDir(newPath) {
			subRet, err := parsePath(newPath + "/...")
			if err != nil {
				return nil, err
			}
			ret = append(ret, subRet...)
		} else if isFile(newPath) {
			ret = append(ret, newPath)
		}
	}
	return ret, nil
}

// prepareLicenseNotice create license notice from given template
func prepareLicenseNotice(cfg Config) (ret []byte, err error) {
	var template string
	if cfg.CustomLicenseNotice != "" {
		templateBytes, err := ioutil.ReadFile(cfg.CustomLicenseNotice)
		if err != nil {
			return []byte{}, err
		}
		template = string(templateBytes[:])
	} else {
		header := static.LICENSE_NOTICE_TEMPLATE["head"]
		for i := range cfg.Authors {
			template += header
			if i < len(cfg.Authors)-1 {
				template += "\n"
			}
		}
		template += static.LICENSE_NOTICE_TEMPLATE["body"]
	}
	license, err := getLicense(cfg.License)
	if err != nil {
		return
	}
	if cfg.License == "unlicense" {
		ret = []byte(fmt.Sprintf("// Unlicense, see the accompanying file LICENSE or %s\n\n", license["link"]))
	} else {

		// Set license name
		retStr := strings.Replace(template, "<license name>", license["name"], -1)

		// Set license link
		retStr = strings.Replace(retStr, "<license link>", license["link"], -1)

		// Set authors
		for _, author := range cfg.Authors {
			retStr = strings.Replace(retStr, "<author>", author.Name, 1)
			retStr = strings.Replace(retStr, "<year>", author.Year, 1)
		}

		// Set comments
		retStr = strings.Replace(retStr, "<comment>", "//", -1)

		// Set other fields if custom license notice template is provided
		if cfg.CustomLicenseNotice != "" {

			// Set program name
			retStr = strings.Replace(retStr, "<program name>", cfg.ProgramName, -1)

			// Set program description
			retStr = strings.Replace(retStr, "<program description>", cfg.ProgramDescription, -1)
		}
		ret = []byte(retStr)
	}
	return
}
