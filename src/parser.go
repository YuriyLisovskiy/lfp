// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"fmt"
	"strings"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"github.com/YuriyLisovskiy/lofp/src/static"
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
		for _, path := range paths {
			fileData, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			resultFileData, err := transform(fileData, cfg)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(path, resultFileData, 0644)
			if err != nil {
				return err
			}
		}
	} else if cfg.AddLicenseFile {
		createLicenseFile(cfg)
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

// transform appends license notice to given file content.
func transform(fileContent []byte, cfg Config) ([]byte, error) {
	l, err := getLicense(cfg.License)
	if err != nil {
		return []byte{}, err
	}
	var licenseNotice []byte
	if cfg.CustomLicenseNotice != "" {
		licenseNotice, err = ioutil.ReadFile(cfg.CustomLicenseNotice)
		if err != nil {
			return []byte{}, err
		}
	} else {
		if cfg.License == "unlicense" {
			licenseNotice = []byte(fmt.Sprintf("// Unlicense, see the accompanying file LICENSE or %s\n\n", l["link"]))
		} else {
			licenseNotice = []byte(fmt.Sprintf(static.LICENSE_NOTICE_TEMPLATE, cfg.Year, cfg.Author, l["name"], l["link"]))
		}
	}
	return append(licenseNotice, fileContent...), nil
}
