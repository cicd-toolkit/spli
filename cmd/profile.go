package cmd

import (
	"log"

	config "github.com/cicd-toolkit/spli/pkg/config_manager"
	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "configure profile",
	Long:  `This command allows you to set the active profile.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		cfg, err := config.NewConfig()
		if err != nil {
			log.Fatalf("Error initializing config: %v", err)
		}

		if !ContainsString(cfg.Sections(), args[0]) {
			log.Fatalf("profile '%s' not found", args[0])
		}

		err = cfg.SetString("", "active_profile", args[0])
		if err != nil {
			log.Fatalf("Error setting string value: %v", err)
		}
		cmd.Printf("Active profile : %s ", args[0])

	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
