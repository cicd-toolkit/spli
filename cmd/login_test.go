package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// TestLoginCommand tests the login command
func TestLoginCommand(t *testing.T) {
	spliFile := ".spli"

	defer func() {
		os.Remove(spliFile)
	}()
	// Create a new root command for testing
	rootCmd := &cobra.Command{Use: "login"}
	rootCmd.AddCommand(loginCmd)

	// Capture the output
	var stdoutBuf bytes.Buffer
	rootCmd.SetOut(io.Writer(&stdoutBuf))

	// Set arguments for the command (simulate CLI input)
	rootCmd.SetArgs([]string{"login", "--url", "http://test.com", "--username", "testuser", "--password", "testpass"})

	// Run the command
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("Failed to execute command: %v", err)
	}

	// Verify the .spli file is created
	if _, err := os.Stat(spliFile); os.IsNotExist(err) {
		t.Fatalf("Expected .spli file to be created")
	}

	// Read and verify the contents of the file
	fileContent, err := ioutil.ReadFile(spliFile)
	if err != nil {
		t.Fatalf("Failed to read .spli file: %v", err)
	}

	var loginData map[string]interface{}
	if err := json.Unmarshal(fileContent, &loginData); err != nil {
		t.Fatalf("Failed to unmarshal JSON from .spli file: %v", err)
	}

	assert.NoError(t, err, "Should read credentials without error")
	assert.Equal(t, "http://test.com", loginData["url"].(string), "URL should match expected value")
	assert.Equal(t, "testuser", loginData["username"].(string), "Username should match expected value")
	assert.Equal(t, "testpass", loginData["password"].(string), "Password should match expected value")

}
