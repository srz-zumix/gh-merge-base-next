package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/srz-zumix/gh-merge-base-next/version"
	"github.com/srz-zumix/go-gh-extension/pkg/actions"
)

var rootCmd = &cobra.Command{
	Use:     "gh-merge-base-next",
	Short:   "A tool to find the next commit in a merge base",
	Long:    `gh-merge-base-next is a tool to find the next commit in a merge base.`,
	Version: version.Version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	if actions.IsRunsOn() {
		rootCmd.SetErrPrefix(actions.GetErrorPrefix())
	}
}
