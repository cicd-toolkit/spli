package cmd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ServerRoleCommand(t *testing.T) {
	out, err := executeCommand(rootCmd, "server_roles")
	assert.NoError(t, err)
	expectedOutput := `["indexer", "license_master", "license_manager", "kv_store"]`
	cleanOut := strings.TrimSuffix(string(out), "\n")
	assert.Equal(t, expectedOutput, cleanOut)
}
