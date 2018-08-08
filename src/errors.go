// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import "errors"

var (
	// Configuration errors
	ConfigErrAuthorRequired               = errors.New(PROGRAM_NAME + " error: author required")
	ConfigErrPathsRequired                = errors.New(PROGRAM_NAME + " error: path(s) required")
	ConfigErrLicenseRequired              = errors.New(PROGRAM_NAME + " error: license required")
	ConfigErrYearsAuthors                 = errors.New(PROGRAM_NAME + " error: year and(or) author is empty")
	ErrCantRetrieveUserName               = errors.New(PROGRAM_NAME + " error: can't get user name for licensing")
	ConfigErrAddLicenseFileNoticeRequired = errors.New(PROGRAM_NAME + " error: license file option of license notice option required")

	// Other errors
	ErrIndexOutOfRange      = errors.New("error: index out of range")
	ErrCommentNotFound      = errors.New(PROGRAM_NAME + " error: comment not found")
	ErrLicenseNotFound      = errors.New(PROGRAM_NAME + " error: license not found")
	ErrPathDoesNotExist     = errors.New(PROGRAM_NAME + " error: path does not exist")
	ErrMissingConfigPath    = errors.New(PROGRAM_NAME + " error: missing config path")
	ErrHelpRedundantArgs    = errors.New(PROGRAM_NAME + " error: help was called with redundant argument(s)")
	ErrVersionRedundantArgs = errors.New(PROGRAM_NAME + " error: version was called with redundant argument(s)")
)
