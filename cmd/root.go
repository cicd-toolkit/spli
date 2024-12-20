package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spli",
	Short: "splunk cli",
	Long:  `splunk cli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Splunk CLI")
	},
}

func SetVersionInfo(version, commit string) {
	re := regexp.MustCompile(`\d+\.\d+\.\d+`)
	// Find the first match of the pattern in the version string
	semver := re.FindString(version)
	rootCmd.Version = fmt.Sprintf("%s (Git SHA %s)", semver, commit)
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
