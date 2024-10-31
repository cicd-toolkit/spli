package cmd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ServerRoleCommand(t *testing.T) {
	out, err := executeCommand(rootCmd, "server_roles")
	assert.NoError(t, err)
	expectedOutput := `["indexer", "license_master", "kv_store"]`

	isGTE, _, err := IsVersionGTE("9.0.0")
	if err != nil {
		t.Fatalf("Error checking SPLUNK_VERSION: %v", err)
	}

	if isGTE {
		t.Log("Running test for SPLUNK_VERSION >= 9.0.0")
		expectedOutput = `["indexer", "license_master", "license_manager", "kv_store"]`
	}

	cleanOut := strings.TrimSuffix(string(out), "\n")
	assert.Equal(t, expectedOutput, cleanOut)
}
