// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package updater

import (
	"os"
	"fmt"
	"time"
	"bytes"
	"errors"
	"regexp"
	"context"
	"runtime"
	"strings"
	"os/exec"

	"github.com/mholt/archiver"
	"github.com/google/go-github/github"
)

func startUpdate(version string) error {
	start := time.Now()
	downloadUrl, err := getReleaseUrl(version)
	if err != nil {
		return err
	}
	split := strings.Split(downloadUrl, "/")
	cwd, err := getCWD()
	if err != nil {
		return err
	}
	path := cwd + "/" + split[len(split)-1]

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err = downloadFile(path, downloadUrl)
			if err != nil {
				return err
			}
		}
	}
	lfpLoc, err := os.Executable()
	if err != nil {
		return errors.New(updater + ": error: unable to get executable location")
	}
	split = strings.Split(lfpLoc, "/")
	return installUpdate(path, strings.Join(split[:len(split)-1], "/"), start)
}

func getReleases() ([]*github.RepositoryRelease, error) {
	client := github.NewClient(nil)
	releases, resp, err := client.Repositories.ListReleases(context.Background(), author, repo, nil)
	if err != nil {
		return nil, err
	}
	if resp.Status == "200 OK" {
		return releases, nil
	}
	return nil, ErrCantFetchReleases
}

// getVersionUrl retrieves browser download url for the user-defined release
func getReleaseUrl(version string) (string, error) {
	version = strings.ToLower(version)

	// Fetch all available releases
	releases, err := getReleases()
	if err != nil {
		return "", err
	}
	if len(releases) < 1 {
		return "", ErrReleasesNotFound
	}
	url := ""
	switch version {
	case "latest":
		latest := releases[0]

		cmd := exec.Command("lfp", "--version")
		var out, errOut bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &errOut
		cmd.Run()
		if errOut.String() != "" {
			fmt.Println(errOut.String())
		}

		currentVersion := ""
		rx, _ := regexp.Compile(`\d.\d.\d`)

		split := strings.Split(out.String(), "\n")

		if rx.MatchString(split[0]) {
			currentVersion = rx.FindString(split[0])
		} else {
			return "", ErrCantRetrieveLfpVer
		}

		// Check if the newest version is already installed
		if *latest.TagName == currentVersion {
			return "", ErrTheLatestAlreadyInstalled
		}

		// Check if version is not greater than the latest release
		if *latest.TagName < currentVersion {
			return "", ErrLFPIsBroken
		}

		// Get latest release
		url, err = chooseArchive(latest.Assets)
	default:
		var targetRelease *github.RepositoryRelease = nil
		version = strings.TrimPrefix(version, "v")

		// Try to get user-defined release
		for _, release := range releases {
			if strings.TrimPrefix(strings.ToLower(*release.TagName), "v") == version {
				targetRelease = release
			}
		}

		// If release was not found, return an error
		if targetRelease == nil {
			return "", ErrVersionNotFound
		}
		url, err = chooseArchive(targetRelease.Assets)
	}
	return url, err
}

func chooseArchive(assets []github.ReleaseAsset) (string, error) {

	// Get target operating system and architecture for retrieving a suitable url
	targetOs := runtime.GOOS
	targetArch := runtime.GOARCH

	// Try to get suitable url
	for _, asset := range assets {
		url := *asset.BrowserDownloadURL
		if strings.Contains(url, targetOs) && strings.Contains(url, targetArch) {
			return url, nil
		}
	}
	return "", ErrNoReleaseForTargetOs
}

// installUpdate installs downloaded update using daemon process
func installUpdate(path, exec string, start time.Time) error {
	fmt.Println("LFP Updater: Updating LFP tool...")
	targetOs := runtime.GOOS
	switch targetOs {
	case "windows":
		err := archiver.Zip.Open(path, exec)
		if err != nil {
			return errors.New(fmt.Sprintf(updater + ": error: opening downloaded archive with executable, %s", err))
		}
	case "linux":
		err := archiver.TarGz.Open(path, exec)
		if err != nil {
			return errors.New(fmt.Sprintf(updater + ": error: opening downloaded archive with executable, %s", err))
		}
	}
	elapsed := time.Since(start)
	fmt.Printf(updater + ": LFP tool has been updated successfully, time elapsed: %d sec\n", int64(elapsed / time.Second))
	return nil
}
