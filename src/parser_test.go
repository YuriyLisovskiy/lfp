// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import "testing"

var Test_transformData = []struct {
	content  []byte
	cfg      Config
	expected []byte
}{
	{
		content:  []byte("package main\n\nfunc main() {\n\n}"),
		cfg: Config{
			License:  "mit",
			Author:   "John Smith",
			Year:     "2010",
		},
		expected: []byte("// Copyright (c) 2010 John Smith\n// Distributed under the MIT License,\n// see the accompanying file LICENSE or https://opensource.org/licenses/MIT\n\npackage main\n\nfunc main() {\n\n}"),
	},
}

func Test_transform(test *testing.T) {
	for _, data := range Test_transformData {
		actual, _ := transform(data.content, data.cfg)
		if len(actual) != len(data.expected) {
			test.Errorf("parser.Test_transform: actual len != expected len:\n\t%d != %d\n", len(actual), len(data.expected))
		}
		for i, b := range actual {
			if b != data.expected[i] {
				test.Errorf("parser.Test_transform: actual != expected:\n\t%x != %x\n", actual, data.expected)
			}
		}
	}
}

var Test_transformDataErrLicenseNotFound = []struct {
	content  []byte
	cfg      Config
	expected error
}{
	{
		content:  []byte("package main\n\nfunc main() {\n\n}"),
		cfg: Config{
			License:  "some-unknown-license",
			Author:   "John Smith",
			Year:     "2010",
		},
		expected: ErrLicenseNotFound,
	},
}

func Test_transformErrLicenseNotFound(test *testing.T) {
	for _, data := range Test_transformDataErrLicenseNotFound {
		_, err := transform(data.content, data.cfg)
		if err != data.expected {
			test.Errorf("parser.Test_transformErrLicenseNotFound: actual != expected:\n\t%s != %s\n", err, data.expected)
		}
	}
}

var Test_parseConfigData = struct {
	input    []byte
	expected Config
}{
	input: []byte(`
author: John Smith
year: 2018
program_name: Skynet
paths:
  - parser/...
  - generator/...
  - execute.go
  - main.go

license: apache-v2
project_root: /home/johnsmith/go/src/github.com/author/Skynet
add_license_file: true
add_license_notice: true
`),
	expected: Config{
		Author:              "John Smith",
		Year:                "2018",
		ProgramName:         "Skynet",
		Paths:               []string{"parser/...", "generator/...", "execute.go", "main.go"},
		License:             "apache-v2",
		ProjectRoot:         "/home/johnsmith/go/src/github.com/author/Skynet",
		CustomLicenseNotice: "",
		AddLicenseFile:      true,
		AddLicenseNotice:    true,
	},
}

func Test_parseConfig(test *testing.T) {
	data := Test_parseConfigData
	actual, _ := parseConfig(data.input)
	if actual.Author != data.expected.Author {
		test.Errorf("parser.Test_parseConfig: actual != expected:\n\t%s != %s\n", actual.Author, data.expected.Author)
	}
	if actual.Year != data.expected.Year {
		test.Errorf("parser.Test_parseConfig: actual != expected:\n\t%s != %s\n", actual.Year, data.expected.Year)
	}
	if actual.ProgramName != data.expected.ProgramName {
		test.Errorf("parser.Test_parseConfig: actual != expected:\n\t%s != %s\n", actual.ProgramName, data.expected.ProgramName)
	}
	if len(actual.Paths) != len(data.expected.Paths) {
		test.Errorf("parser.Test_parseConfig: actual != expected:\n\t%d != %d\n", len(actual.Paths), len(data.expected.Paths))
	}
	for i, path := range actual.Paths {
		if path != data.expected.Paths[i] {
			test.Errorf("parser.Test_parseConfig: actual != expected:\n\t%s != %s\n", path, data.expected.Paths[i])
		}
	}
	if actual.License != data.expected.License {
		test.Errorf("parser.Test_parseConfig: actual != expected:\n\t%s != %s\n", actual.License, data.expected.License)
	}
	if actual.ProjectRoot != data.expected.ProjectRoot {
		test.Errorf("parser.Test_parseConfig: actual != expected:\n\t%s != %s\n", actual.ProjectRoot, data.expected.ProjectRoot)
	}
	if actual.CustomLicenseNotice != data.expected.CustomLicenseNotice {
		test.Errorf("parser.Test_parseConfig: actual != expected:\n\t%s != %s\n", actual.CustomLicenseNotice, data.expected.CustomLicenseNotice)
	}
	if actual.AddLicenseFile != data.expected.AddLicenseFile {
		test.Errorf("parser.Test_parseConfig: actual != expected:\n\t%t != %t\n", actual.AddLicenseFile, data.expected.AddLicenseFile)
	}
	if actual.AddLicenseNotice != data.expected.AddLicenseNotice {
		test.Errorf("parser.Test_parseConfig: actual != expected:\n\t%t != %t\n", actual.AddLicenseNotice, data.expected.AddLicenseNotice)
	}
}
