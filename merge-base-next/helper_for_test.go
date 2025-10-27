package mergebasenext

import (
	"strings"
	"testing"
)

var testRepository = "srz-zumix/gh-merge-base-next"

var testClient *Client

func getTestClient(t *testing.T) *Client {
	if testClient != nil {
		return testClient
	}
	c, err := NewClient(testRepository)
	if err != nil {
		t.Fatalf("Failed to create mergebasenext client: %v", err)
	}
	testClient = c
	return testClient
}

type TestCase struct {
	Name  string
	Base  string
	Head  string
	SHA   string
	Depth int
	Desc  string
}

type ErrorTestCase struct {
	Name  string
	Base  string
	Head  string
	Error string
	Desc  string
}

func (tc *TestCase) Run(t *testing.T) {
	client := getTestClient(t)
	result, err := client.GetMergeBaseNext(tc.Base, tc.Head)
	if err != nil {
		t.Fatalf("GetMergeBaseNext failed: %v", err)
	}
	if result.SHA != tc.SHA {
		t.Errorf("Expected SHA %s, got %s", tc.SHA, result.SHA)
	}
	if result.Depth != tc.Depth {
		t.Errorf("Expected Depth %d, got %d", tc.Depth, result.Depth)
	}
}

func (etc *ErrorTestCase) Run(t *testing.T) {
	client := getTestClient(t)
	_, err := client.GetMergeBaseNext(etc.Base, etc.Head)
	if err == nil {
		t.Fatalf("Expected error but got none")
	}
	if !strings.Contains(err.Error(), etc.Error) {
		t.Errorf("Expected error to contain '%s', got '%s'", etc.Error, err.Error())
	}
}
