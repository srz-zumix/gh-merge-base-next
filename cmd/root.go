package cmd

import (
	"os"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
	mergebasenext "github.com/srz-zumix/gh-merge-base-next/merge-base-next"
	"github.com/srz-zumix/gh-merge-base-next/version"
	"github.com/srz-zumix/go-gh-extension/pkg/actions"
	"github.com/srz-zumix/go-gh-extension/pkg/render"
)

type Options struct {
	Exporter cmdutil.Exporter
	Repo     string
	WalkTo   string
}

var opts Options
var rootCmd = &cobra.Command{
	Use:     "gh-merge-base-next <base> <head>",
	Short:   "A tool to find the next commit in a merge base",
	Long:    `gh-merge-base-next is a tool to find the next commit in a merge base.`,
	Version: version.Version,
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		base := args[0]
		head := args[1]
		if opts.WalkTo == "base" {
			base, head = head, base
		}
		err := RunMergeBaseNext(cmd, base, head)
		return err
	},
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
	f := rootCmd.Flags()
	f.StringVarP(&opts.Repo, "repo", "R", "", "Target repository in the format 'owner/repo'")
	f.StringVarP(&opts.WalkTo, "walk-to", "T", "head", "Specifies whether the next commit of a merge base should walk to the base or the head")
	cmdutil.AddFormatFlags(rootCmd, &opts.Exporter)
}

func RunMergeBaseNext(cmd *cobra.Command, base string, head string) error {
	client, err := mergebasenext.NewClient(opts.Repo)
	if err != nil {
		return err
	}
	result, err := client.GetMergeBaseNext(base, head)
	if err != nil {
		return err
	}

	renderer := render.NewRenderer(opts.Exporter)
	if opts.Exporter != nil {
		renderer.RenderExportedData(result)
		return nil
	}
	cmd.Println(result.Next.GetSHA())
	return nil
}
