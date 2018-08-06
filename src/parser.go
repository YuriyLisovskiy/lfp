// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"fmt"
	"github.com/YuriyLisovskiy/lofp/src/static"
)

func processPaths(paths []string, license, root string, addNotice, addLicense bool) {

}

func processPath(path string) {

}

func transform(fileContent []byte, license, author string, year int) ([]byte, error) {
	l, err := getLicense(license)
	if err != nil {
		return []byte{}, err
	}
	return append(
		[]byte(fmt.Sprintf(static.LICENSE_NOTICE_TEMPLATE, year, author, l["name"], l["link"])),
		fileContent...
	), nil
}
