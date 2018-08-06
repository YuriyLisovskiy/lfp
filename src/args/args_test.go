// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package args

import "testing"

var Test_isArgData = []struct {
	input    string
	expected bool
}{
	{
		"-p",
		true,
	},
	{
		"-lf",
		true,
	},
	{
		"-l",
		true,
	},
	{
		"-n",
		true,
	},
	{
		"-r",
		true,
	},
	{
		"-h",
		true,
	},
	{
		"-v",
		true,
	},
	{
		"unknown-argument",
		false,
	},
}

func Test_isArg(test *testing.T) {
	for _, data := range Test_isArgData {
		actual := isArg(data.input)
		if actual != data.expected {
			test.Errorf("args.Test_isArg: actual != expected:\n\t%t != %t\n", actual, data.expected)
		}
	}
}


var Test_ValidatePathsData = struct {
	input    []string
	expected error
}{
	[]string{"/home/", "/home/username/", "./"},
	nil,
}

func Test_ValidatePaths(test *testing.T) {
	actual := ValidatePaths(Test_ValidatePathsData.input)
	if actual != Test_ValidatePathsData.expected {
		test.Errorf("args.Test_ValidatePaths: actual != expected:\n\t%s != %s\n", actual, Test_ValidatePathsData.expected)
	}
}


var Test_ValidatePathsDataErrPathsContainArgs = struct {
	input    []string
	expected error
}{
	[]string{"/home/", "-v", "/home/username/", "./"},
	ErrPathsContainArgs,
}

func Test_validatePathsErrPathsContainArgs(test *testing.T) {
	actual := ValidatePaths(Test_ValidatePathsDataErrPathsContainArgs.input)
	if actual != Test_ValidatePathsDataErrPathsContainArgs.expected {
		test.Errorf("args.Test_ValidatePathsErrPathsContainArgs: actual != expected:\n\t%s != %s\n",
			actual, Test_ValidatePathsDataErrPathsContainArgs.expected,
		)
	}
}


var Test_ValidatePathsDataErrPathsRequired = struct {
	input    []string
	expected error
}{
	[]string{},
	ErrPathsRequired,
}

func Test_ValidatePathsErrPathsRequired(test *testing.T) {
	actual := ValidatePaths(Test_ValidatePathsDataErrPathsRequired.input)
	if actual != Test_ValidatePathsDataErrPathsRequired.expected {
		test.Errorf("args.Test_ValidatePathsErrPathsRequired: actual != expected:\n\t%s != %s\n",
			actual, Test_ValidatePathsDataErrPathsRequired.expected,
		)
	}
}


var Test_RetrievePathsData = []struct {
	input    []string
	expected []string
}{
	{
		[]string{"-d", "/home/", "/home/username/", "./", "-n"},
		[]string{"/home/", "/home/username/", "./"},
	},
	{
		[]string{"-l", "-n", "-d", "/home/", "/home/username/", "./"},
		[]string{"/home/", "/home/username/", "./"},
	},
	{
		[]string{"-l", "-n", "-d", "/home/", "/home/username/", "./", "-lf"},
		[]string{"/home/", "/home/username/", "./"},
	},
}

func Test_RetrievePaths(test *testing.T) {
	for i, data := range Test_RetrievePathsData {
		actual := RetrievePaths(data.input)
		for j, item := range actual {
			if item != data.expected[j] {
				test.Errorf("args.Test_RetrievePaths: actual != expected, item %d:\n\t%s != %s\n",
					i, item, data.expected[j],
				)
			}
		}
	}
}
