// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"testing"

	"github.com/YuriyLisovskiy/lfp/src/licenses"
	"github.com/YuriyLisovskiy/lfp/src/licenses/bsd"
	"github.com/YuriyLisovskiy/lfp/src/licenses/gnu"
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
	_, err := getLicense(Test_getLicenseErrLicenseNotFoundData.input)
	if err != ErrLicenseNotFound {
		test.Errorf("util.Test_getLicenseErrLicenseNotFound: actual error != expected error:\n\t%s != %s",
			err,
			Test_getLicenseErrLicenseNotFoundData.expected,
		)
	}
}

var Test_aggregate_Data = []struct {
	input     []string
	separator string
	expected  string
}{
	{
		input:     []string{"some", "words", "to", "aggregate", "separated", "by", "comma"},
		separator: ", ",
		expected:  "some, words, to, aggregate, separated, by, comma",
	},
	{
		input:     []string{"some", "another", "words", "to", "aggregate"},
		separator: " ",
		expected:  "some another words to aggregate",
	},
	{
		input:     []string{},
		separator: ".",
		expected:  "",
	},
}

func Test_aggregate(test *testing.T) {
	for _, data := range Test_aggregate_Data {
		actual := aggregate(data.input, data.separator)
		if actual != data.expected {
			test.Errorf("util.Test_aggregate: actual != expected:\n\t%s != %s\n", actual, data.expected)
		}
	}
}

var Test_prepareLicense_Data = struct {
	old      []string
	new      []string
	count    []int
	template string
	expected string
}{
	template: "Some license template written by <authors> in <years> and one more time - <authors> =)",
	old:      []string{"<authors>", "<years>"},
	new: []string{
		aggregate([]string{"Author 1", "Author 2"}, ", "),
		"2018",
	},
	count:    []int{2, 1},
	expected: "Some license template written by Author 1, Author 2 in 2018 and one more time - Author 1, Author 2 =)",
}

func Test_prepareLicense(test *testing.T) {
	data := Test_prepareLicense_Data
	actual, _ := prepareLicense(data.template, data.old, data.new, data.count)
	if actual != data.expected {
		test.Errorf("util.Test_prepareLicense: actual != expected:\n\t%s != %s\n", actual, data.expected)
	}
}

var Test_prepareLicense_ErrOldNewCountInvalidLenData = []struct {
	old      []string
	new      []string
	count    []int
	template string
	expected error
}{
	{
		template: "Some license template written by <authors> in <years> and one more time - <authors> =)",
		old:      []string{"<years>"},
		new: []string{
			aggregate([]string{"Author 1", "Author 2"}, ", "),
			"2018",
		},
		count:    []int{2, 1},
		expected: ErrOldNewCountInvalidLen,
	},
	{
		template: "Some license template written by <authors> in <years> and one more time - <authors> =)",
		old:      []string{"<authors>", "<years>"},
		new: []string{
			"2018",
		},
		count:    []int{2, 1},
		expected: ErrOldNewCountInvalidLen,
	},
	{
		template: "Some license template written by <authors> in <years> and one more time - <authors> =)",
		old:      []string{"<authors>", "<years>"},
		new: []string{
			aggregate([]string{"Author 1", "Author 2"}, ", "),
			"2018",
		},
		count:    []int{2},
		expected: ErrOldNewCountInvalidLen,
	},
}

func Test_prepareLicense_ErrOldNewCountInvalidLen(test *testing.T) {
	for _, data := range Test_prepareLicense_ErrOldNewCountInvalidLenData {
		_, actual := prepareLicense(data.template, data.old, data.new, data.count)
		if actual != data.expected {
			test.Errorf("util.Test_prepareLicense: actual != expected:\n\t%s != %s\n", actual, data.expected)
		}
	}
}
