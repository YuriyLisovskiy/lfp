// Copyright (c) 2018 Yuriy Lisovskiy
//
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package lfp

import (
	"fmt"
	"errors"
	"regexp"
	"strings"
	"io/ioutil"
	"encoding/xml"
	"encoding/json"
	
	"gopkg.in/yaml.v2"
	"github.com/YuriyLisovskiy/licenses/api/golang"
)

func process(cfg Config) error {

	// Check if at least one of license file or license option was provided
	if !cfg.AddLicenseFile && !cfg.AddLicenseNotice {
		return ConfigErrAddLicenseFileNoticeRequired
	}

	if cfg.AddLicenseNotice {

		// Parse all paths and its children if specified
		var paths []string
		for _, path := range cfg.Paths {
			subPaths, err := parsePath(cfg.ProjectRoot + "/" + path)
			if err != nil {
				return err
			}
			paths = append(paths, subPaths...)
		}

		// Map for holding license notices, it prevents from generating
		// copies of notices
		notices := make(map[string][]byte)
		for _, path := range paths {
			fileData, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			ext := getExtension(path)

			// Check if such notice already exists in map
			if _, ok := notices[ext]; !ok {

				// If not then create one more and add it to map
				licenseNotice, err := prepareLicenseNotice(cfg, ext)

				// If comment was not found, ignore this file
				if err == ErrCommentNotFound {
					println(PROGRAM_NAME + ": info: file/extension " + ext + " is not supported")
					continue
				} else if err != nil {
					return err
				}
				notices[ext] = licenseNotice
			}
			resultFileData := append(notices[ext], fileData...)
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

// parseConfig parses given configuration data.
func parseConfig(data []byte, cfgFile string) (cfg Config, err error) {
	switch strings.ToLower(cfgFile) {
	case "yml", "yaml":
		err = yaml.Unmarshal(data, &cfg)
	case "json":
		err = json.Unmarshal(data, &cfg)
	case "xml":
		err = xml.Unmarshal(data, &cfg)
	default:
		err = errors.New("invalid config file")
	}
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

// prepareLicenseNotice create license notice from given template.
func prepareLicenseNotice(cfg Config, ext string) (ret []byte, err error) {
	noticeTemplate := getNotice(cfg.License)
	if cfg.CustomLicenseNotice != "" {
		templateBytes, err := ioutil.ReadFile(cfg.CustomLicenseNotice)
		if err != nil {
			return []byte{}, err
		}
		noticeTemplate = string(templateBytes[:])
	}
	crRegex, _ := regexp.Compile(`{{.+}}`)
	loc := crRegex.FindStringIndex(noticeTemplate)
	template := ""
	if loc != nil {
		header := crRegex.FindString(noticeTemplate)
		for i := range cfg.Authors {
			template += header
			if i < len(cfg.Authors)-1 {
				template += "\n"
			}
		}
	}
	start := noticeTemplate[:loc[0]]
	if start == "\n" {
		start = ""
	}
	template = start + template + noticeTemplate[loc[1]:]
	commentStart, commentEnd, err := getComments(ext)
	if err != nil {
		return
	}
	template = shift(template, " ")
	if commentStart != "" && commentEnd != "" {
		template = "<comment>\n" + template + "\n<comment>"
	} else {
		template = "<comment>" + strings.Replace(template, "\n", "\n<comment>", -1)
	}
	template = strings.Replace(strings.Replace(template, "{", "", -1), "}", "", -1) + "\n\n"
	license, err := getLicense(cfg.License)
	if err != nil {
		return
	}
	if cfg.License == "unlicense" {
		ret = []byte(fmt.Sprintf("// Unlicense, see the accompanying file LICENSE or %s\n\n", license.Link()))
	} else {
		ret = []byte(replaceKeys(template, license, cfg, commentStart, commentEnd))
	}
	return
}

// getNotice downloads license notice, if the notice does not exist, returns
// default license notice.
func getNotice(license string) string {
	client := golang.Client{}
	notice, err := client.GetHeader(license)
	if err != nil {
		return LICENSE_NOTICE_TEMPLATE
	}
	return notice
}

// replaceKeys replaces template keywords by an actual data.
func replaceKeys(template string, license golang.License, cfg Config, cStart, cEnd string) string {
	// Set license name
	retStr := strings.Replace(template, "<license name>", license.Name(), -1)

	// Set license link
	retStr = strings.Replace(retStr, "<license link>", license.Link(), -1)

	// Set authors
	for _, author := range cfg.Authors {
		if author.Name != "" {
			retStr = strings.Replace(retStr, "<author>", author.Name, 1)
		}
		if author.Year != "" {
			retStr = strings.Replace(retStr, "<year>", author.Year, 1)
		}
	}

	// Set comments
	if cStart != "" && cEnd != "" {
		retStr = strings.Replace(retStr, "<comment>", cStart, 1)
		retStr = strings.Replace(retStr, "<comment>", cEnd, 1)
	} else {
		retStr = strings.Replace(retStr, "<comment>", cStart, -1)
	}

	// Set program name
	if cfg.ProgramName != "" {
		retStr = strings.Replace(retStr, "<program name>", cfg.ProgramName, -1)
	}

	// Set program description
	if cfg.ProgramDescription != "" {
		retStr = strings.Replace(retStr, "<program description>", cfg.ProgramDescription, -1)
	}
	return retStr
}
