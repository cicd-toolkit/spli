package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const UA = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36"

var sbCmd = &cobra.Command{
	Use:     "splunkbase",
	Short:   "splunkbase",
	Aliases: []string{"sb"},
	Long:    `splunkbase group command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("please provide a sub-command")
	},
}

func init() {
	rootCmd.AddCommand(sbCmd)
}
