package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type LoginData struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Variables to hold the flag values
var host, username, password string

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "setup with Host, username, and password",
	Long:  `This command allows you to setup URL, username, and password.`,
	Run: func(cmd *cobra.Command, args []string) {
		// If flags are not provided, prompt the user
		if host == "" {
			host = promptInput("Host")
		}
		if username == "" {
			username = promptInput("Username")
		}
		if password == "" {
			password = promptPassword("Password")
		}

		fmt.Printf("Logging in with Host: %s, Username: %s\n", host, username)
		// Save the login info to file
		saveLoginData(LoginData{Host: host, Username: username, Password: password})
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	setupCmd.Flags().StringVarP(&host, "host", "u", "", "splunk host")
	setupCmd.Flags().StringVarP(&username, "username", "n", "", "Username for login")
	setupCmd.Flags().StringVarP(&password, "password", "p", "", "Password for login")
}

func saveLoginData(data LoginData) {
	file, err := os.Create(".spli")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
	} else {
		fmt.Println("Login data saved to .spli")
	}
}
