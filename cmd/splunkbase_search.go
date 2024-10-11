package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

// curl 'https://api.splunkbase.splunk.com/v3/suggestions?q=aws' \
//   -H 'accept: application/json, text/plain, */*' \
//   -H 'accept-language: en-GB,en-US;q=0.9,en;q=0.8' \
//   -H 'cache-control: no-cache' \
//   -H 'origin: https://splunkbase.splunk.com' \
//   -H 'pragma: no-cache' \
//   -H 'priority: u=1, i' \
//   -H 'referer: https://splunkbase.splunk.com/' \
//   -H 'sec-ch-ua: "Google Chrome";v="129", "Not=A?Brand";v="8", "Chromium";v="129"' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'sec-ch-ua-platform: "macOS"' \
//   -H 'sec-fetch-dest: empty' \
//   -H 'sec-fetch-mode: cors' \
//   -H 'sec-fetch-site: same-site' \
//   -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36'

var sbSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "search",
	Args:  cobra.ExactArgs(1),
	Long:  `This command allows you to setup splunkbase  username and password.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		client := resty.New()
		resp, err := client.R().
			SetQueryParams(map[string]string{
				"q": args[0],
			}).
			SetHeader("user-agent", UA).
			Get("https://api.splunkbase.splunk.com/v3/suggestions")
		if err != nil {
			return fmt.Errorf("failed executing api : %v", err)
		}
		var formattedData []map[string]interface{}
		result := gjson.Get(resp.String(), "@pretty")
		result.ForEach(func(key, value gjson.Result) bool {
			person := map[string]interface{}{
				"name": value.Get("text").String(),
				"id":   value.Get("id").Int(),
			}
			formattedData = append(formattedData, person)
			return true // Keep iterating
		})

		formattedJSON, err := json.MarshalIndent(formattedData, "", "  ")
		if err != nil {
			log.Fatalf("Error marshalling JSON: %v", err)
		}

		// Print the formatted JSON string
		fmt.Println(string(formattedJSON))
		return nil

	},
}

func init() {
	sbCmd.AddCommand(sbSearchCmd)
}
