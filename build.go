// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package main

import (
	"runtime"
	"os/exec"
)

func main() {
	platform := runtime.GOOS
	binary := "lofp"
	switch platform {
	case "windows":
		binary += ".exe"
	default:
	}
	exec.Command("go", "build", "-o", "bin/" + binary, "main.go").Run()
}
