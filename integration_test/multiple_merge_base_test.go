package integrationtest

/*
## Multiple Merge Base Commit Graph

```text
* d9b05a4 (branch2) I: Branch2 specific development
*   beb86f7 H: Merge branchA into branch2
|\
| | * a7d32e1 (branch1) G: Branch1 specific development
| | *   0e66bd3 F: Merge branchB into branch1
| |/|\
| |/ /
|/| /
* | | 8c7a9d4 (branchB) E: Common work on path B
* | | b265c24 C: Branch B initial work
| | * 24150a4 (branchA) D: Common work on path A
| | * a8184be B: Branch A initial work
| |/
|/
* 120ed78 (main) A: Initial commit (root)
```

## Multiple Merge Base Analysis

When running `git merge-base --all branch1 branch2`, it returns:
- 8c7a9d4b44737dee39a8893358c3cfdd81172ffd (E: Common work on path B)
- 24150a474baddbc32827c30e2147fe6b076be7d2 (D: Common work on path A)

This occurs because both branches have divergent merge paths that create
multiple potential merge bases, representing different common ancestors
that are not ancestors of each other.

## Commit Details

- 120ed78: A: Initial commit (root)
- a8184be: B: Branch A initial work
- b265c24: C: Branch B initial work
- 24150a4: D: Common work on path A
- 8c7a9d4: E: Common work on path B
- 0e66bd3: F: Merge branchB into branch1
- a7d32e1: G: Branch1 specific development
- beb86f7: H: Merge branchA into branch2
- d9b05a4: I: Branch2 specific development
*/

import (
	"testing"
)

// TestMultipleMergeBase tests scenarios where git merge-base --all returns multiple commits
func TestMultipleMergeBase(t *testing.T) {
	testCases := []TestCase{
		{
			Name:  "MultipleMergeBaseBranch1ToBranch2",
			Base:  "testdata/multiple-merge-base/branch1",
			Head:  "testdata/multiple-merge-base/branch2",
			SHA:   "beb86f714f463ead9c78e993715a168d3514a20a", // H: Merge branchA into branch2
			Depth: 2,
			Desc:  "Find next commit with multiple merge bases - should return merge commit",
		},
		{
			Name:  "MultipleMergeBaseBranch2ToBranch1",
			Base:  "testdata/multiple-merge-base/branch2",
			Head:  "testdata/multiple-merge-base/branch1",
			SHA:   "0e66bd36c2ce00729842c094da0cfb14b217533d", // F: Merge branchB into branch1
			Depth: 2,
			Desc:  "Find next commit with multiple merge bases - should return merge commit",
		},
		{
			Name:  "FromSingleMergeBaseToHead",
			Base:  "8c7a9d4b44737dee39a8893358c3cfdd81172ffd", // E: Common work on path B
			Head:  "testdata/multiple-merge-base/branch1",
			SHA:   "a8184be5d737c2751486ea88bcc9151e9ce4b8e1", // B: Branch A initial work
			Depth: 4,
			Desc:  "Find next commit from one of the merge bases to branch1",
		},
		{
			Name:  "FromAnotherMergeBaseToHead",
			Base:  "24150a474baddbc32827c30e2147fe6b076be7d2", // D: Common work on path A
			Head:  "testdata/multiple-merge-base/branch2",
			SHA:   "b265c2435cd016f47413df25f4655962cc5bbaf5", // C: Branch B initial work
			Depth: 4,
			Desc:  "Find next commit from another merge base to branch2",
		},
		{
			Name:  "NavigateBetweenMergeBases",
			Base:  "8c7a9d4b44737dee39a8893358c3cfdd81172ffd", // E: Common work on path B
			Head:  "24150a474baddbc32827c30e2147fe6b076be7d2", // D: Common work on path A
			SHA:   "a8184be5d737c2751486ea88bcc9151e9ce4b8e1", // B: Branch A initial work (common ancestor)
			Depth: 2,
			Desc:  "Navigate between merge bases via common ancestor",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Run(t)
		})
	}
}
