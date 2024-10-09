package cmd

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/url"

	restclient "github.com/cicd-toolkit/spli/pkg/splunk_client"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "bump",
	Long:  `bump`,
	RunE: func(cmd *cobra.Command, args []string) error {
		api, err := restclient.SplunkClient()
		if err != nil {
			return fmt.Errorf("failed api : %v", err)
		}

		u, err := url.Parse(api.URL)
		if err != nil {
			panic(err)
		}
		host, _, _ := net.SplitHostPort(u.Host)

		newUrl := fmt.Sprintf("http://%s:8000/en-GB", host)
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

		respBump, err := client.R().
			EnableTrace().
			SetCookies(respLogin.Cookies()).
			Get(newUrl + "/_bump")
		splunkFormKey := extractField(`name="splunk_form_key" value="([^"]+)"`, respBump.String())

		resp, err := client.R().
			EnableTrace().
			SetCookies(respLogin.Cookies()).
			SetFormData(map[string]string{
				"splunk_form_key": splunkFormKey,
			}).
			Post(newUrl + "/_bump")
		fmt.Println("  Status     :", resp.Status())
		ver := extractField(`Current version:\s*(\d+)`, resp.String())
		fmt.Println("  Bumped version  : ", ver)
		return nil
	},
}

func init() {
	// Adding the version command to the root command
	rootCmd.AddCommand(bumpCmd)
}
