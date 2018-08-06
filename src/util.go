// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"github.com/YuriyLisovskiy/lofp/src/args"
	"github.com/YuriyLisovskiy/lofp/src/licenses"
	"github.com/YuriyLisovskiy/lofp/src/licenses/bsd"
	"github.com/YuriyLisovskiy/lofp/src/licenses/gnu"
	)

func getLicense(license string) (map[string]string, error) {
	var res map[string]string
	switch license {
	case "apache-v2":
		res = licenses.APACHE_LICENSE_2_0
	case "bsd-2-clause":
		res = bsd.BSD_2_CLAUSE_LICENSE
	case "bsd-3-clause":
		res = bsd.BSD_3_CLAUSE_LICENSE
	case "eclipse-pl-v2":
		res = licenses.ECLIPSE_PUBLIC_LICENSE_V2_0
	case "gnu-gpl-v2":
		res = gnu.GNU_GPL_V2_0
	case "gnu-gpl-v3":
		res = gnu.GNU_GPL_V3_0
	case "gnu-affero-gpl-v3":
		res = gnu.GNU_AFFERO_GPL_V_3
	case "gnu-lesser-gpl-v2.1":
		res = gnu.GNU_LESSER_GPL_V2_1
	case "gnu-lesser-gpl-v3":
		res = gnu.GNU_LESSER_GPL_V3
	case "mit":
		res = licenses.MIT_LICENSE
	case "mozilla-pl-v2":
		res = licenses.MOZILLA_PUBLIC_LICENSE_V2
	case "unlicense":
		res = licenses.UNLICENSE
	default:
		return map[string]string{}, args.ErrLicenseNotFound
	}
	return res, nil
}

func createLicenseFile() {

}
