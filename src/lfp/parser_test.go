// Copyright (c) 2018 Yuriy Lisovskiy
//
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package lfp

import "testing"

var Test_prepareLicenseNoticeData = []struct {
	content  []byte
	cfg      Config
	ext      string
	expected []byte
}{
	{
		cfg: Config{
			License: "mit",
			Authors: []Author{
				{
					Name: "John Smith",
					Year: "2010",
				},
				{
					Name: "John Smith 2",
					Year: "2011",
				},
			},
		},
		ext: "cpp",
		expected: []byte(
			`// Copyright (c) 2010 John Smith
// Copyright (c) 2011 John Smith 2
// 
// Distributed under the MIT License,
// see the accompanying file LICENSE or https://opensource.org/licenses/MIT

`),
	},
	{
		cfg: Config{
			License: "mit",
			Authors: []Author{
				{
					Name: "John Smith",
					Year: "2010",
				},
				{
					Name: "John Smith 2",
					Year: "2011",
				},
			},
		},
		ext: "html",
		expected: []byte(
			`<!--
 Copyright (c) 2010 John Smith
 Copyright (c) 2011 John Smith 2
 
 Distributed under the MIT License,
 see the accompanying file LICENSE or https://opensource.org/licenses/MIT
-->

`),
	},
}

