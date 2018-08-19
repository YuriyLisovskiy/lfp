// Copyright (c) 2018 Yuriy Lisovskiy
//
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package lfp

import "testing"

var Test_ValidatePathsData = struct {
	input    []string
	expected error
}{
	[]string{"/home/", "./"},
	nil,
}

func Test_ValidatePaths(test *testing.T) {
	actual := validatePaths(Test_ValidatePathsData.input)
	if actual != Test_ValidatePathsData.expected {
		test.Errorf("args.Test_ValidatePaths: actual != expected:\n\t%s != %s\n", actual, Test_ValidatePathsData.expected)
	}
}


var Test_ValidatePathsDataErrPathsRequired = struct {
	input    []string
	expected error
}{
	[]string{},
	ConfigErrPathsRequired,
}

func Test_ValidatePathsErrPathsRequired(test *testing.T) {
	actual := validatePaths(Test_ValidatePathsDataErrPathsRequired.input)
	if actual != Test_ValidatePathsDataErrPathsRequired.expected {
		test.Errorf("args.Test_ValidatePathsErrPathsRequired: actual != expected:\n\t%s != %s\n",
			actual, Test_ValidatePathsDataErrPathsRequired.expected,
		)
	}
}
