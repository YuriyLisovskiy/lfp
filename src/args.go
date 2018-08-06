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
		if (path) {
			return ErrPathsContainArgs
		}
	}
	return nil
}
