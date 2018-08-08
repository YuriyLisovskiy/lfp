// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package updater

import (
	"context"
	"runtime"
	"strings"

	"github.com/google/go-github/github"
)

func StartUpdate(version string) error {
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
	err = DownloadFile(path, downloadUrl)
	if err != nil {
		return err
	}
	startDaemon(path)
	return nil
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

		// Get latest release
		url, err = chooseArchive(releases[0].Assets)
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
