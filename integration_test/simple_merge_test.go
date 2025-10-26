package integrationtest

/*
## Commit Graph

```text
* 904d00b (testdata/simple-merge/feature) D: Feature development 2
* d761e77                                 C: Feature development 1
| * 0eb5947 (testdata/simple-merge/main)  B: Main branch development
|/
* cdddb51                                 A: Initial commit (merge-base)
```

## Commit Details

- cdddb51: A: Initial commit (merge-base of main and feature)
- 0eb5947: B: Main branch development
- d761e77: C: Feature development 1
- 904d00b: D: Feature development 2
*/

import (
	"fmt"
	"strings"
	"testing"
)

// TestSimpleMerge tests the simple merge scenarios with commit hash and depth validation
func TestSimpleMerge(t *testing.T) {
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
			name:     "FindNextFromMergeBaseToMain",
			base:     "cdddb51",
			head:     "testdata/simple-merge/main",
			walkTo:   "head",
			expected: "0eb5947",
			depth:    1,
			desc:     "Find next commit from merge-base (A: Initial commit) to main branch",
		},
		{
			name:     "FindNextFromMergeBaseToFeature",
			base:     "cdddb51",
			head:     "testdata/simple-merge/feature",
			walkTo:   "head",
			expected: "d761e77",
			depth:    2,
			desc:     "Find next commit from merge-base (A: Initial commit) to feature branch",
		},
		{
			name:     "AutoDetectMergeBaseWalkToHead",
			base:     "testdata/simple-merge/main",
			head:     "testdata/simple-merge/feature",
			walkTo:   "head",
			expected: "d761e77",
			depth:    2,
			desc:     "Auto-detect merge-base between main and feature, walk to head (feature)",
		},
		{
			name:     "AutoDetectMergeBaseWalkToBase",
			base:     "testdata/simple-merge/main",
			head:     "testdata/simple-merge/feature",
			walkTo:   "base",
			expected: "0eb5947",
			depth:    1,
			desc:     "Auto-detect merge-base between main and feature, walk to base (main)",
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
