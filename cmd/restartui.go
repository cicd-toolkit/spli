package cmd

import (
	"crypto/tls"
	"fmt"

	restclient "github.com/cicd-toolkit/spli/pkg/splunk_client"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var restartuiCmd = &cobra.Command{
	Use:   "restartui",
	Short: "restartui",
	Long:  `restartui`,
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
			Post(fmt.Sprintf("https://%s:%s/services/server/control/restart_webui", api.Host, api.AdminPort))
		if err != nil {
			return fmt.Errorf("failed executing api : %v", err)
		}
		fmt.Print(gjson.Get(resp.String(), "@pretty"))
		return nil
	},
}

func init() {
	// Adding the version command to the root command
	rootCmd.AddCommand(restartuiCmd)
}
