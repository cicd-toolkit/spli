package cmd

import (
	"fmt"
	"log"

	config "github.com/cicd-toolkit/spli/pkg/config_manager"
	"github.com/spf13/cobra"
)

var sbSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "setup",
	Long:  `This command allows you to setup splunkbase  username and password.`,
	Run: func(cmd *cobra.Command, args []string) {
		if username == "" {
			username = promptInput("Username")
		}
		if password == "" {
			password = promptPassword("Password")
		}

		fmt.Printf("Logging in splunkbase with Username: %s\n", username)
		cfg, err := config.NewConfig()
		if err != nil {
			log.Fatalf("Error initializing config: %v", err)
		}

		err = cfg.SetValue("splunkbase_username", username)
		if err != nil {
			log.Fatalf("Error setting string value: %v", err)
		}
		err = cfg.SetValue("splunkbase_password", password)
		if err != nil {
			log.Fatalf("Error setting string value: %v", err)
		}

	},
}

func init() {
	sbCmd.AddCommand(sbSetupCmd)
	sbSetupCmd.Flags().StringVarP(&username, "username", "n", "", "Username for login")
	sbSetupCmd.Flags().StringVarP(&password, "password", "p", "", "Password for login")
}
