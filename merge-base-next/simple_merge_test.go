package mergebasenext

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
	"testing"
)

// TestSimpleMerge tests the simple merge scenarios with commit hash and depth validation
func TestSimpleMerge(t *testing.T) {
	testCases := []TestCase{
		{
			Name:  "FindNextFromMergeBaseToMain",
			Base:  "cdddb51",
			Head:  "testdata/simple-merge/main",
			SHA:   "0eb59474ced5e6cd338c9ef1406acb4b4522d9fc",
			Depth: 1,
			Desc:  "Find next commit from merge-base (A: Initial commit) to main branch",
		},
		{
			Name:  "FindNextFromMergeBaseToFeature",
			Base:  "cdddb51",
			Head:  "testdata/simple-merge/feature",
			SHA:   "d761e77fe7bbc5fd65e8ee14b8a65ea2ff1f0043",
			Depth: 2,
			Desc:  "Find next commit from merge-base (A: Initial commit) to feature branch",
		},
		{
			Name:  "AutoDetectMergeBaseWalkToHead",
			Base:  "testdata/simple-merge/main",
			Head:  "testdata/simple-merge/feature",
			SHA:   "d761e77fe7bbc5fd65e8ee14b8a65ea2ff1f0043",
			Depth: 2,
			Desc:  "Auto-detect merge-base between main and feature, walk to feature",
		},
		{
			Name:  "AutoDetectMergeBaseWalkToBase",
			Base:  "testdata/simple-merge/feature",
			Head:  "testdata/simple-merge/main",
			SHA:   "0eb59474ced5e6cd338c9ef1406acb4b4522d9fc",
			Depth: 1,
			Desc:  "Auto-detect merge-base between main and feature, walk to main",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Run(t)
		})
	}
}
