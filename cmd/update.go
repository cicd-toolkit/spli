package cmd

import (
	"crypto/tls"
	"fmt"
	"regexp"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update",
	Long:  `check if there is a new version of the cli`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := resty.New()
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		resp, err := client.R().
			Get("https://api.github.com/repos/cicd-toolkit/spli/releases/latest")

		if err != nil {
			return fmt.Errorf("failed executing api : %v", err)
		}
		re := regexp.MustCompile(`\d+\.\d+\.\d+`)
		// Find the first match of the pattern in the version string
		semver := re.FindString(rootCmd.Version)

		lastVersion := gjson.Get(resp.String(), "name|@pretty")
		if lastVersion.String() == semver {
			fmt.Print("Nothing to update")
		} else {
			fmt.Println("Your version      : " + semver)
			fmt.Println("New version found : " + lastVersion.String())
			fmt.Println("curl -sSL https://raw.githubusercontent.com/cicd-toolkit/spli/master/scripts/install | bash")
		}

		return nil

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
