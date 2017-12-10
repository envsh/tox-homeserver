package gofia

import (
	"fmt"
	"log"
	"strings"
)

// build info
var GitCommit, GitBranch, GitState, GitSummary, BuildDate, Version string

func printBuildInfo(full bool) { log.Println(getBuildInfo(full)) }
func getBuildInfo(full bool) string {
	s := btversion
	s = strings.Replace(s, "-X ", "", -1)
	s = strings.Replace(s, "gofia.", "", -1)
	s = strings.Replace(s, "GOVVV-", "", -1)
	s = strings.Replace(s, "=", ":", -1)
	return fmt.Sprintf("govvv: %s", s)
}
