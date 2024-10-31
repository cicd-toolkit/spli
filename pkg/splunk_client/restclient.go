package splunk_client

import (
	"crypto/tls"
	"fmt"
	"log"

	config "github.com/cicd-toolkit/spli/pkg/config_manager"
	"github.com/go-resty/resty/v2"
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

func (a *API) DoLogin() (*resty.Response, error) {

	newUrl := fmt.Sprintf("%s://%s:%s/en-GB", a.WebProto, a.Host, a.WebPort)
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	resp1, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Get(newUrl + "/account/login")

	cvalStr := extractField(`"cval":(\d+)`, resp1.String())

	respLogin, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"username":          a.Username,
			"password":          a.Password,
			"set_has_logged_in": "false",
			"cval":              cvalStr,
		}).
		Post(newUrl + "/account/login")

	// fmt.Println("Response Info:")
	// fmt.Println("  Error      :", err)
	// fmt.Println("  Status Code:", respLogin.StatusCode())
	// fmt.Println("  Status     :", respLogin.Status())
	// fmt.Println("  Proto      :", respLogin.Proto())
	// fmt.Println("  Time       :", respLogin.Time())
	// fmt.Println("  Body       :\n", respLogin)

	if err != nil {
		return nil, fmt.Errorf("failed executing api : %v", err)
	}

	if respLogin.IsError() {
		return nil, fmt.Errorf("API responded with errorx: %s - %s", respLogin.Status(), respLogin.Body())
	}
	// fmt.Printf("%+v\n", resp)
	return respLogin, nil

}
