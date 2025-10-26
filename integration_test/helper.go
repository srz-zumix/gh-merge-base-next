package integrationtest

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// TestHelper provides utility functions for integration tests
type TestHelper struct {
	rootDir    string
	binaryPath string
	repoFlag   string // GitHub repository to test against
}

// NewTestHelper creates a new test helper instance
func NewTestHelper(t *testing.T) *TestHelper {
	// Get the project root directory
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	rootDir := wd
	if strings.HasSuffix(wd, "/integration_test") {
		rootDir = strings.TrimSuffix(wd, "/integration_test")
	}

	binaryPath := "./gh-merge-base-next" // Use local binary

	// Use the actual repository for testing
	repoFlag := "srz-zumix/gh-merge-base-next"

	return &TestHelper{
		rootDir:    rootDir,
		binaryPath: binaryPath,
		repoFlag:   repoFlag,
	}
}

// BuildBinary ensures the binary is available
func (h *TestHelper) BuildBinary(t *testing.T) {
	// Build the binary if it doesn't exist
	cmd := exec.Command("make", "build")
	cmd.Dir = h.rootDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
}

// RunCommand executes the gh-merge-base-next command with given arguments
func (h *TestHelper) RunCommand(t *testing.T, args ...string) (string, string, error) {
	// Add repository flag to all commands
	fullArgs := append([]string{"--repo", h.repoFlag}, args...)

	cmd := exec.Command(h.binaryPath, fullArgs...)
	cmd.Dir = h.rootDir

	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

// AssertCommandSuccess runs a command and asserts it succeeds
func (h *TestHelper) AssertCommandSuccess(t *testing.T, args ...string) string {
	stdout, stderr, err := h.RunCommand(t, args...)
	if err != nil {
		t.Errorf("Command failed: %v\nArgs: %v\nStdout: %s\nStderr: %s", err, args, stdout, stderr)
		return ""
	}

	// Return the actual output (might be in stdout or stderr)
	result := strings.TrimSpace(stdout)
	if result == "" && stderr != "" {
		result = strings.TrimSpace(stderr)
	}

	return result
}

// AssertCommandError runs a command and asserts it fails
func (h *TestHelper) AssertCommandError(t *testing.T, expectedError string, args ...string) string {
	stdout, stderr, err := h.RunCommand(t, args...)
	if err == nil {
		t.Errorf("Expected command to fail but it succeeded\nArgs: %v\nStdout: %s", args, stdout)
		return ""
	}

	errorOutput := strings.TrimSpace(stderr)
	if expectedError != "" && !strings.Contains(errorOutput, expectedError) {
		t.Errorf("Expected error message '%s' not found in stderr: %s", expectedError, errorOutput)
	}

	return errorOutput
}

// GetTestRepository returns the repository used for testing
func (h *TestHelper) GetTestRepository() string {
	return h.repoFlag
}
