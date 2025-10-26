package mergebasenext

import (
	"fmt"

	"github.com/google/go-github/v73/github"
	"github.com/srz-zumix/go-gh-extension/pkg/gh"
)

type MergeBaseNext struct {
	Next  *github.RepositoryCommit
	Depth int
}

func (c *Client) GetMergeBaseNext(base string, head string) (MergeBaseNext, error) {
	commitsComparison, err := gh.CompareCommits(c.ctx, c.client, c.repo, base, head)
	if err != nil {
		return MergeBaseNext{}, err
	}

	headCommit, err := gh.GetCommit(c.ctx, c.client, c.repo, head)
	if err != nil {
		return MergeBaseNext{}, err
	}
	headRepositoryCommit, err := findCommit(commitsComparison, headCommit.GetSHA())
	if err != nil {
		return MergeBaseNext{}, err
	}

	nextCommit, depth := walkToFirstParent(commitsComparison, headRepositoryCommit, 1)

	return MergeBaseNext{
		Next:  nextCommit,
		Depth: depth,
	}, nil
}

func findCommit(commitsComparison *github.CommitsComparison, sha string) (*github.RepositoryCommit, error) {
	for i := len(commitsComparison.Commits) - 1; i >= 0; i-- {
		commit := commitsComparison.Commits[i]
		if commit.GetSHA() == sha {
			return commit, nil
		}
	}
	return nil, fmt.Errorf("commit not found")
}

func walkToFirstParent(commitsComparison *github.CommitsComparison, commit *github.RepositoryCommit, depth int) (*github.RepositoryCommit, int) {
	if len(commit.Parents) == 0 {
		return commit, depth
	}
	parentCommit, err := findCommit(commitsComparison, commit.Parents[0].GetSHA())
	if err != nil {
		return commit, depth
	}
	return walkToFirstParent(commitsComparison, parentCommit, depth+1)
}
