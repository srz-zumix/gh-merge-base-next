package mergebasenext

import (
	"context"
	"fmt"

	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/srz-zumix/go-gh-extension/pkg/gh"
	"github.com/srz-zumix/go-gh-extension/pkg/parser"
)

type Client struct {
	client *gh.GitHubClient
	ctx    context.Context
	repo   repository.Repository
}

func NewClient(repo string) (*Client, error) {
	repository, err := parser.Repository(parser.RepositoryInput(repo))
	if err != nil {
		return nil, fmt.Errorf("error parsing repository: %w", err)
	}

	ctx := context.Background()
	client, err := gh.NewGitHubClientWithRepo(repository)
	if err != nil {
		return nil, fmt.Errorf("error creating GitHub client: %w", err)
	}

	return &Client{
		client: client,
		ctx:    ctx,
		repo:   repository,
	}, nil
}
