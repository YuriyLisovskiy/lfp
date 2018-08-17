// Copyright (c) 2018 Yuriy Lisovskiy
//
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package lfp

import (
	"os"
	"regexp"
	"strings"

	"github.com/YuriyLisovskiy/licenses/api/golang"
)

func getLicense(license string) (golang.License, error) {
	client := golang.Client{}
	return client.GetLicense(license)
}

// createLicenseFile generates license from a template
func createLicenseFile(cfg Config) ([]byte, error) {
	var ret []byte
	path := cfg.ProjectRoot + "/LICENSE"

	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return ret, err
		}
		defer file.Close()
	}
	license, err := getLicense(cfg.License)
	if err != nil {
		return ret, err
	}
	licenseContent := license.Content()
	switch cfg.License {
	case "apache-2.0", "mit", "bsd-2-clause", "bsd-3-clause":
		licenseContent, err = prepareLicense(licenseContent, cfg.Authors, map[string]string{})
	case "gpl-3.0":
		licenseContent, err = prepareLicense(licenseContent, cfg.Authors, map[string]string{
			"<program name>": cfg.ProgramName,
			"<program description>": cfg.ProgramDescription,
		})
	case "lgpl-2.1", "gpl-2.0", "agpl-3.0":
		licenseContent, err = prepareLicense(licenseContent, cfg.Authors, map[string]string{
			"<program description>": cfg.ProgramDescription,
		})
	default:
	}
	if err != nil {
		return ret, err
	}
	ret = []byte(licenseContent)
	return ret, nil
}

// prepareLicense replaces all given keywords to actual data
func prepareLicense(template string, authors []Author, data map[string]string) (string, error) {
	ret := template
	for key, value := range data {
		ret = strings.Replace(ret, key, value, -1)
	}
	crRegex, _ := regexp.Compile(`{{.+}}`)
	for crRegex.MatchString(ret) {

		// Find header template location
		loc := crRegex.FindStringIndex(ret)

		header := processHeader(crRegex.FindString(ret), authors, findIndentReverse(ret[:loc[0]]))

		// Replace header template with aggregated headers
		ret = ret[:loc[0]] + header + ret[loc[1]:]
	}
	return ret, nil
}

// findIndentReverse searches for an indent at the end of template fragment before header template
func findIndentReverse(templateFragment string) string {
	indent := ""
	start := len(templateFragment)-1
	for start >= 0 {
		if templateFragment[start] == ' ' || templateFragment[start] == '\t' {
			indent += string(templateFragment[start])
			start--
		} else {
			break
		}
	}
	return indent
}

// processHeader aggregates headers from a template for all given authors
func processHeader(header string, authors []Author, indent string) string {
	header = strings.Replace(header, "{", "", -1)
	header = strings.Replace(header, "}", "", -1)
	ret := ""
	for i, author := range authors {
		headerDone := strings.Replace(header, "<year>", author.Year, -1)
		headerDone = strings.Replace(headerDone, "<author>", author.Name, -1)
		if i < len(authors)-1 {
			headerDone += "\n"
		}
		if i != 0 {
			ret += indent
		}
		ret += headerDone
	}
	return ret
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode().IsDir()
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return true
	}
	return info.Mode().IsRegular()
}

func shift(str, indent string) string {
	ret := strings.Replace(str, "\n", "\n" + indent, -1)
	return indent + ret
}
