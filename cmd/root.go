package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spli",
	Short: "splunk cli",
	Long:  `splunk cli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use --help for more information.")
	},
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
