package main

import (
	"github.com/cicd-toolkit/spli/cmd"
)

var (
	version = "0.0.0"
	commit  = "localdev"
)

func main() {
	cmd.SetVersionInfo(version, commit)
	cmd.Execute()
}
