// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package updater

import (
	"os"
	"log"
	"fmt"
	"time"
	"runtime"

	"github.com/mholt/archiver"
	"github.com/gen2brain/beeep"
	"github.com/sevlyar/go-daemon"
	"github.com/YuriyLisovskiy/lfp/src/var"
)

func startDaemon(archiveLoc string) {
	cwd, err := getCWD()
	if err != nil {
		log.Println(_var.PROGRAM_NAME + ": error: can't get current working directory")
	}
	cntxt := &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0644,
		LogFileName: "log",
		LogFilePerm: 0640,
		WorkDir:     cwd,
		Umask:       027,
		Args:        []string{"[lfp-updater]"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal(_var.PROGRAM_NAME + ": error: unable to run updater")
	}
	if d != nil {
		return
	}
	defer cntxt.Release()
	lfpLoc, err := os.Executable()
	if err != nil {
		log.Fatal(_var.PROGRAM_NAME + ": error: unable to get executable location")
	}
	go installUpdate(archiveLoc, lfpLoc)
}

// installUpdate installs downloaded update using daemon process
func installUpdate(path, exec string) {
	time.Sleep(1 * time.Second)
	start := time.Now()
	err := beeep.Notify("LFP Updater", "Updating LFP tool...", "")
	if err != nil {
		log.Fatal(err)
	}
	targetOs := runtime.GOOS
	switch targetOs {
	case "windows":
		err := archiver.Zip.Open(path, exec)
		if err != nil {
			err := beeep.Notify("LFP Updater", "Error occurred while opening downloaded archive with executable", "")
			if err != nil {
				log.Fatal(err)
			}
		}
	case "linux":
		err := archiver.TarGz.Open(path, exec)
		if err != nil {
			err := beeep.Notify("LFP Updater", "Error occurred while opening downloaded archive with executable", "")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	elapsed := time.Since(start)
	err = beeep.Notify(
		"LFP Updater",
		fmt.Sprintf("LFP tool has been updated successfully\nTime elapsed: %d sec", int64(elapsed * time.Second)),
		"",
	)
	if err != nil {
		log.Fatal(err)
	}
}
