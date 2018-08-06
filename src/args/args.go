// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package args

import (
	"flag"
	"strings"
)

var (
	Lofp = flag.NewFlagSet("lofp", flag.ContinueOnError)

	pathsPtr       = Lofp.String("p", "", "sets paths")
	licenseTypePtr = Lofp.String("l", "", "specify license")
	rootPtr        = Lofp.String("r", "", "sets project root")

	setLicensePtr = Lofp.Bool("lf", false, "adds LICENSE file to the project root folder")
	setNoticePtr  = Lofp.Bool("n", false, "adds license notice to specified files")
	helpPtr       = Lofp.Bool("h", false, "prints usage")
	versionPtr    = Lofp.Bool("v", false, "prints version")
)

// validateArgs checks if args is correct.
func ValidateArgs(args []string) error {
	if *helpPtr && len(args) > 2 {
		return ErrHelpRedundantArgs
	}
	if *versionPtr && len(args) > 2 {
		return ErrVersionRedundantArgs
	}
	if *licenseTypePtr == "" {
		return ErrLicenseNotSpecified
	}
	if *rootPtr == "" {
		return ErrProjectRootRequired
	}
	if !*setLicensePtr && !*setNoticePtr {
		return ErrLicensingOptionNotSpecified
	}
	if *pathsPtr == "" {
		return ErrPathsRequired
	}
	return nil
}

// isArg checks if given string is an argument.
func isArg(arg string) bool {
	switch strings.Replace(arg, "-", "", -1) {
	case "p", "lf", "l", "n", "r", "h", "v":
		return true
	default:
		return false
	}
}

// validatePaths checks if given paths contain arguments.
func ValidatePaths(paths []string) error {
	if len(paths) == 0 {
		return ErrPathsRequired
	}
	for _, path := range paths {
		if isArg(path) {
			return ErrPathsContainArgs
		}
	}
	return nil
}

// retrievePaths finds -d argument, cuts and returns the rest arguments as paths.
func RetrievePaths(args []string) []string {
	var ret []string

	// find -d flag to get beginning of paths array
	for i, arg := range args {
		if strings.Replace(arg, "-", "", -1) == "d" {
			ret = args[i+1:]
			break
		}
	}

	// find any other flag to get ending of paths array
	for i, item := range ret {
		if isArg(item) {
			ret = ret[:i]
			break
		}
	}
	return ret
}
