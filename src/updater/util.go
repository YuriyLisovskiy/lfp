// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"os"
	"path/filepath"
)

func getCWD() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
