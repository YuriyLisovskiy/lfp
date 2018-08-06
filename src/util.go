// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"os"

	"github.com/YuriyLisovskiy/lofp/src/licenses"
	"github.com/YuriyLisovskiy/lofp/src/licenses/bsd"
	"github.com/YuriyLisovskiy/lofp/src/licenses/gnu"
	"io/ioutil"
	"fmt"
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
	case "apache-v2", "mit", "bsd-2-clause", "bsd-3-clause":
		licenseContent = fmt.Sprintf(licenseContent, cfg.Year, cfg.Author)
	case "gnu-lesser-gpl-v2.1", "gnu-gpl-v2":
		licenseContent = fmt.Sprintf(licenseContent, cfg.ProgramName, cfg.Year, cfg.Author, cfg.ProgramName, cfg.Year, cfg.Author)
	case "gnu-gpl-v3", "gnu-affero-gpl-v3":
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
