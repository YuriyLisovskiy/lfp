// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"errors"

	"github.com/YuriyLisovskiy/lfp/src/static"
)

var (
	// Configuration errors
	ConfigErrYearRequired                 = errors.New(static.PROGRAM_NAME + " error: year required")
	ConfigErrAuthorRequired               = errors.New(static.PROGRAM_NAME + " error: author required")
	ConfigErrPathsRequired                = errors.New(static.PROGRAM_NAME + " error: path(s) required")
	ConfigErrLicenseRequired              = errors.New(static.PROGRAM_NAME + " error: license required")
	ConfigErrProjectRootRequired          = errors.New(static.PROGRAM_NAME + " error: project root required")
	ConfigErrAddLicenseFileNoticeRequired = errors.New(static.PROGRAM_NAME + " error: license file option of license notice option required")

	// Other errors
	ErrIndexOutOfRange       = errors.New("error: index out of range")
	ErrOldNewCountInvalidLen = errors.New("old, new and count must have the same len")
	ErrLicenseNotFound       = errors.New(static.PROGRAM_NAME + " error: license not found")
	ErrPathDoesNotExist      = errors.New(static.PROGRAM_NAME + " error: path does not exist")
	ErrMissingConfigPath     = errors.New(static.PROGRAM_NAME + " error: missing config path")
	ErrHelpRedundantArgs     = errors.New(static.PROGRAM_NAME + " error: help was called with redundant argument(s)")
	ErrVersionRedundantArgs  = errors.New(static.PROGRAM_NAME + " error: version was called with redundant argument(s)")
)
