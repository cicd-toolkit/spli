package cmd

import (
	"crypto/tls"
	"fmt"

	restclient "github.com/cicd-toolkit/spli/pkg/splunk_client"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var appInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "info",
	Long:  `info`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
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
			Get(fmt.Sprintf("https://%s:%s/services/apps/local/%s", api.Host, api.AdminPort, args[0]))
		if err != nil {
			return fmt.Errorf("failed executing api : %v", err)
		}
		fmt.Print(gjson.Get(resp.String(), "entry|@pretty"))
		return nil
	},
}

func init() {
	appCmd.AddCommand(appInfoCmd)
}
