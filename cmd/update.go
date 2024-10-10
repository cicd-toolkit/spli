package cmd

import (
	"crypto/tls"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update",
	Long:  `update`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := resty.New()
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		resp, err := client.R().
			Get("https://api.github.com/repos/cicd-toolkit/spli/releases/latest")

		if err != nil {
			return fmt.Errorf("failed executing api : %v", err)
		}

		lastVersion := gjson.Get(resp.String(), "name|@pretty")
		if lastVersion.String() == rootCmd.Version {
			fmt.Print("Nothing to update")
		} else {
			fmt.Println("New version found : " + lastVersion.String())
			fmt.Println("update your cli with\n")
			fmt.Println(" curl https://raw.githubusercontent.com/cicd-toolkit/spli/master/scripts/install | bash")
		}

		return nil

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
