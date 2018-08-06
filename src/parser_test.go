// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"testing"

	"github.com/YuriyLisovskiy/lofp/src/args"
)

var Test_transformData = []struct {
	content  []byte
	license  string
	author   string
	year     int
	expected []byte
}{
	{
		content: []byte("package main\n\nfunc main() {\n\n}"),
		license: "mit",
		author: "John Smith",
		year: 2010,
		expected: []byte("// Copyright (c) 2010 John Smith\n// Distributed under the MIT License,\n// see the accompanying file LICENSE or https://opensource.org/licenses/MIT\n\npackage main\n\nfunc main() {\n\n}"),
	},
}

func Test_transform(test *testing.T) {
	for _, data := range Test_transformData {
		actual, _ := transform(data.content, data.license, data.author, data.year)
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
	license  string
	author   string
	year     int
	expected error
}{
	{
		content: []byte("package main\n\nfunc main() {\n\n}"),
		license: "some-unknown-license",
		author: "John Smith",
		year: 2010,
		expected: args.ErrLicenseNotFound,
	},
}

func Test_transformErrLicenseNotFound(test *testing.T) {
	for _, data := range Test_transformDataErrLicenseNotFound {
		_, err := transform(data.content, data.license, data.author, data.year)
		if err != data.expected {
			test.Errorf("parser.Test_transformErrLicenseNotFound: actual != expected:\n\t%s != %s\n", err, data.expected)
		}
	}
}
