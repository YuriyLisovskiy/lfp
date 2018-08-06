// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"testing"

	"github.com/YuriyLisovskiy/lofp/src/licenses"
	"github.com/YuriyLisovskiy/lofp/src/licenses/bsd"
	"github.com/YuriyLisovskiy/lofp/src/licenses/gnu"
)

var Test_getLicenseData = []struct {
	input    string
	expected map[string]string
} {
	{
		"apache-v2",
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
		"eclipse-pl-v2",
		licenses.ECLIPSE_PUBLIC_LICENSE_V2_0,
	},
	{
		"gnu-gpl-v2",
		gnu.GNU_GPL_V2_0,
	},
	{
		"gnu-gpl-v3",
		gnu.GNU_GPL_V3_0,
	},
	{
		"gnu-affero-gpl-v3",
		gnu.GNU_AFFERO_GPL_V_3,
	},
	{
		"gnu-lesser-gpl-v2.1",
		gnu.GNU_LESSER_GPL_V2_1,
	},
	{
		"gnu-lesser-gpl-v3",
		gnu.GNU_LESSER_GPL_V3,
	},
	{
		"mit",
		licenses.MIT_LICENSE,
	},
	{
		"mozilla-pl-v2",
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
	input string
	expected error
}{
	"some-unknown-license",
	ErrLicenseNotFound,
}

func Test_getLicenseErrLicenseNotFound(test *testing.T) {
	_, err := getLicense(Test_getLicenseErrLicenseNotFoundData.input)
	if err != ErrLicenseNotFound {
		test.Errorf("util.Test_getLicenseErrLicenseNotFound: actual error != expected error:\n\t%s != %s",
			err,
			Test_getLicenseErrLicenseNotFoundData.expected,
		)
	}
}
