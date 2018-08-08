// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package lfp

import "flag"

var (
	lfp       = flag.NewFlagSet("lfp", flag.ContinueOnError)
	helpPtr    = lfp.Bool("help", false, "prints usage")
	configPtr  = lfp.String("config", "", "set config path")
	versionPtr = lfp.Bool("version", false, "prints version")
)

// validateArgs checks if args is correct.
func validateArgs(args []string) error {
	if *helpPtr && len(args) > 2 {
		return ErrHelpRedundantArgs
	}
	if *versionPtr && len(args) > 2 {
		return ErrVersionRedundantArgs
	}
	if *configPtr == "" && !*helpPtr && !*versionPtr {
		return ErrMissingConfigPath
	}
	return nil
}

// validatePaths checks if given paths contain arguments.
func validatePaths(paths []string) error {
	if len(paths) == 0 {
		return ConfigErrPathsRequired
	}
	for _, path := range paths {
		exists, err := pathExists(path)
		if err != nil {
			return err
		}
		if !exists {
			return ErrPathDoesNotExist
		}
	}
	return nil
}
