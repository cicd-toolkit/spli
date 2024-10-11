package cmd

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Long:  `version`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Print(rootCmd.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
