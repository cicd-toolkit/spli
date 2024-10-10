package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// LoginData holds the login information
type LoginData struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Variables to hold the flag values
var loginURL, username, password string

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "setup with URL, username, and password",
	Long:  `This command allows you to setup URL, username, and password.`,
	Run: func(cmd *cobra.Command, args []string) {
		// If flags are not provided, prompt the user
		if loginURL == "" {
			loginURL = promptInput("URL")
		}
		if username == "" {
			username = promptInput("Username")
		}
		if password == "" {
			password = promptPassword("Password")
		}

		fmt.Printf("Logging in with URL: %s, Username: %s\n", loginURL, username)
		// Save the login info to file
		saveLoginData(LoginData{URL: loginURL, Username: username, Password: password})
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Add flags for the login command
	setupCmd.Flags().StringVarP(&loginURL, "url", "u", "", "URL for login")
	setupCmd.Flags().StringVarP(&username, "username", "n", "", "Username for login")
	setupCmd.Flags().StringVarP(&password, "password", "p", "", "Password for login")
}

// saveLoginData saves the login data to a .spli file in JSON format
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
