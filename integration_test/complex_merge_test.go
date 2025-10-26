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
	"testing"
)

// TestComplexMerge tests complex merge scenarios with multiple branches, merge commits, and depth validation
func TestComplexMerge(t *testing.T) {
	testCases := []TestCase{
		{
			Name:  "FindNextFromMergeBaseToMain",
			Base:  "00506e8",
			Head:  "testdata/complex-merge/main",
			SHA:   "8aa27fa091aa909e06c42dd61c83c1395f034fb8",
			Depth: 3,
			Desc:  "Find next commit from merge-base (A: Initial commit) to main branch with merge commits",
		},
		{
			Name:  "FindNextFromMergeBaseToFeature2",
			Base:  "00506e8",
			Head:  "testdata/complex-merge/feature2",
			SHA:   "8bea32a6b9ff20cbfe949e5e8f82ddbf84197ad1",
			Depth: 2,
			Desc:  "Find next commit from merge-base (A: Initial commit) to feature2 branch",
		},
		{
			Name:  "AutoDetectMergeBaseMainToFeature2WalkToHead",
			Base:  "testdata/complex-merge/main",
			Head:  "testdata/complex-merge/feature2",
			SHA:   "8bea32a6b9ff20cbfe949e5e8f82ddbf84197ad1",
			Depth: 2,
			Desc:  "Auto-detect merge-base between main and feature2, walk to feature2",
		},
		{
			Name:  "AutoDetectMergeBaseMainToFeature2WalkToBase",
			Base:  "testdata/complex-merge/feature2",
			Head:  "testdata/complex-merge/main",
			SHA:   "8aa27fa091aa909e06c42dd61c83c1395f034fb8",
			Depth: 3,
			Desc:  "Auto-detect merge-base between main and feature2, walk to main",
		},
		{
			Name:  "FindNextAfterMergeCommit",
			Base:  "12898c7",
			Head:  "testdata/complex-merge/main",
			SHA:   "b1f8984e815df71ba66cbbbbe7559365d0c82d4f",
			Depth: 1,
			Desc:  "Find next commit after merge commit (E: Merge feature1 into main) to main",
		},
		{
			Name:  "FindNextFromMergeBaseToFeature1",
			Base:  "00506e8",
			Head:  "testdata/complex-merge/feature1",
			SHA:   "75ff397267f653296190ade85624b664920353b4",
			Depth: 1,
			Desc:  "Find next commit from merge-base (A: Initial commit) to feature1 branch",
		},
		{
			Name:  "AutoDetectMergeBaseMainToFeature1WalkToHead",
			Base:  "testdata/complex-merge/main",
			Head:  "testdata/complex-merge/feature1",
			SHA:   "",
			Depth: 0,
			Desc:  "Auto-detect merge-base between main and feature1, walk to feature1 (head)",
		},
		{
			Name:  "AutoDetectMergeBaseMainToFeature1WalkToBase",
			Base:  "testdata/complex-merge/feature1",
			Head:  "testdata/complex-merge/main",
			SHA:   "8aa27fa091aa909e06c42dd61c83c1395f034fb8",
			Depth: 3,
			Desc:  "Auto-detect merge-base between main and feature1, walk to main (base)",
		},
		{
			Name:  "AutoDetectMergeBaseFeature2ToFeature1WalkToHead",
			Base:  "testdata/complex-merge/feature2",
			Head:  "testdata/complex-merge/feature1",
			SHA:   "75ff397267f653296190ade85624b664920353b4",
			Depth: 1,
			Desc:  "Auto-detect merge-base between feature2 and feature1, walk to feature1 (head)",
		},
		{
			Name:  "AutoDetectMergeBaseFeature2ToFeature1WalkToBase",
			Base:  "testdata/complex-merge/feature1",
			Head:  "testdata/complex-merge/feature2",
			SHA:   "8bea32a6b9ff20cbfe949e5e8f82ddbf84197ad1",
			Depth: 2,
			Desc:  "Auto-detect merge-base between feature2 and feature1, walk to feature2 (base)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Run(t)
		})
	}
}
