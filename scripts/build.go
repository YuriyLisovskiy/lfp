// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"bytes"
	"os/exec"
	"runtime"
)

func main() {
	platform := runtime.GOOS
	binaryLfp := "lfp"
	binaryUpdater := "lfp-updater"
	switch platform {
	case "windows":
		binaryLfp += ".exe"
		binaryUpdater += ".exe"
	default:
	}
	build("bin/" + binaryLfp, "./src/main-lfp.go")
	build("bin/" + binaryUpdater, "./src/main-updater.go")
}

func build(destination, source string) {
	cmd := exec.Command("go", "build", "-o", destination, source)
	var out bytes.Buffer
	cmd.Stderr = &out
	cmd.Run()
	if out.String() != "" {
		fmt.Println(out.String())
	}
}
