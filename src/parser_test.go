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
		cfg: Config{
			License:  "mit",
			Authors:   []string{"John Smith"},
			Years:     []string{"2010", "2011"},
		},
		expected: []byte("// Copyright (c) 2010, 2011 John Smith\n// Distributed under the MIT License,\n// see the accompanying file LICENSE or https://opensource.org/licenses/MIT\n\n"),
	},
}

func Test_prepareLicenseNotice(test *testing.T) {
	for _, data := range Test_transformData {
		actual, _ := prepareLicenseNotice(data.cfg)
		if len(actual) != len(data.expected) {
			test.Errorf("parser.Test_prepareLicenseNotice: actual len != expected len:\n\t%d != %d\n", len(actual), len(data.expected))
		}
		for i, b := range actual {
			if b != data.expected[i] {
				test.Errorf("parser.Test_prepareLicenseNotice: actual != expected:\n\t%s != %s\n", string(actual), string(data.expected))
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
			Authors:   []string{"John Smith"},
			Years:     []string{"2010"},
		},
		expected: ErrLicenseNotFound,
	},
}

func Test_prepareLicenseNoticeErrLicenseNotFound(test *testing.T) {
	for _, data := range Test_transformDataErrLicenseNotFound {
		_, err := prepareLicenseNotice(data.cfg)
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
authors:
  - John Smith
years:
  - 2018
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
		Authors:              []string{"John Smith"},
		Years:                []string{"2018"},
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
	if len(actual.Authors) != len(data.expected.Authors) {
		test.Errorf("parser.Test_parseConfig: actual != expected:\n\t%d != %d\n", len(actual.Authors), len(data.expected.Authors))
	}
	for i, author := range actual.Authors {
		if author != data.expected.Authors[i] {
			test.Errorf("parser.Test_parseConfig: actual != expected:\n\t%s != %s\n", author, data.expected.Authors[i])
		}
	}
	if len(actual.Years) != len(data.expected.Years) {
		test.Errorf("parser.Test_parseConfig: actual != expected:\n\t%d != %d\n", len(actual.Years), len(data.expected.Years))
	}
	for i, year := range actual.Years {
		if year != data.expected.Years[i] {
			test.Errorf("parser.Test_parseConfig: actual != expected:\n\t%s != %s\n", year, data.expected.Years[i])
		}
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
