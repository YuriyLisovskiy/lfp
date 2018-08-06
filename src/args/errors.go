// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package args

import (
	"errors"
	"github.com/YuriyLisovskiy/lofp/src/static"
)

var (
	ErrPathsRequired               = errors.New(static.PROGRAM_NAME + " error: path(s) required")
	ErrLicenseNotFound             = errors.New(static.PROGRAM_NAME + " error: license not found")
	ErrProjectRootRequired         = errors.New(static.PROGRAM_NAME + " error: project root required")
	ErrPathsContainArgs            = errors.New(static.PROGRAM_NAME + " error: paths contain argument(s)")
	ErrLicenseNotSpecified         = errors.New(static.PROGRAM_NAME + " error: license was not specified")
	ErrHelpRedundantArgs           = errors.New(static.PROGRAM_NAME + " error: help was called with redundant argument(s)")
	ErrVersionRedundantArgs        = errors.New(static.PROGRAM_NAME + " error: version was called with redundant argument(s)")
	ErrLicensingOptionNotSpecified = errors.New(static.PROGRAM_NAME + " error: at least one of -l or -n arguments required, type -h for details")
)
