// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"os"
	"fmt"
	"io/ioutil"

	"github.com/YuriyLisovskiy/lfp/src/licenses"
	"github.com/YuriyLisovskiy/lfp/src/licenses/bsd"
	"github.com/YuriyLisovskiy/lfp/src/licenses/gnu"
)

func getLicense(license string) (map[string]string, error) {
	var res map[string]string
	switch license {
	case "apache-2.0":
		res = licenses.APACHE_LICENSE_2_0
	case "bsd-2-clause":
		res = bsd.BSD_2_CLAUSE_LICENSE
	case "bsd-3-clause":
		res = bsd.BSD_3_CLAUSE_LICENSE
	case "epl-2.0":
		res = licenses.ECLIPSE_PUBLIC_LICENSE_V2_0
	case "gpl-2.0":
		res = gnu.GNU_GPL_V2_0
	case "gpl-3.0":
		res = gnu.GNU_GPL_V3_0
	case "agpl-3.0":
		res = gnu.GNU_AFFERO_GPL_V_3
	case "lgpl-2.1":
		res = gnu.GNU_LESSER_GPL_V2_1
	case "lgpl-3.0":
		res = gnu.GNU_LESSER_GPL_V3
	case "mit":
		res = licenses.MIT_LICENSE
	case "mpl-2.0":
		res = licenses.MOZILLA_PUBLIC_LICENSE_V2
	case "unlicense":
		res = licenses.UNLICENSE
	default:
		return map[string]string{}, ErrLicenseNotFound
	}
	return res, nil
}

func createLicenseFile(cfg Config) error {

	path := cfg.ProjectRoot + "/LICENSE"

	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	license, err := getLicense(cfg.License)
	if err != nil {
		return err
	}
	licenseContent := license["text"]
	switch cfg.License {
	case "apache-2.0", "mit", "bsd-2-clause", "bsd-3-clause":
		licenseContent = fmt.Sprintf(licenseContent, cfg.Year, cfg.Author)
	case "lgpl-2.1", "gpl-2.0":
		licenseContent = fmt.Sprintf(licenseContent, cfg.ProgramName, cfg.Year, cfg.Author, cfg.ProgramName, cfg.Year, cfg.Author)
	case "gpl-3.0", "agpl-3.0":
		licenseContent = fmt.Sprintf(licenseContent, cfg.ProgramName, cfg.Year, cfg.Author)
	default:
	}

	// write to file generated license
	err = ioutil.WriteFile(path, []byte(licenseContent), 0644)
	if err != nil {
		return err
	}
	return nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode().IsDir()
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return true
	}
	return info.Mode().IsRegular()
}
