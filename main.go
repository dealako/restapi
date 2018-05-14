package main

import (
	"runtime"

	log "github.com/sirupsen/logrus"

	"github.com/dealako/restapi/cmd"
)

// Build and version variables defined and set during the build process
var (
	// Name the application name
	name string
	// Version the application version
	version string
	// Build/Commit the application build number
	commit string
	// Build date
	buildDate string
)

func main() {
	// Show the version and build info
	log.Infof("Name                  : %s", name)
	log.Infof("Version               : %s", version)
	log.Infof("Git commit hash       : %s", commit)
	log.Infof("Build date            : %s", buildDate)
	log.Infof("Golang OS             : %s", runtime.GOOS)
	log.Infof("Golang Arch           : %s", runtime.GOARCH)

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
