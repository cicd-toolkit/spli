package splunk_client

import (
	"log"
	"os"
	"regexp"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func extractField(rex string, body string) string {
	re := regexp.MustCompile(rex)

	// Find the first match in the string
	match := re.FindStringSubmatch(body)
	if len(match) < 2 {
		log.Fatalf("No match found for key 'cval'")
	}

	// Extract the numerical value from the matched string
	return match[1]
}
