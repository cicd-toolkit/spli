package cmd

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

// promptInput prompts the user for input and returns the entered value
func promptInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", prompt)
	input, _ := reader.ReadString('\n')
	return input[:len(input)-1]
}

// promptPassword prompts the user for a password and returns the entered value (without echoing input)
func promptPassword(prompt string) string {
	fmt.Printf("%s: ", prompt)
	bytePassword, _ := terminal.ReadPassword(0)
	fmt.Println()
	return string(bytePassword)
}

// Helper function to parse the form data string into a map[string]interface{}
func parseFormData(data string) (map[string]interface{}, error) {
	formData := make(map[string]interface{})

	// Split the data string by "&" to get key-value pairs
	pairs := strings.Split(data, "&")

	for _, pair := range pairs {
		// Split each pair by "=" to separate key and value
		kv := strings.SplitN(pair, "=", 2)

		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid form data: %s", pair)
		}

		key, value := kv[0], kv[1]
		formData[key] = value
	}

	return formData, nil
}

// Convert a map[string]interface{} to url.Values
func mapToURLValues(data map[string]interface{}) url.Values {
	values := url.Values{}

	for key, value := range data {
		values.Set(key, fmt.Sprintf("%v", value))
	}

	return values
}