func Test_prepareLicenseNotice(test *testing.T) {
	for _, data := range Test_prepareLicenseNoticeData {
		actual, _ := prepareLicenseNotice(data.cfg, data.ext)
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

var Test_prepareLicenseNoticeErrLicenseNotFound_Data = []struct {
	content  []byte
	cfg      Config
	ext      string
}{
	{
		content: []byte("package main\n\nfunc main() {\n\n}"),
		cfg: Config{
			License: "some-unknown-license",
			Authors: []Author{
				{
					Name: "John Smith",
					Year: "2010",
				},
			},
		},
		ext:      "go",
	},
	{
		content: []byte("package main\n\nfunc main() {\n\n}"),
		cfg: Config{
			License: "apache-2.0",
			Authors: []Author{
				{
					Name: "John Smith",
					Year: "2010",
				},
			},
		},
		ext:      "some-ext",
	},
}

func Test_prepareLicenseNoticeErr(test *testing.T) {
	for _, data := range Test_prepareLicenseNoticeErrLicenseNotFound_Data {
		_, err := prepareLicenseNotice(data.cfg, data.ext)
		if err == nil {
			test.Errorf(
				"parser.Test_prepareLicenseNoticeErr: func does not return an error for input %s",
				data.cfg.License,
			)
		}
	}
}

var Test_parseConfig_Data = []struct {
	input    []byte
	ext      string
	expected Config
} {
	{
		input: []byte(`
authors:
  - name: John Smith
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
		ext: "yml",
		expected: Config{
			Authors: []Author{
				{
					Name: "John Smith",
					Year: "2018",
				},
			},
			ProgramName:         "Skynet",
			Paths:               []string{"parser/...", "generator/...", "execute.go", "main.go"},
			License:             "apache-v2",
			ProjectRoot:         "/home/johnsmith/go/src/github.com/author/Skynet",
			CustomLicenseNotice: "",
			AddLicenseFile:      true,
			AddLicenseNotice:    true,
		},
	},
	{
		input: []byte(`
authors:
  - name: John Smith
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
		ext: "yaml",
		expected: Config{
			Authors: []Author{
				{
					Name: "John Smith",
					Year: "2018",
				},
			},
			ProgramName:         "Skynet",
			Paths:               []string{"parser/...", "generator/...", "execute.go", "main.go"},
			License:             "apache-v2",
			ProjectRoot:         "/home/johnsmith/go/src/github.com/author/Skynet",
			CustomLicenseNotice: "",
			AddLicenseFile:      true,
			AddLicenseNotice:    true,
		},
	},
	{
		input: []byte(`
{
	"authors": [
		{
			"name": "John Smith",
			"year": "2018"
		}
	],
	"program_name": "Skynet",
	"paths": [
		"parser/...",
        "generator/...",
        "execute.go",
        "main.go"
	],
	"license": "apache-v2",
	"project_root": "/home/johnsmith/go/src/github.com/author/Skynet",
	"add_license_file": true,
	"add_license_notice": true
}

`),
		ext: "json",
		expected: Config{
			Authors: []Author{
				{
					Name: "John Smith",
					Year: "2018",
				},
			},
			ProgramName:         "Skynet",
			Paths:               []string{"parser/...", "generator/...", "execute.go", "main.go"},
			License:             "apache-v2",
			ProjectRoot:         "/home/johnsmith/go/src/github.com/author/Skynet",
			CustomLicenseNotice: "",
			AddLicenseFile:      true,
			AddLicenseNotice:    true,
		},
	},
	{
		input: []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<Config>
	<Authors>
		<Author>
			<Name>John Smith</Name>
			<Year>2018</Year>
		</Author>
	</Authors>
	<ProgramName>Skynet</ProgramName>
	<Paths>
		<Path>parser/...</Path>
		<Path>generator/...</Path>
        <Path>execute.go</Path>
        <Path>main.go</Path>
	</Paths>
	<License>apache-v2</License>
	<ProjectRoot>/home/johnsmith/go/src/github.com/author/Skynet</ProjectRoot>
	<AddLicenseFile>true</AddLicenseFile>
	<AddLicenseNotice>true</AddLicenseNotice>
</Config>
`),
		ext: "xml",
		expected: Config{
			Authors: []Author{
				{
					Name: "John Smith",
					Year: "2018",
				},
			},
			ProgramName:         "Skynet",
			Paths:               []string{"parser/...", "generator/...", "execute.go", "main.go"},
			License:             "apache-v2",
			ProjectRoot:         "/home/johnsmith/go/src/github.com/author/Skynet",
			CustomLicenseNotice: "",
			AddLicenseFile:      true,
			AddLicenseNotice:    true,
		},
	},
}

func Test_parseConfig(test *testing.T) {
	for _, data := range Test_parseConfig_Data {
		actual, _ := parseConfig(data.input, data.ext)
		if len(actual.Authors) != len(data.expected.Authors) {
			test.Errorf("parser.Test_parseConfig[%s]: actual != expected:\n\t%d != %d\n",
				data.ext, len(actual.Authors), len(data.expected.Authors),
			)
		}
		for i, author := range actual.Authors {
			if author != data.expected.Authors[i] {
				test.Errorf("parser.Test_parseConfig[%s]: actual != expected:\n\t%s != %s\n",
					data.ext, author, data.expected.Authors[i],
				)
			}
		}
		if len(actual.Authors) != len(data.expected.Authors) {
			test.Errorf("parser.Test_parseConfig[%s]: actual != expected:\n\t%d != %d\n",
				data.ext, len(actual.Authors), len(data.expected.Authors),
			)
		}
		for i, author := range actual.Authors {
			if author.Name != data.expected.Authors[i].Name {
				test.Errorf("parser.Test_parseConfig[%s]: actual != expected:\n\t%s != %s\n",
					data.ext, author.Name, data.expected.Authors[i].Name,
				)
			}
			if author.Year != data.expected.Authors[i].Year {
				test.Errorf("parser.Test_parseConfig[%s]: actual != expected:\n\t%s != %s\n",
					data.ext, author.Year, data.expected.Authors[i].Year,
				)
			}
		}
		if actual.ProgramName != data.expected.ProgramName {
			test.Errorf("parser.Test_parseConfig[%s]: actual != expected:\n\t%s != %s\n",
				data.ext, actual.ProgramName, data.expected.ProgramName,
			)
		}
		if len(actual.Paths) != len(data.expected.Paths) {
			test.Errorf("parser.Test_parseConfig[%s]: actual != expected:\n\t%d != %d\n",
				data.ext, len(actual.Paths), len(data.expected.Paths),
			)
		}
		for i, path := range actual.Paths {
			if path != data.expected.Paths[i] {
				test.Errorf("parser.Test_parseConfig[%s]: actual != expected:\n\t%s != %s\n",
					data.ext, path, data.expected.Paths[i],
				)
			}
		}
		if actual.License != data.expected.License {
			test.Errorf("parser.Test_parseConfig[%s]: actual != expected:\n\t%s != %s\n",
				data.ext, actual.License, data.expected.License,
			)
		}
		if actual.ProjectRoot != data.expected.ProjectRoot {
			test.Errorf("parser.Test_parseConfig[%s]: actual != expected:\n\t%s != %s\n",
				data.ext, actual.ProjectRoot, data.expected.ProjectRoot,
			)
		}
		if actual.CustomLicenseNotice != data.expected.CustomLicenseNotice {
			test.Errorf("parser.Test_parseConfig[%s]: actual != expected:\n\t%s != %s\n",
				data.ext, actual.CustomLicenseNotice, data.expected.CustomLicenseNotice,
			)
		}
		if actual.AddLicenseFile != data.expected.AddLicenseFile {
			test.Errorf("parser.Test_parseConfig[%s]: actual != expected:\n\t%t != %t\n",
				data.ext, actual.AddLicenseFile, data.expected.AddLicenseFile,
			)
		}
		if actual.AddLicenseNotice != data.expected.AddLicenseNotice {
			test.Errorf("parser.Test_parseConfig[%s]: actual != expected:\n\t%t != %t\n",
				data.ext, actual.AddLicenseNotice, data.expected.AddLicenseNotice,
			)
		}
	}
}
