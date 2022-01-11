package selfupdate

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/byxorna/nycmesh-tool/pkg/version"
	"github.com/google/go-github/v41/github"
)

var (
	ErrNoNewRelease    = fmt.Errorf("no newer release available")
	ErrConnectionError = fmt.Errorf("network connection error")
)

type Client struct {
	github *github.Client
}

type Release struct {
	Name          string
	Body          string
	TagName       string
	Published     time.Time
	DownloadURL   string
	URL           string
	Commit        string
	DownloadCount int
}

func NewUpdaterGithub() (*Client, error) {
	return &Client{
		github: github.NewClient(nil),
	}, nil
}

func (c *Client) HasNewerRelease(ctx context.Context, releaseAssetName string) (*Release, error) {
	org, repo, err := getGithubOrgRepoFromPkgPath()
	if err != nil {
		return nil, fmt.Errorf("unable to determine github org and repo from package path: %w", err)
	}
	rel, err := c.getLatestRelease(ctx, org, repo, releaseAssetName)
	if err != nil {
		return nil, fmt.Errorf("unable to get latest release: %w", err)
	}

	if version.Release == rel.TagName {
		return nil, ErrNoNewRelease
	}

	if exePath, err := os.Executable(); err == nil {
		// if the exe we were invoked with is still available, stat it and
		// see whether mtime is older than when the release was published
		if exeStat, err := os.Stat(exePath); err == nil {
			if exeStat.ModTime().After(rel.Published) {
				// assuming local binary built newer than release, you
				// are running a newer build than released
				return nil, fmt.Errorf("running executable is release %s, which is newer than available release %s", version.Release, rel.TagName)
			}
		}
	}

	if version.Release == rel.TagName {
		return nil, ErrNoNewRelease
	}
	return rel, nil
}

func (c *Client) getLatestRelease(ctx context.Context, ghorg, ghrepo, releaseAssetName string) (*Release, error) {
	ghrel, _, err := c.github.Repositories.GetLatestRelease(ctx, ghorg, ghrepo)
	if err != nil {
		return nil, fmt.Errorf("unable to resolve latest release: %w", err)
	}
	rel := Release{
		Name:      ghrel.GetName(),
		TagName:   ghrel.GetTagName(),
		Body:      ghrel.GetBody(),
		Published: ghrel.GetPublishedAt().Time,
		URL:       ghrel.GetURL(),
		Commit:    ghrel.GetTargetCommitish(),
	}

	if releaseAssetName != "" {
		// try to find the download URL matching
		for _, ra := range ghrel.Assets {
			if ra.GetName() == releaseAssetName && ra.BrowserDownloadURL != nil {
				rel.DownloadURL = ra.GetBrowserDownloadURL()
				rel.DownloadCount = ra.GetDownloadCount()
				break
			}
		}
	}

	return &rel, nil
}

// look at our import path to figure out how we are named, and if we are using the
// github.com import path, pull out the org and repo
func getGithubOrgRepoFromPkgPath() (string, string, error) {
	pc, _, _, _ := runtime.Caller(2)
	partsx := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(partsx)
	packageName := ""
	funcName := partsx[pl-1]

	if partsx[pl-2][0] == '(' {
		funcName = partsx[pl-2] + "." + funcName
		packageName = strings.Join(partsx[0:pl-2], ".")
	} else {
		packageName = strings.Join(partsx[0:pl-1], ".")
	}

	parts := strings.Split(packageName, "/")
	if len(parts) <= 3 || parts[0] != "github.com" || parts[2] != "nycmesh-tool" {
		return "", "", fmt.Errorf("unable to determine github org and repo from pkg import - version check will be disabled %+v", parts)
	}
	return parts[1], parts[2], nil
}
