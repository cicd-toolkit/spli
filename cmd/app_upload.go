package cmd

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"os"
	"path/filepath"

	restclient "github.com/cicd-toolkit/spli/pkg/splunk_client"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var appUploadCmd = &cobra.Command{
	Use:     "upload",
	Short:   "upload",
	Long:    `upload`,
	Aliases: []string{"add", "up"},
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		api, err := restclient.SplunkClient()
		if err != nil {
			return fmt.Errorf("failed api : %v", err)
		}
		newUrl := fmt.Sprintf("%s://%s:%s/en-GB", api.WebProto, api.Host, api.WebPort)
		client := resty.New()
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

		respLogin, err := api.DoLogin()
		if err != nil {
			return fmt.Errorf("failed executing api : %v", err)
		}

		respUpPage, err := client.R().
			EnableTrace().
			SetCookies(respLogin.Cookies()).
			Post(newUrl + "/manager/appinstall/_upload")

		splunkFormKey := extractField(`name="splunk_form_key" value="([^"]+)"`, respUpPage.String())
		stateValue := extractField(`name="state" value="([^"]+)"`, respUpPage.String())

		fileBytes, _ := os.ReadFile(args[0])
		fileName := filepath.Base(args[0])
		// // https://github.com/go-resty/resty/issues/109

		resp, err := client.R().
			EnableTrace().
			SetCookies(respLogin.Cookies()).
			SetFileReader("appfile", fileName, bytes.NewReader([]byte(fileBytes))).
			SetFormData(map[string]string{
				"state":           stateValue,
				"splunk_form_key": splunkFormKey,
			}).
			Post(newUrl + "/manager/appinstall/_upload")

		fmt.Println("Response Info:")
		fmt.Println("  Error      :", err)
		fmt.Println("  Status Code:", resp.StatusCode())
		fmt.Println("  Status     :", resp.Status())
		fmt.Println("  Proto      :", resp.Proto())
		fmt.Println("  Time       :", resp.Time())
		// fmt.Println("  Body       :\n", resp)
		// fmt.Println("  Body       :\n", respUp)
		fmt.Println()

		return nil
	},
}

func init() {
	appCmd.AddCommand(appUploadCmd)
}
