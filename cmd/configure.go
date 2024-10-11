package cmd

import (
	"log"

	config "github.com/cicd-toolkit/spli/pkg/config_manager"
	"github.com/spf13/cobra"
)

type LoginData struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Variables to hold the flag values
var host, username, password, profileName, adminPort, webPort, webProto string

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "configure with Host, username, and password",
	Long:  `This command allows you to configure Host, username, and password.`,
	Run: func(cmd *cobra.Command, args []string) {

		host = GetInputWithDefault("SPLUNK_HOST", "Host", "localhost")
		username = GetInputWithDefault("SPLUNK_USERNAME", "Username", "admin")
		password = GetInputWithDefault("SPLUNK_PASSWORD", "Password", "admin")
		adminPort = GetInputWithDefault("SPLUNK_ADMINPORT", "admin port", "8089")
		webPort = GetInputWithDefault("SPLUNK_PORT", "web port", "8000")
		webProto = GetInputWithDefault("SPLUNK_PROTO", "protocol", "http")

		cfg, err := config.NewConfig()

		if err != nil {
			log.Fatalf("Error initializing config: %v", err)
		}
		if profileName == "" {
			profileName = cfg.GetString("", "active_profile")
		}
		if profileName == "" {
			profileName = "default"
		}

		_ = cfg.SetString(profileName, "host", host)
		_ = cfg.SetString(profileName, "username", username)
		_ = cfg.SetString(profileName, "password", password)
		_ = cfg.SetString(profileName, "admin_port", adminPort)
		_ = cfg.SetString(profileName, "web_port", webPort)
		_ = cfg.SetString(profileName, "protocol", webProto)

		cmd.Print("Done")

	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.Flags().StringVarP(&host, "host", "", "", "splunk host")
	configureCmd.Flags().StringVarP(&username, "username", "", "", "Username for login")
	configureCmd.Flags().StringVarP(&password, "password", "", "", "Password for login")
	configureCmd.Flags().StringVarP(&adminPort, "adminport", "", "", "admin Port")
	configureCmd.Flags().StringVarP(&webPort, "webport", "", "", "web Port")
	configureCmd.Flags().StringVarP(&webProto, "webproto", "", "", "web Protocol")
	configureCmd.Flags().StringVarP(&profileName, "profile", "p", "", "profile name")
}
