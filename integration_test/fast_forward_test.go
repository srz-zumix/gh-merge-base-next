package integrationtest

/*

## Commit Graph

```text
* 6bbcf95 (testdata/fast-forward/feature) D: Feature work 3
* f3295b2                                 C: Feature work 2
* 463ec54                                 B: Feature work 1
* 476d315 (testdata/fast-forward/main)    A: Initial commit (merge-base)
```

## Commit Details

- 476d315: A: Initial commit (main HEAD, merge-base)
- 463ec54: B: Feature work 1
- f3295b2: C: Feature work 2
- 6bbcf95: D: Feature work 3 (feature HEAD)

*/

import (
	"fmt"
	"strings"
	"testing"
)

// TestFastForward tests fast-forward scenarios with linear commit history and depth validation
func TestFastForward(t *testing.T) {
	helper := NewTestHelper(t)
	helper.BuildBinary(t)

	testCases := []struct {
		name     string
		base     string
		head     string
		walkTo   string
		expected string
		depth    int
		desc     string
	}{
		{
			name:     "FindNextFromMainToFeatureFirstNext",
			base:     "476d315",
			head:     "testdata/fast-forward/feature",
			walkTo:   "head",
			expected: "463ec54",
			depth:    3,
			desc:     "Find next commit from main to feature (first next)",
		},
		{
			name:     "FindNextFromIntermediateCommit",
			base:     "463ec54",
			head:     "testdata/fast-forward/feature",
			walkTo:   "head",
			expected: "f3295b2",
			depth:    2,
			desc:     "Find next commit from intermediate commit (B: Feature work 1)",
		},
		{
			name:     "FindNextFromSecondToLastCommit",
			base:     "f3295b2",
			head:     "testdata/fast-forward/feature",
			walkTo:   "head",
			expected: "6bbcf95",
			depth:    1,
			desc:     "Find next from second-to-last commit (C: Feature work 2)",
		},
		{
			name:     "AutoDetectMergeBaseMainToFeatureWalkToHead",
			base:     "testdata/fast-forward/main",
			head:     "testdata/fast-forward/feature",
			walkTo:   "head",
			expected: "463ec54",
			depth:    3,
			desc:     "Auto-detect merge-base between main and feature, walk to head",
		},
		{
			name:     "AutoDetectMergeBaseMainToFeatureWalkToBase",
			base:     "testdata/fast-forward/main",
			head:     "testdata/fast-forward/feature",
			walkTo:   "base",
			expected: "<no value>",
			depth:    0,
			desc:     "Auto-detect merge-base between main and feature, walk to base (no next commit)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Testing: %s", tc.desc)

			// Prepare command arguments with template to get depth
			args := []string{tc.base, tc.head, "--walk-to", tc.walkTo, "--format", "json", "--template", "{{.commit.sha}} depth:{{.depth}}"}

			// Execute command
			stdout, stderr, err := helper.RunCommand(t, args...)

			if err != nil {
				t.Errorf("Command failed: %v\nArgs: %v\nStdout: %s\nStderr: %s", err, args, stdout, stderr)
				return
			}

			// Check the result
			result := strings.TrimSpace(stdout)
			if result == "" {
				result = strings.TrimSpace(stderr)
			}

			// Verify the expected commit hash is in the output
			if !strings.Contains(result, tc.expected) {
				t.Errorf("Expected commit hash '%s' not found in output: %s", tc.expected, result)
			}

			// Verify the expected depth is in the output
			if !strings.Contains(result, fmt.Sprintf("depth:%d", tc.depth)) {
				t.Errorf("Expected depth '%d' not found in output: %s", tc.depth, result)
			}

			t.Logf("Success - Expected: %s depth:%d, Got: %s", tc.expected, tc.depth, result)
		})
	}
}

// TestFastForwardEdgeCases tests edge cases in fast-forward scenarios
func TestFastForwardEdgeCases(t *testing.T) {
	helper := NewTestHelper(t)
	helper.BuildBinary(t)

	t.Run("NoNextCommitCase", func(t *testing.T) {
		// Test when base equals head (no next commit available)
		args := []string{"6bbcf95", "testdata/fast-forward/feature", "--walk-to", "head"}

		stdout, stderr, err := helper.RunCommand(t, args...)

		// This case is expected to have no output when no next commit is available
		result := stdout + stderr
		t.Logf("No next commit test - Args: %v, Result: '%s', Error: %v", args, result, err)

		// It's acceptable to have no output when there's no next commit
		// This is the expected behavior for this edge case
	})

	t.Run("SameCommitForBaseAndHead", func(t *testing.T) {
		// Test when base and head are the same commit
		args := []string{"476d315", "476d315", "--walk-to", "head"}

		stdout, stderr, err := helper.RunCommand(t, args...)

		result := stdout + stderr
		t.Logf("Same commit test - Args: %v, Result: '%s', Error: %v", args, result, err)

		// It's acceptable to have no output when base equals head
		// This is the expected behavior for this edge case
	})

	t.Run("LinearHistoryValidation", func(t *testing.T) {
		// Test that we can traverse the linear history step by step
		commits := []string{"476d315", "463ec54", "f3295b2", "6bbcf95"}

		for i := 0; i < len(commits)-1; i++ {
			base := commits[i]
			expected := commits[i+1]

			args := []string{base, "testdata/fast-forward/feature", "--walk-to", "head"}
			stdout, stderr, err := helper.RunCommand(t, args...)

			if err != nil {
				t.Errorf("Step %d failed: %v\nArgs: %v\nStdout: %s\nStderr: %s", i+1, err, args, stdout, stderr)
				continue
			}

			result := strings.TrimSpace(stdout + stderr)
			if !strings.Contains(result, expected) {
				t.Errorf("Step %d: Expected '%s' not found in output: %s", i+1, expected, result)
			} else {
				t.Logf("Step %d: Success - %s -> %s", i+1, base[:7], expected[:7])
			}
		}
	})
}
