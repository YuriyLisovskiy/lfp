// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"flag"
	"errors"
	"strings"
)

var (
	lofp = flag.NewFlagSet("lofp", flag.ContinueOnError)
	setLicensePtr = lofp.Bool("l", false, "add LICENSE file to the project root folder")
	setNoticePtr = lofp.Bool("n", false, "add license notice to specified files")
	rootPtr = lofp.String("r", "", "set project root")
)

func validateArgs(paths []string) error {
	if len(paths) == 0 {
		return errors.New("path(s) required")
	}
	if *rootPtr == "" {
		return errors.New("project root required")
	}
	if !*setLicensePtr && !*setNoticePtr {
		return errors.New("at least one of -l or -n arguments required, type -h for details")
	}
	return nil
}

func cleanUpArgs() {
	*rootPtr = strings.TrimSpace(*rootPtr)
}
