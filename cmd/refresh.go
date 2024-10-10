package cmd

import (
	"crypto/tls"
	"fmt"

	restclient "github.com/cicd-toolkit/spli/pkg/splunk_client"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "refresh",
	Long:  `refresh`,
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := restclient.SplunkClient()
		if err != nil {
			return fmt.Errorf("failed api : %v", err)
		}
		newUrl := fmt.Sprintf("%s://%s:%s/en-GB", api.WebProto, api.Host, api.WebPort)
		client := resty.New()
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

		resp1, err := client.R().
			EnableTrace().
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			Get(newUrl + "/account/login")
		cvalStr := extractField(`"cval":(\d+)`, resp1.String())

		respLogin, err := client.R().
			EnableTrace().
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			SetFormData(map[string]string{
				"username":          api.Username,
				"password":          api.Password,
				"set_has_logged_in": "false",
				"cval":              cvalStr,
			}).
			Post(newUrl + "/account/login")

		if err != nil {
			return fmt.Errorf("failed executing api : %v", err)
		}

		respRefresh, err := client.R().
			EnableTrace().
			SetCookies(respLogin.Cookies()).
			Get(newUrl + "/debug/refresh")
		splunkFormKey := extractField(`name="splunk_form_key" value="([^"]+)"`, respRefresh.String())

		resp, err := client.R().
			EnableTrace().
			SetCookies(respLogin.Cookies()).
			SetFormData(map[string]string{
				"splunk_form_key": splunkFormKey,
			}).
			Post(newUrl + "/debug/refresh")
		fmt.Println("  Status   :", resp.Status())
		fmt.Println("  Body     :\n", resp)
		return nil
	},
}

func init() {
	// Adding the version command to the root command
	rootCmd.AddCommand(refreshCmd)
}
