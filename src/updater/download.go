// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package updater

import (
	"io"
	"os"
	"log"
	"fmt"
	"time"
	"strconv"
	"net/http"
)

func printStatus(percent int64) {
	fmt.Print("\r[")
	for i := int64(0); i < 100 / 2; i++ {
		if i < percent / 2 {
			fmt.Print("#")
		} else if i == percent / 2 {
			fmt.Print(">")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("] %d", int64(percent))
	fmt.Print("%")
}

func printDownloadPercent(done chan int64, path string, total int64) {
	stop := false
	for {
		select {
		case <-done:
			stop = true
		default:
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			fi, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}
			file.Close()
			size := fi.Size()
			if size == 0 {
				size = 1
			}
			percent := float64(size) / float64(total) * 100
			printStatus(int64(percent))
		}
		if stop {
			printStatus(100)
			fmt.Println("\nDownload completed.")
			break
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func downloadFile(filePath string, url string) error {
	fmt.Printf("Downloading %s\n", url)

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	done := make(chan int64)
	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return err
	}
	go printDownloadPercent(done, filePath, int64(size))

	// Write the body to file
	written, err := io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	done <- written
	return nil
}
