// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package lfp

import (
	"testing"

	"github.com/YuriyLisovskiy/lfp/src/lfp/licenses"
	"github.com/YuriyLisovskiy/lfp/src/lfp/licenses/bsd"
	"github.com/YuriyLisovskiy/lfp/src/lfp/licenses/gnu"
)

var Test_getLicenseData = []struct {
	input    string
	expected map[string]string
}{
	{
		"apache-2.0",
		licenses.APACHE_LICENSE_2_0,
	},
	{
		"bsd-2-clause",
		bsd.BSD_2_CLAUSE_LICENSE,
	},
	{
		"bsd-3-clause",
		bsd.BSD_3_CLAUSE_LICENSE,
	},
	{
		"epl-2.0",
		licenses.ECLIPSE_PUBLIC_LICENSE_V2_0,
	},
	{
		"gpl-2.0",
		gnu.GNU_GPL_V2_0,
	},
	{
		"gpl-3.0",
		gnu.GNU_GPL_V3_0,
	},
	{
		"agpl-3.0",
		gnu.GNU_AFFERO_GPL_V_3,
	},
	{
		"lgpl-2.1",
		gnu.GNU_LESSER_GPL_V2_1,
	},
	{
		"lgpl-3.0",
		gnu.GNU_LESSER_GPL_V3,
	},
	{
		"mit",
		licenses.MIT_LICENSE,
	},
	{
		"mpl-2.0",
		licenses.MOZILLA_PUBLIC_LICENSE_V2,
	},
	{
		"unlicense",
		licenses.UNLICENSE,
	},
}

var utilTestErr = "util.%s: actual[\"%s\"] != expected[\"%s\"]:\n\t%s != %s\n"

func Test_getLicense(test *testing.T) {
	for _, data := range Test_getLicenseData {
		actual, _ := getLicense(data.input)
		if actual["name"] != data.expected["name"] {
			test.Errorf(utilTestErr, "Test_getLicense", "name", "name", actual["name"], data.expected["name"])
		}
		if actual["link"] != data.expected["link"] {
			test.Errorf(utilTestErr, "Test_getLicense", "link", "link", actual["link"], data.expected["link"])
		}
		if actual["text"] != data.expected["text"] {
			test.Errorf(utilTestErr, "Test_getLicense", "text", "text", actual["text"], data.expected["text"])
		}
	}
}

var Test_getLicenseErrLicenseNotFoundData = struct {
	input    string
	expected error
}{
	"some-unknown-license",
	ErrLicenseNotFound,
}

func Test_getLicenseErrLicenseNotFound(test *testing.T) {
	data := Test_getLicenseErrLicenseNotFoundData
	_, err := getLicense(data.input)
	if err != ErrLicenseNotFound {
		test.Errorf("util.Test_getLicenseErrLicenseNotFound: actual error != expected error:\n\t%s != %s", err, data.expected)
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
