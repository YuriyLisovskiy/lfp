// Copyright (c) 2018 Yuriy Lisovskiy
//
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package lfp

import "testing"

var Test_getLicenseErrLicenseNotFoundData = struct {
	input    string
}{
	"some-unknown-license",
}

func Test_getLicenseErrLicenseNotFound(test *testing.T) {
	data := Test_getLicenseErrLicenseNotFoundData
	_, err := getLicense(data.input)
	if err == nil {
		test.Errorf("util.Test_getLicenseErrLicenseNotFound: func does not return an error for input %s", data.input)
	}
}

var Test_findIndentReverse_Data = []struct {
	input string
	expected string
}{
	{
		input: "Hello, world\n       ",
		expected: "       ",
	},
	{
		input: "Hello, world                \n",
		expected: "",
	},
	{
		input: "",
		expected: "",
	},
}

func Test_findIndentReverse(test *testing.T) {
	for _, data := range Test_findIndentReverse_Data {
		actual := findIndentReverse(data.input)
		if actual != data.expected {
			test.Errorf("util.Test_findIndentReverse: actual != expected:\n\t%s != %s", actual, data.expected)
		}
	}
}
