// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import "testing"

var Test_validate_Data = struct {
	cfg Config
	err error
}{
	cfg: Config{
		Authors: []Author{
			{
				"John Smith",
				"2010",
			},
		},
		ProgramName: "some program",
		Paths: []string{
			"test/...",
			"some/",
			"./dir/...",
		},
		License:             "apache-v2",
		AddLicenseFile:      true,
		AddLicenseNotice:    false,
		ProjectRoot:         "/home/root/path/to/project/some_program",
		CustomLicenseNotice: "",
	},
	err: nil,
}

func Test_validate(test *testing.T) {
	if err := Test_validate_Data.cfg.validate(); err != nil {
		test.Errorf("config.Test_validate: actual != expected:\n\t%s != %s\n", err, Test_validate_Data.err)
	}
}

var Test_normalize_Data = struct {
	cfg Config
	err error
}{
	cfg: Config{
		Authors: []Author{
			{
				"John Smith",
				"2010",
			},
		},
		Paths: []string{
			"/test/...",
			"some/",
			"./dir/...",
			"./",
		},
		License:          "apache-v2",
		AddLicenseFile:   true,
		AddLicenseNotice: false,
		ProjectRoot:      "/home/root/path/to/project/program_name/",
	},
	err: nil,
}

func Test_normalize(test *testing.T) {
	data := Test_normalize_Data
	data.cfg, _ = data.cfg.normalize()
	if len(data.cfg.Paths) != 3 {
		test.Errorf("config.Test_normalize: actual != expected:\n\t%d != %d\n", len(data.cfg.Paths), 3)
	}
	if data.cfg.Paths[0] != "test/..." {
		test.Errorf("config.Test_normalize: actual != expected:\n\t%s != %s\n", data.cfg.Paths[0], "test/...")
	}
	if data.cfg.Paths[1] != "some/" {
		test.Errorf("config.Test_normalize: actual != expected:\n\t%s != %s\n", data.cfg.Paths[1], "some/")
	}
	if data.cfg.Paths[2] != "dir/..." {
		test.Errorf("config.Test_normalize: actual != expected:\n\t%s != %s\n", data.cfg.Paths[2], "dir/...")
	}
	if data.cfg.ProjectRoot != "/home/root/path/to/project/program_name" {
		test.Errorf("config.Test_normalize: actual != expected:\n\t%s != %s\n", data.cfg.ProjectRoot, "/home/root/path/to/project/program_name")
	}
}

var Test_removeItem_Data = []struct {
	input    []string
	pos      int
	expected []string
}{
	{
		pos: 1,
		input: []string{
			"1", "2", "3",
		},
		expected: []string{
			"1", "3",
		},
	},
	{
		pos: 2,
		input: []string{
			"1", "2", "3",
		},
		expected: []string{
			"1", "2",
		},
	},
}

func Test_removeItem(test *testing.T) {
	for _, data := range Test_removeItem_Data {
		actual, _ := removeItem(data.input, data.pos)
		if len(actual) != len(data.expected) {
			test.Errorf("config.Test_removeItem: actual != expected:\n\t%s != %s\n", actual, data.expected)
		}
		for i, item := range actual {
			if item != data.expected[i] {
				test.Errorf("config.Test_removeItem: actual != expected:\n\t%s != %s\n", actual, data.expected)
			}
		}
	}
}

var Test_removeItem_ErrIndexOutOfRangeData = struct {
	input    []string
	pos      int
	expected error
}{
	pos: 5,
	input: []string{
		"1", "2", "3",
	},
	expected: ErrIndexOutOfRange,
}

func Test_removeItem_ErrIndexOutOfRange(test *testing.T) {
	data := Test_removeItem_ErrIndexOutOfRangeData
	_, actual := removeItem(data.input, data.pos)
	if actual != data.expected {
		test.Errorf("config.Test_removeItem_ErrIndexOutOfRange: actual != expected:\n\t%s != %s\n", actual, data.expected)
	}
}

var Test_validate_ErrData = []struct {
	cfg      Config
	expected error
}{
	{
		cfg: Config{
			Authors:     []Author{},
			ProgramName: "some program",
			Paths: []string{
				"test/...",
				"some/",
				"./dir/...",
			},
			License:             "apache-v2",
			AddLicenseFile:      true,
			AddLicenseNotice:    false,
			ProjectRoot:         "/home/root/path/to/project/some_program",
			CustomLicenseNotice: "",
		},
		expected: ConfigErrAuthorRequired,
	},
	{
		cfg: Config{
			Authors:     []Author{
				{
					"John Smith",
					"",
				},
			},
			ProgramName: "some program",
			Paths: []string{
				"test/...",
				"some/",
				"./dir/...",
			},
			License:             "apache-v2",
			AddLicenseFile:      true,
			AddLicenseNotice:    false,
			ProjectRoot:         "/home/root/path/to/project/some_program",
			CustomLicenseNotice: "",
		},
		expected: ConfigErrYearsAuthors,
	},
	{
		cfg: Config{
			Authors:     []Author{
				{
					"John Smith",
					"2010",
				},
				{
					"",
					"",
				},
			},
			ProgramName: "some program",
			Paths: []string{
				"test/...",
				"some/",
				"./dir/...",
			},
			License:             "apache-v2",
			AddLicenseFile:      true,
			AddLicenseNotice:    false,
			ProjectRoot:         "/home/root/path/to/project/some_program",
			CustomLicenseNotice: "",
		},
		expected: ConfigErrYearsAuthors,
	},
	{
		cfg: Config{
			Authors:     []Author{
				{
					"John Smith",
					"2010",
				},
			},
			ProgramName:         "some program",
			License:             "apache-v2",
			AddLicenseFile:      true,
			AddLicenseNotice:    false,
			ProjectRoot:         "/home/root/path/to/project/some_program",
			CustomLicenseNotice: "",
		},
		expected: ConfigErrPathsRequired,
	},
	{
		cfg: Config{
			Authors:     []Author{
				{
					"John Smith",
					"2010",
				},
			},
			ProgramName: "some program",
			Paths: []string{
				"test/...",
				"some/",
				"./dir/...",
			},
			AddLicenseFile:      true,
			AddLicenseNotice:    false,
			ProjectRoot:         "/home/root/path/to/project/some_program",
			CustomLicenseNotice: "",
		},
		expected: ConfigErrLicenseRequired,
	},
	{
		cfg: Config{
			Authors:     []Author{
				{
					"John Smith",
					"2010",
				},
			},
			ProgramName: "some program",
			Paths: []string{
				"test/...",
				"some/",
				"./dir/...",
			},
			License:             "apache-v2",
			AddLicenseFile:      true,
			AddLicenseNotice:    false,
			CustomLicenseNotice: "",
		},
		expected: ConfigErrProjectRootRequired,
	},
	{
		cfg: Config{
			Authors:     []Author{
				{
					"John Smith",
					"2010",
				},
			},
			ProgramName: "some program",
			Paths: []string{
				"test/...",
				"some/",
				"./dir/...",
			},
			License:             "apache-v2",
			ProjectRoot:         "/home/root/path/to/project/some_program",
			CustomLicenseNotice: "",
		},
		expected: ConfigErrAddLicenseFileNoticeRequired,
	},
}

func Test_validate_Err(test *testing.T) {
	for _, data := range Test_validate_ErrData {
		if actual := data.cfg.validate(); actual == nil {
			test.Errorf("config.Test_validate_Err: actual != expected:\n\t%s != %s\n", actual, data.expected)
		}
	}
}
