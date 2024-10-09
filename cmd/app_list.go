package cmd

import (
	"crypto/tls"
	"fmt"

	restclient "github.com/cicd-toolkit/spli/pkg/splunk_client"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var appListCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long:  `list`,
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
			Get(api.URL + "/services/apps/local")

		fmt.Print(gjson.Get(resp.String(), "entry.#.name|@pretty"))
		return nil
	},
}

func init() {
	appCmd.AddCommand(appListCmd)
}
