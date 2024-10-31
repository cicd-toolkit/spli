package cmd

import (
	"crypto/tls"
	"fmt"
	"strings"

	restclient "github.com/cicd-toolkit/spli/pkg/splunk_client"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var messsagesCmd = &cobra.Command{
	Use:   "messsage",
	Short: "messsage",
	Long: strings.Replace(`
this command is passing the system messages and print over cli
example:

spli msg
'''
[
  {
    "NoAllowedDomainsList": "Security risk warning: Found an empty value for allowedDomainList in the alert_actions.conf configuration file. If you do not configure this setting, then users can send email alerts with search results to any domain. You can add values for allowedDomainList either in the alert_actions.conf file or in Server Settings > Email Settings > Email Domains in Splunk Web.",
    "capabilities": ["admin_all_objects"],
    "eai:acl": null,
    "help": "",
    "message": "Security risk warning: Found an empty value for allowedDomainList in the alert_actions.conf configuration file. If you do not configure this setting, then users can send email alerts with search results to any domain. You can add values for allowedDomainList either in the alert_actions.conf file or in Server Settings > Email Settings > Email Domains in Splunk Web.",
    "message_alternate": "",
    "server": "36c2770ea6dc",
    "severity": "warn",
    "timeCreated_epochSecs": 1730382326,
    "timeCreated_iso": "2024-10-31T13:45:26+00:00"
  }
]
'''

use '--jsonpath' to select any output
'''
spli msg --jsonpath "entry.#.content.severity|@pretty"
["warn"]
'''
`, "'", "`", -1),
	Aliases: []string{"msg"},
	RunE: func(cmd *cobra.Command, args []string) error {

		api, err := restclient.SplunkClient()
		if err != nil {
			return fmt.Errorf("failed api : %v", err)
		}
		newUrl := fmt.Sprintf("%s://%s:%s/en-GB", api.WebProto, api.Host, api.WebPort)

		respLogin, err := api.DoLogin()
		if err != nil {
			return fmt.Errorf("failed executing api : %v", err)
		}
		client := resty.New()
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		// http://localhost:8000/en-GB/splunkd/__raw/services/messages?output_mode=json&sort_key=timeCreated_epochSecs&sort_dir=desc&count=1000&_=1730392300714
		resp, err := client.R().
			EnableTrace().
			SetCookies(respLogin.Cookies()).
			SetQueryParams(map[string]string{
				"count":       "1000",
				"sort_dir":    "desc",
				"sort_key":    "timeCreated_epochSecs",
				"output_mode": "json",
			}).
			Get(newUrl + "/splunkd/__raw/services/messages")

		// fmt.Println("Response Info:")
		// fmt.Println("  Error      :", err)
		// fmt.Println("  Status Code:", resp.StatusCode())
		// fmt.Println("  Status     :", resp.Status())
		// fmt.Println("  Proto      :", resp.Proto())
		// fmt.Println("  Time       :", resp.Time())
		// fmt.Println("  Body       :\n", resp)
		// fmt.Println()

		cmd.Print(gjson.Get(resp.String(), messsagesJsonPath).String())

		return nil
	},
}

var messsagesJsonPath string

func init() {
	rootCmd.AddCommand(messsagesCmd)
	messsagesCmd.Flags().StringVarP(&messsagesJsonPath, "jsonpath", "", "entry.#.content|@pretty", "jsonpath using gjson")
}
