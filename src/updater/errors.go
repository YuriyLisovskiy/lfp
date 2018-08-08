// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package updater

import (
	"errors"

	"github.com/YuriyLisovskiy/lfp/src/var"
)

var (
	ErrVersionNotFound      = errors.New(_var.PROGRAM_NAME + ": update error: version not found")
	ErrReleasesNotFound     = errors.New(_var.PROGRAM_NAME + ": update error: releases not found")
	ErrCantFetchReleases    = errors.New(_var.PROGRAM_NAME + ": update error: can't fetch program releases")
	ErrNoReleaseForTargetOs = errors.New(_var.PROGRAM_NAME + ": update error: no suitable release for target operating system")
)
