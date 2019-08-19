package version

import "fmt"

var (
	// Version is the queue's version will set on build time
	Version string
	// GitVersion is the latest commit hash while building
	GitVersion string
	// BuildTime is the time at which build is executed
	BuildTime string
)

// PrintVersion will print the version info to the stdout
func PrintVersion() {
	fmt.Printf("version\t\t\t%s\ngit version\t\t%s\nbuild time\t\t%s", Version, GitVersion, BuildTime)
}
