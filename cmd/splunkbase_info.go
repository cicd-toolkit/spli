package cmd

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var sbInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "info",
	Args:  cobra.ExactArgs(1),
	Long:  `This command allows you get info from for an App`,
	RunE: func(cmd *cobra.Command, args []string) error {

		client := resty.New()
		respLogin, err := client.R().
			SetHeader("user-agent", UA).
			Get("https://api.splunkbase.splunk.com/limelight-login/")

		fmt.Println("Response Info:")
		fmt.Println("  Error      :", err)
		fmt.Println("  Status Code:", respLogin.StatusCode())

		return nil

	},
}

func init() {
	sbCmd.AddCommand(sbInfoCmd)
}
