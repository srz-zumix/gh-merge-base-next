package mergebasenext

/*
## Cross-Merge Commit Graph

```text
*   1c97b65 (testdata/cross-merge/feature1) M2: Merge feature2 into feature1 (cross-merge point)
|\
| * 46bf43e (testdata/cross-merge/feature2) F2: Feature2 development
| *   18a2ae8 G: Merge feature1 into feature2 (first cross-merge)
| |\
| * | 8e030f4 (testdata/cross-merge/main) C: Main development 2
* | | 2ba409a F1: Feature1 development
| |/
|/|
* | 74ec9ee B: Feature1 initial development
|/
* 9b93392 A: Initial commit (common ancestor)
```

## Commit Details

- 9b93392: A: Initial commit (common ancestor)
- 74ec9ee: B: Feature1 initial development
- 8e030f4: C: Main development 2
- 18a2ae8: G: Merge feature1 into feature2 (first cross-merge)
- 2ba409a: F1: Feature1 development
- 46bf43e: F2: Feature2 development
- 1c97b65: M2: Merge feature2 into feature1 (cross-merge point)
*/

import (
	"testing"
)

// TestCrossMerge tests cross-merge scenarios with complex merge relationships
func TestCrossMerge(t *testing.T) {
	testCases := []TestCase{
		{
			Name:  "CrossMergeFromCommonAncestorToFeature1",
			Base:  "9b93392",
			Head:  "testdata/cross-merge/feature1",
			SHA:   "74ec9ee6dd54286b15638541a9590d770cf039f2",
			Depth: 3,
			Desc:  "Find next commit from common ancestor to final cross-merge branch",
		},
		{
			Name:  "CrossMergeFromCommonAncestorToFeature2",
			Base:  "9b93392",
			Head:  "testdata/cross-merge/feature2",
			SHA:   "8e030f4bf6138adbb1a8c7c67d879bd6433c5130",
			Depth: 3,
			Desc:  "Find next commit from common ancestor to feature2 branch",
		},
		{
			Name:  "BetweenCrossMergePoints",
			Base:  "2ba409ae3a792d20d3425b4b596515b2996301c7",
			Head:  "testdata/cross-merge/feature1",
			SHA:   "1c97b651f23907fd29956141d3e7a6bbc8c24671",
			Depth: 1,
			Desc:  "Find next commit between cross merge points",
		},
		{
			Name:  "AutoDetectCrossMergeMainToFeature1",
			Base:  "testdata/cross-merge/main",
			Head:  "testdata/cross-merge/feature1",
			SHA:   "74ec9ee6dd54286b15638541a9590d770cf039f2",
			Depth: 3,
			Desc:  "Auto-detect cross merge from main to feature1",
		},
		{
			Name:  "CrossMergeMainToFeature2",
			Base:  "testdata/cross-merge/main",
			Head:  "testdata/cross-merge/feature2",
			SHA:   "18a2ae89310b28343e7cd19200a8f2f565864d87",
			Depth: 2,
			Desc:  "Find cross merge path from main to feature2",
		},
		{
			Name:  "CrossMergeFromFirstMergePointToFeature2",
			Base:  "18a2ae89310b28343e7cd19200a8f2f565864d87",
			Head:  "testdata/cross-merge/feature2",
			SHA:   "46bf43eb956daf9599b77d4177f00b7e47c1f9d5",
			Depth: 1,
			Desc:  "Navigate from first cross-merge point to feature2 final commit",
		},
		{
			Name:  "ComplexCrossMergeDepthTest",
			Base:  "74ec9ee6dd54286b15638541a9590d770cf039f2",
			Head:  "testdata/cross-merge/feature2",
			SHA:   "8e030f4bf6138adbb1a8c7c67d879bd6433c5130",
			Depth: 3,
			Desc:  "Test depth calculation in complex cross-merge",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Run(t)
		})
	}
}
