// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import "os"

func RunCLI() error {

	// parse command line arguments
	if err := lofp.Parse(os.Args[1:]); err != nil {
		return err
	}
	paths := lofp.Args()

	// check if there is no errors in given arguments
	if err := validateArgs(paths); err != nil {
		return err
	}

	// clean up arguments
	cleanUpArgs()

	for _, path := range paths {
		println(path)
	}

	return nil
}
