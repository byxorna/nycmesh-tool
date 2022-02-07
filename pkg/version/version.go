package version

import (
	"fmt"
)

var (
	GitCommit string = "unknown"
	GitBranch string = "unknown"
	Release   string
	BuildDate string
)

func UserAgent() string {
	return fmt.Sprintf("nycmesh-tool/%s-%s", Release, GitCommit)
}
