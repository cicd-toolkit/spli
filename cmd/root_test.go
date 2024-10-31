package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/blang/semver"
	"github.com/spf13/cobra"
)

func IsVersionInRange(minVersion, maxVersion string) (bool, error) {
	currentVersion := os.Getenv("SPLUNK_VERSION")
	if currentVersion == "" {
		return false, fmt.Errorf("SPLUNK_VERSION is not set")
	}

	current, err := semver.ParseTolerant(currentVersion)
	if err != nil {
		return false, fmt.Errorf("failed to parse SPLUNK_VERSION: %v", err)
	}

	min, err := semver.ParseTolerant(minVersion)
	if err != nil {
		return false, fmt.Errorf("failed to parse minimum version: %v", err)
	}

	max, err := semver.ParseTolerant(maxVersion)
	if err != nil {
		return false, fmt.Errorf("failed to parse maximum version: %v", err)
	}

	// Check if current version is within the range (minVersion, maxVersion)
	return current.GT(min) && current.LT(max), nil
}

// IsVersionGTE returns true if SPLUNK_VERSION is greater than or equal to minVersion.
func IsVersionGTE(minVersion string) (bool, error) {
	currentVersion := os.Getenv("SPLUNK_VERSION")
	if currentVersion == "" {
		return false, fmt.Errorf("SPLUNK_VERSION is not set")
	}

	current, err := semver.ParseTolerant(currentVersion)
	if err != nil {
		return false, fmt.Errorf("failed to parse SPLUNK_VERSION: %v", err)
	}

	min, err := semver.ParseTolerant(minVersion)
	if err != nil {
		return false, fmt.Errorf("failed to parse minimum version: %v", err)
	}

	// Check if current version is greater than or equal to the minVersion
	return current.GTE(min), nil
}

func IsVersionLTE(maxVersion string) (bool, error) {
	currentVersion := os.Getenv("SPLUNK_VERSION")
	if currentVersion == "" {
		return false, fmt.Errorf("SPLUNK_VERSION is not set")
	}

	current, err := semver.ParseTolerant(currentVersion)
	if err != nil {
		return false, fmt.Errorf("failed to parse SPLUNK_VERSION: %v", err)
	}

	max, err := semver.ParseTolerant(maxVersion)
	if err != nil {
		return false, fmt.Errorf("failed to parse maximum version: %v", err)
	}

	// Check if current version is less than or equal to the maxVersion
	return current.LTE(max), nil
}

func executeCommand(root *cobra.Command, arg string) (output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(strings.Fields(arg))

	err = root.Execute()
	return strings.TrimSpace(buf.String()), err
}
