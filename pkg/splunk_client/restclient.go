package splunk_client

import (
	"log"

	config "github.com/cicd-toolkit/spli/pkg/config_manager"
)

// API struct to hold the base URL of the REST API and bearer token
type API struct {
	Host      string
	AdminPort string
	WebPort   string
	WebProto  string
	Username  string
	Password  string
}

var profileName string

// SplunkClient creates a new API instance with authentication
func SplunkClient() (*API, error) {

	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}
	profileName = cfg.GetString("", "active_profile")
	if profileName == "" {
		profileName = cfg.GetString("", "active_profile")
	}
	if profileName == "" {
		profileName = "default"
	}
	return &API{
		Host:      cfg.GetString(profileName, "host"),
		AdminPort: cfg.GetString(profileName, "admin_port"),
		WebPort:   cfg.GetString(profileName, "web_port"),
		WebProto:  cfg.GetString(profileName, "protocol"),
		Username:  cfg.GetString(profileName, "username"),
		Password:  cfg.GetString(profileName, "password"),
	}, nil
}
