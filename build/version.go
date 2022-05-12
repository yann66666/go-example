// Author: yann
// Date: 2021/12/22
// Desc: build

package build

var Commit string
var BuildType string
var V string

const (
	BuildDevelop    = "develop"
	BuildPreRelease = "pr"
	BuildRelease    = "release"
)

func toString(buildType string) string {
	switch buildType {
	case BuildDevelop:
		return "+develop"
	case BuildRelease:
		return "+release"
	case BuildPreRelease:
		return "+pre-release"
	default:
		return "+huh?"
	}
}

func Version() string {
	return V + toString(BuildType) + "." + Commit
}
