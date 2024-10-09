package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

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
