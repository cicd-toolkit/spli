package main

import (
	"github.com/cicd-toolkit/spli/cmd"
)

var (
	version = "dev"
	commit  = "none"
)

func main() {
	cmd.SetVersionInfo(version, commit)
	cmd.Execute()
}
