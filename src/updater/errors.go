// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"errors"

	)

var (
	ErrTooManyArguments     = errors.New(updater + " error: too many arguments")
	ErrVersionNotFound      = errors.New(updater + ": error: version not found")
	ErrReleasesNotFound     = errors.New(updater + ": error: releases not found")
	ErrCantFetchReleases    = errors.New(updater + ": error: can't fetch program releases")
	ErrNoReleaseForTargetOs = errors.New(updater + ": error: no suitable release for target operating system")
)
