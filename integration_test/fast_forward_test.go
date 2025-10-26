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
	"testing"
)

// TestFastForward tests fast-forward scenarios with linear commit history and depth validation
func TestFastForward(t *testing.T) {

	testCases := []TestCase{
		{
			Name:  "FindNextFromMainToFeatureFirstNext",
			Base:  "476d315",
			Head:  "testdata/fast-forward/feature",
			SHA:   "463ec54c18fd7544fcf276aec8a3c1e185b61b6c",
			Depth: 3,
			Desc:  "Find next commit from main to feature (first next)",
		},
		{
			Name:  "FindNextFromIntermediateCommit",
			Base:  "463ec54c18fd7544fcf276aec8a3c1e185b61b6c",
			Head:  "testdata/fast-forward/feature",
			SHA:   "f3295b2c26f9cc8f96e3be8dcf11b34022bc1951",
			Depth: 2,
			Desc:  "Find next commit from intermediate commit (B: Feature work 1)",
		},
		{
			Name:  "FindNextFromSecondToLastCommit",
			Base:  "f3295b2c26f9cc8f96e3be8dcf11b34022bc1951",
			Head:  "testdata/fast-forward/feature",
			SHA:   "6bbcf95e2eb308ae611d6c8ef67e1b46d04d8a5a",
			Depth: 1,
			Desc:  "Find next commit from second to last commit (C: Feature work 2)",
		},
		{
			Name:  "FindNextFromLastCommit",
			Base:  "6bbcf95",
			Head:  "testdata/fast-forward/feature",
			SHA:   "",
			Depth: 0,
			Desc:  "Test when base equals head (no next commit available)",
		},
		{
			Name:  "FindNextFromSameCommit",
			Base:  "476d315",
			Head:  "476d315",
			SHA:   "",
			Depth: 0,
			Desc:  "Test when base equals head (no next commit available)",
		},
		{
			Name:  "AutoDetectMergeBaseMainToFeatureWalkToHead",
			Base:  "testdata/fast-forward/main",
			Head:  "testdata/fast-forward/feature",
			SHA:   "463ec54c18fd7544fcf276aec8a3c1e185b61b6c",
			Depth: 3,
			Desc:  "Auto-detect merge-base between main and feature, walk to feature",
		},
		{
			Name:  "AutoDetectMergeBaseMainToFeatureWalkToBase",
			Base:  "testdata/fast-forward/feature",
			Head:  "testdata/fast-forward/main",
			SHA:   "",
			Depth: 0,
			Desc:  "Auto-detect merge-base between main and feature, walk to main (no next commit)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Run(t)
		})
	}
}
