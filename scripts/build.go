// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"bytes"
	"runtime"
	"os/exec"
)

func main() {
	platform := runtime.GOOS
	binary := "lfp"
	switch platform {
	case "windows":
		binary += ".exe"
	default:
	}
	cmd := exec.Command("go", "build", "-o", "bin/" + binary, "main.go")
	var out bytes.Buffer
	cmd.Stderr = &out
	cmd.Run()
	if out.String() != "" {
		fmt.Println(out.String())
	}
}
