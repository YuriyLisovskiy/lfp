// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package updater

import "os"

func RunCLI() error {
	// if there are no arguments
	if len(os.Args) == 1 {
		print(about)
		lfpUpdater.Usage()
		return nil
	}

	if *helpPtr {
		lfpUpdater.Usage()
		return nil
	}

	// parse command line arguments
	if err := lfpUpdater.Parse(os.Args[1:]); err != nil {
		return nil
	}
	arguments := lfpUpdater.Args()

	// check if there is no errors in given arguments
	if err := validateArgs(arguments); err != nil {
		return err
	}
	return processUpdate()
}

func processUpdate() error {
	return startUpdate(*updatePtr)
}
