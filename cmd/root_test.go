package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/blang/semver"
	"github.com/spf13/cobra"
)

// getParsedSplunkVersion fetches and parses SPLUNK_VERSION environment variable.
func getParsedSplunkVersion() (semver.Version, string, error) {
	currentVersion := os.Getenv("SPLUNK_VERSION")
	if currentVersion == "" {
		return semver.Version{}, "", fmt.Errorf("SPLUNK_VERSION is not set")
	}

	parsedVersion, err := semver.ParseTolerant(currentVersion)
	if err != nil {
		return semver.Version{}, currentVersion, fmt.Errorf("failed to parse SPLUNK_VERSION: %v", err)
	}

	return parsedVersion, currentVersion, nil
}

// IsVersionEQ returns true if SPLUNK_VERSION is exactly equal to targetVersion.
func IsVersionEQ(targetVersion string) (bool, string, error) {
	current, currentVersion, err := getParsedSplunkVersion()
	if err != nil {
		return false, currentVersion, err
	}

	target, err := semver.ParseTolerant(targetVersion)
	if err != nil {
		return false, currentVersion, fmt.Errorf("failed to parse target version: %v", err)
	}

	return current.Equals(target), currentVersion, nil
}

// IsVersionGTE returns true if SPLUNK_VERSION is greater than or equal to minVersion.
func IsVersionGTE(minVersion string) (bool, string, error) {
	current, currentVersion, err := getParsedSplunkVersion()
	if err != nil {
		return false, currentVersion, err
	}

	min, err := semver.ParseTolerant(minVersion)
	if err != nil {
		return false, currentVersion, fmt.Errorf("failed to parse minimum version: %v", err)
	}

	return current.GTE(min), currentVersion, nil
}

// IsVersionLTE returns true if SPLUNK_VERSION is less than or equal to maxVersion.
func IsVersionLTE(maxVersion string) (bool, string, error) {
	current, currentVersion, err := getParsedSplunkVersion()
	if err != nil {
		return false, currentVersion, err
	}

	max, err := semver.ParseTolerant(maxVersion)
	if err != nil {
		return false, currentVersion, fmt.Errorf("failed to parse maximum version: %v", err)
	}

	return current.LTE(max), currentVersion, nil
}

// IsVersionInRange returns true if SPLUNK_VERSION is greater than minVersion and less than maxVersion.
func IsVersionInRange(minVersion, maxVersion string) (bool, string, error) {
	current, currentVersion, err := getParsedSplunkVersion()
	if err != nil {
		return false, currentVersion, err
	}

	min, err := semver.ParseTolerant(minVersion)
	if err != nil {
		return false, currentVersion, fmt.Errorf("failed to parse minimum version: %v", err)
	}

	max, err := semver.ParseTolerant(maxVersion)
	if err != nil {
		return false, currentVersion, fmt.Errorf("failed to parse maximum version: %v", err)
	}

	return current.GT(min) && current.LT(max), currentVersion, nil
}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	err = root.Execute()
	return strings.TrimSpace(buf.String()), err
}
