package integrationtest

import (
	"strings"
	"testing"
)

// TestCommandInterface tests the command line interface
func TestCommandInterface(t *testing.T) {
	helper := NewTestHelper(t)
	helper.BuildBinary(t)

	t.Run("TestHelpCommand", func(t *testing.T) {
		stdout, stderr, _ := helper.RunCommand(t, "--help")
		output := stdout + stderr

		// Verify help contains expected content
		expectedStrings := []string{
			"gh-merge-base-next",
			"<base>",
			"<head>",
			"--walk-to",
			"--repo",
		}

		for _, expected := range expectedStrings {
			if !strings.Contains(output, expected) {
				t.Errorf("Help output should contain '%s', got: %s", expected, output)
			}
		}
	})

	t.Run("TestVersionCommand", func(t *testing.T) {
		stdout, stderr, _ := helper.RunCommand(t, "--version")
		output := stdout + stderr

		// Verify version output
		if !strings.Contains(output, "gh-merge-base-next") {
			t.Errorf("Version output should contain tool name, got: %s", output)
		}
	})
}

// TestErrorHandling tests various error conditions
func TestErrorHandling(t *testing.T) {
	helper := NewTestHelper(t)
	helper.BuildBinary(t)

	testCases := []struct {
		name string
		args []string
		desc string
	}{
		{
			name: "NoArguments",
			args: []string{},
			desc: "Command with no arguments should show usage",
		},
		{
			name: "OneArgument",
			args: []string{"commit1"},
			desc: "Command with only one argument should show error",
		},
		{
			name: "InvalidCommitFormat",
			args: []string{"invalid", "also-invalid"},
			desc: "Command with invalid commit format should show error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Testing: %s", tc.desc)

			// These should all fail, but we're testing that they fail gracefully
			stdout, stderr, err := helper.RunCommand(t, tc.args...)

			// The command should either fail or show usage
			if err == nil && !strings.Contains(stdout+stderr, "Usage:") {
				t.Errorf("Expected command to fail or show usage for case: %s", tc.name)
			}
		})
	}
}
