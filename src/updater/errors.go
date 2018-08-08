// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package updater

import "errors"

var (
	ErrTooManyArguments          = errors.New(updater + " error: too many arguments")
	ErrVersionNotFound           = errors.New(updater + ": error: version not found")
	ErrReleasesNotFound          = errors.New(updater + ": error: releases not found")
	ErrCantFetchReleases         = errors.New(updater + ": error: can't fetch program releases")
	ErrTheLatestAlreadyInstalled = errors.New(updater + ": you have the newest version of lfp tool")
	ErrNoReleaseForTargetOs      = errors.New(updater + ": error: no suitable release for target operating system")
	ErrCantRetrieveLfpVer        = errors.New(updater + ": fatal: can't retrieve current version of LFP tool, please update or reinstall lfp-updater")
	ErrLFPIsBroken               = errors.New(updater + ": fatal: you have unknown version of lfp tool, we recommend to reinstall it to avoid any bugs in production use")
)
