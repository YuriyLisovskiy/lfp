// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package updater

import "flag"

var (
	lfpUpdater = flag.NewFlagSet("lfpupdater", flag.ContinueOnError)
	updatePtr  = lfpUpdater.String("update", "", "update LFP tool")
	helpPtr    = lfpUpdater.Bool("help", false, "print help")
)

// validateArgs checks if args is correct.
func validateArgs(args []string) error {
	if *helpPtr {
		if len(args) > 2 {
			return ErrTooManyArguments
		}
	}
	if *updatePtr != "" {
		if len(args) > 3 {
			return ErrTooManyArguments
		}
	}
	return nil
}
