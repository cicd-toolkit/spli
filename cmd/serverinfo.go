package cmd

import (
	"crypto/tls"
	"fmt"

	restclient "github.com/cicd-toolkit/spli/pkg/splunk_client"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var serverinfoCmd = &cobra.Command{
	Use:   "serverinfo",
	Short: "serverinfo",
	Long:  `serverinfo`,
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
			Get("https://" + api.Host + "/services/server/info")
		if err != nil {
			return fmt.Errorf("failed executing api : %v", err)
		}
		fmt.Print(gjson.Get(resp.String(), "@pretty"))
		return nil

	},
}

func init() {
	// Adding the version command to the root command
	rootCmd.AddCommand(serverinfoCmd)
}
