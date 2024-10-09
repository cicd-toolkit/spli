package splunk_client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// API struct to hold the base URL of the REST API and bearer token
type API struct {
	URL      string
	Username string
	Password string
}

const localCreds = ".spli"

// SplunkClient creates a new API instance with authentication
func SplunkClient() (*API, error) {

	if _, err := os.Stat(localCreds); err == nil {
		// File exists, read configuration from file
		creds := make(map[string]interface{})
		file, err := os.Open(localCreds)
		defer file.Close()

		byteValue, err := ioutil.ReadAll(file)

		err = json.Unmarshal(byteValue, &creds)
		if err != nil {
			return nil, errors.New("errro reading json")
		}

		return &API{
			URL:      creds["url"].(string),
			Username: creds["username"].(string),
			Password: creds["password"].(string),
		}, nil
	}

	url := getenv("SPLUNK_ENDPOINT", "http://localhost:8000/")
	if url == "" {
		return nil, errors.New("SPLUNK_ENDPOINT environment variable is not set")
	}

	username := os.Getenv("SPLUNK_USERNAME")
	if username == "" {
		return nil, errors.New("SPLUNK_USERNAME environment variable is not set")
	}
	password := os.Getenv("SPLUNK_PASSWORD")
	if password == "" {
		return nil, errors.New("SPLUNK_PASSWORD environment variable is not set")
	}

	return &API{
		URL:      url,
		Username: username,
		Password: password,
	}, nil
}
