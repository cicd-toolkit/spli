package cmd

import (
	"crypto/tls"
	"fmt"

	restclient "github.com/cicd-toolkit/spli/pkg/splunk_client"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var confirmAppDelFlag bool

var appDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "del",
	Long:  `delete`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		if confirmAppDelFlag == false {
			answer := promptInput("Are you sure to delete app : " + args[0] + " [y/N]")
			if answer != "y" {
				return nil
			}
		}

		api, err := restclient.SplunkClient()
		if err != nil {
			return fmt.Errorf("failed api : %v", err)
		}

		client := resty.New()
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		resp, err := client.R().
			SetQueryParam("output_mode", "json").
			SetBasicAuth(api.Username, api.Password).
			SetHeader("Accept", "application/json").
			Delete(fmt.Sprintf("https://%s:%s/services/apps/local/%s", api.Host, api.AdminPort, args[0]))
		if err != nil {
			return fmt.Errorf("failed executing api : %v", err)
		}
		fmt.Print(gjson.Get(resp.String(), "@pretty"))
		return nil
	},
}

func init() {
	appCmd.AddCommand(appDeleteCmd)
	appDeleteCmd.Flags().BoolVar(&confirmAppDelFlag, "yes", false, "confirm deletion")
}
