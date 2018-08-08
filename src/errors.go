// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"errors"
	"github.com/YuriyLisovskiy/lfp/src/var"
)

var (
	// Configuration errors
	ConfigErrPathsRequired                = errors.New(_var.PROGRAM_NAME + " error: path(s) required")
	ConfigErrLicenseRequired              = errors.New(_var.PROGRAM_NAME + " error: license required")
	ConfigErrYearsAuthors                 = errors.New(_var.PROGRAM_NAME + " error: year and(or) author is empty")
	ErrCantRetrieveUserName               = errors.New(_var.PROGRAM_NAME + " error: can't get user name for licensing")
	ConfigErrAddLicenseFileNoticeRequired = errors.New(_var.PROGRAM_NAME + " error: license file option of license notice option required")

	// Other errors
	ErrIndexOutOfRange      = errors.New("error: index out of range")
	ErrCommentNotFound      = errors.New(_var.PROGRAM_NAME + " error: comment not found")
	ErrLicenseNotFound      = errors.New(_var.PROGRAM_NAME + " error: license not found")
	ErrPathDoesNotExist     = errors.New(_var.PROGRAM_NAME + " error: path does not exist")
	ErrMissingConfigPath    = errors.New(_var.PROGRAM_NAME + " error: missing config path")
	ErrTooManyArguments     = errors.New(_var.PROGRAM_NAME + " error: too many arguments")
)
