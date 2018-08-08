// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import "testing"

var Test_getComments_Data = []struct {
	inputs    []string
	expected1 string
	expected2 string
}{
	{
		inputs: []string{
			"d",    "cu",   "zig",  "php3",  "scala",  "groovy",   "postcss",  "fsscript",
			"h",    "uc",   "uci",  "php4",  "php-s",
			"m",    "as",   "css",  "php5",  "swift",
			"y",    "go",   "dts",  "php7",  "proto",
					"js",   "hpp",  "phps",  "phtml",
					"sc",   "hxx",  "scss",  "ccpph",
					"kt",   "cxx",  "sass",
					"mm",   "c++",  "dtsi",
					"gy",   "cpp",  "tesc",
					"fs",   "gvy",  "tese",
					"cp",   "gsh",  "geom",
					"hh",   "fsi",  "frag",
					"cc",   "fsx",  "comp",
					"rs",   "kts",  "vert",
					"ts",	"php",  "upkg",
							"cfc",  "java",
							"cuh",  "json",
							"jai",  "styl",
							"jsx",  "dart",
							"lds",  "less",
							"qcl",
							"qml",
							"tsx",
		},
		expected1: "//",
		expected2: "",
	},
	{
		inputs: []string{
			"md",   "mkd",  "text",  "mdown",  "mdtext",  "markdown",  "handlebars",
					"hbs",  "mkdn",  "mdtxt",  "cshtml",
					"rmd",  "mdwn",  "polly",
					"xml",  "html",  "rhtml",
					"htm",           "xhtml",
		},
		expected1: "<!--",
		expected2: "-->",
	},
	{
		inputs: []string{
			"t",    "pm",   "yml",  "yaml",  "rdata",  "coffee",  "makefile",
			"r",    "pl",   "pod",  "toml",
			"s",    "py",   "rds",
					"rb",   "rda",
					"jl",   "asm",
					"tf",   "awk",
					"in",   "csh",
					"sh",   "nim",
							"sls",
							"tcl",
							"nix",
							"arr",
		},
		expected1: "#",
		expected2: "",
	},
	{
		inputs: []string{
			"hs",   "lhs",  "lagda",   "agda",
					"adb",  "hlean",   "lidr",
					"ads",             "lean",
					"idr",
					"sql",
					"lua",
		},
		expected1: "--",
		expected2: "",
	},
	{
		inputs: []string{
			"axd",  "asax", "vbhtml",
			"svc",  "ascx",
			"asp",  "ashx",
					"asmx",
					"dbml",
					"edmx",
					"resx",
					"aspx",
		},
		expected1: "<%--",
		expected2: "--%>",
	},
	{
		inputs: []string{
			"pp",   "inc",
					"pas",
		},
		expected1: "{*",
		expected2: "*}",
	},
	{
		inputs: []string{
			"btm",
			"bat",
			"cmd",
		},
		expected1: "REM",
		expected2: "",
	},
	{
		inputs: []string{
			"p",    "oz",   "hrl",
							"erl",
							"pro",
							"sty",
							"tex",
		},
		expected1: "%",
		expected2: "",
	},
	{
		inputs: []string{
			"fr",   "4th",  "forth",
			"fb",   "frt",
			"e4",   "fth",
			"rx",   "f83",
			"ft",   "fpm",
		},
		expected1: "(",
		expected2: ")",
	},
	{
		inputs: []string{
			"f",    "for",
					"ftn",
					"f77",
					"pfo",
					"f08",
					"f90",
					"f95",
					"f03",
		},
		expected1: "!",
		expected2: "",
	},
	{
		inputs: []string{
			"el",   "lsp",  "lisp",
			"ss",   "scm",
					"rkt",
					"ini",
		},
		expected1: ";",
		expected2: "",
	},
	{
		inputs: []string{
			"v",    "ml",   "mli",
					"nb",   "sml",
					"wl",
		},
		expected1: "(*",
		expected2: "*)",
	},
	{
		inputs: []string{
			"txt",  "ihex",
			"hex",
			"rst",
		},
		expected1: "",
		expected2: "",
	},
	{
		inputs: []string{
			"mustache",
		},
		expected1: "{{!",
		expected2: "}}",
	},
}

func Test_getComments(test *testing.T) {
	for _, data := range Test_getComments_Data {
		for _, input := range data.inputs {
			actual1, actual2, _ := getComments(input)
			if actual1 != data.expected1 {
				test.Errorf("extension.Test_getComments: actual != expected:\n\t%s != %s\n", actual1, data.expected1)
			}
			if actual2 != data.expected2 {
				test.Errorf("extension.Test_getComments: actual != expected:\n\t%s != %s\n", actual2, data.expected2)
			}
		}
	}
}

var Test_getExtension_Data = []struct{
	input string
	expected string
} {
	{
		input: "path/to/file.java",
		expected: "java",
	},
	{
		input: "path/to/Makefile",
		expected: "Makefile",
	},
	{
		input: "some_file.go",
		expected: "go",
	},
	{
		input: "some_file",
		expected: "some_file",
	},
}

func Test_getExtension(test *testing.T) {
	for _, data := range Test_getExtension_Data {
		actual := getExtension(data.input)
		if actual != data.expected {
			test.Errorf("extension.Test_getExtension: actual != expected:\n\t%s != %s\n", actual, data.expected)
		}
	}
}
