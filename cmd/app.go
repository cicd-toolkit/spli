package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "app",
	Long:  `app`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("please provide a sub-command")
	},
}

func init() {
	rootCmd.AddCommand(appCmd)
}
