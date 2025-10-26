package integrationtest

/*

## Commit Graph

```text
* b1f8984 (testdata/complex-merge/main)       G: More main development
*   12898c7                                   E: Merge feature1 into main
|\
| * 75ff397 (testdata/complex-merge/feature1) B: Feature1 development
* | 8aa27fa                                   D: Main branch development
|/
| * 66610c6 (testdata/complex-merge/feature2) F: More feature2 development
| * 8bea32a                                   C: Feature2 development
|/
* 00506e8                                     A: Initial commit (merge-base)
```

## Commit Details

- 00506e8: A: Initial commit (common ancestor)
- 75ff397: B: Feature1 development
- 8bea32a: C: Feature2 development
- 8aa27fa: D: Main branch development
- 12898c7: E: Merge feature1 into main (merge commit)
- 66610c6: F: More feature2 development
- b1f8984: G: More main development

*/

import (
	"fmt"
	"strings"
	"testing"
)

// TestComplexMerge tests complex merge scenarios with multiple branches, merge commits, and depth validation
func TestComplexMerge(t *testing.T) {
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
			base:     "00506e8",
			head:     "testdata/complex-merge/main",
			walkTo:   "head",
			expected: "8aa27fa",
			depth:    3,
			desc:     "Find next commit from merge-base (A: Initial commit) to main branch with merge commits",
		},
		{
			name:     "FindNextFromMergeBaseToFeature2",
			base:     "00506e8",
			head:     "testdata/complex-merge/feature2",
			walkTo:   "head",
			expected: "8bea32a",
			depth:    2,
			desc:     "Find next commit from merge-base (A: Initial commit) to feature2 branch",
		},
		{
			name:     "AutoDetectMergeBaseMainToFeature2WalkToHead",
			base:     "testdata/complex-merge/main",
			head:     "testdata/complex-merge/feature2",
			walkTo:   "head",
			expected: "8bea32a",
			depth:    2,
			desc:     "Auto-detect merge-base between main and feature2, walk to feature2 (head)",
		},
		{
			name:     "AutoDetectMergeBaseMainToFeature2WalkToBase",
			base:     "testdata/complex-merge/main",
			head:     "testdata/complex-merge/feature2",
			walkTo:   "base",
			expected: "8aa27fa",
			depth:    3,
			desc:     "Auto-detect merge-base between main and feature2, walk to main (base)",
		},
		{
			name:     "FindNextAfterMergeCommit",
			base:     "12898c7",
			head:     "testdata/complex-merge/main",
			walkTo:   "head",
			expected: "b1f8984",
			depth:    1,
			desc:     "Find next commit after merge commit (E: Merge feature1 into main) to main",
		},
		{
			name:     "FindNextFromMergeBaseToFeature1",
			base:     "00506e8",
			head:     "testdata/complex-merge/feature1",
			walkTo:   "head",
			expected: "75ff397",
			depth:    1,
			desc:     "Find next commit from merge-base (A: Initial commit) to feature1 branch",
		},
		{
			name:     "AutoDetectMergeBaseMainToFeature1WalkToHead",
			base:     "testdata/complex-merge/main",
			head:     "testdata/complex-merge/feature1",
			walkTo:   "head",
			expected: "<no value>",
			depth:    0,
			desc:     "Auto-detect merge-base between main and feature1, walk to feature1 (head)",
		},
		{
			name:     "AutoDetectMergeBaseMainToFeature1WalkToBase",
			base:     "testdata/complex-merge/main",
			head:     "testdata/complex-merge/feature1",
			walkTo:   "base",
			expected: "8aa27fa",
			depth:    3,
			desc:     "Auto-detect merge-base between main and feature1, walk to main (base)",
		},
		{
			name:     "AutoDetectMergeBaseFeature2ToFeature1WalkToHead",
			base:     "testdata/complex-merge/feature2",
			head:     "testdata/complex-merge/feature1",
			walkTo:   "head",
			expected: "75ff397",
			depth:    1,
			desc:     "Auto-detect merge-base between feature2 and feature1, walk to feature1 (head)",
		},
		{
			name:     "AutoDetectMergeBaseFeature2ToFeature1WalkToBase",
			base:     "testdata/complex-merge/feature2",
			head:     "testdata/complex-merge/feature1",
			walkTo:   "base",
			expected: "8bea32a",
			depth:    2,
			desc:     "Auto-detect merge-base between feature2 and feature1, walk to feature2 (base)",
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
