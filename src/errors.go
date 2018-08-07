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
	ConfigErrAuthorRequired               = errors.New(static.PROGRAM_NAME + " error: author required")
	ConfigErrPathsRequired                = errors.New(static.PROGRAM_NAME + " error: path(s) required")
	ConfigErrLicenseRequired              = errors.New(static.PROGRAM_NAME + " error: license required")
	ConfigErrProjectRootRequired          = errors.New(static.PROGRAM_NAME + " error: project root required")
	ConfigErrYearsAuthors                 = errors.New(static.PROGRAM_NAME + " error: year and(or) author is empty")
	ConfigErrAddLicenseFileNoticeRequired = errors.New(static.PROGRAM_NAME + " error: license file option of license notice option required")

	// Other errors
	ErrCommentNotFound      = errors.New("comment not found")
	ErrIndexOutOfRange      = errors.New("error: index out of range")
	ErrLicenseNotFound      = errors.New(static.PROGRAM_NAME + " error: license not found")
	ErrPathDoesNotExist     = errors.New(static.PROGRAM_NAME + " error: path does not exist")
	ErrMissingConfigPath    = errors.New(static.PROGRAM_NAME + " error: missing config path")
	ErrHelpRedundantArgs    = errors.New(static.PROGRAM_NAME + " error: help was called with redundant argument(s)")
	ErrVersionRedundantArgs = errors.New(static.PROGRAM_NAME + " error: version was called with redundant argument(s)")
)
