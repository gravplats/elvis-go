package cmd

import (
	"github.com/mrydengren/elvis/pkg/cmd/top"
	"github.com/spf13/cobra"
	"strings"
)

type TopOptions struct {
	Limit int
}

func NewCmdTop() *cobra.Command {
	opts := TopOptions{}

	cmdTop := cobra.Command{
		Use:     "top <artist>",
		Short:   "Generate a playlist from the last.fm top tracks for an artist",
		Example: "  elvis top opeth",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			artist := strings.Join(args, " ")
			top.Top(artist, opts.Limit)
		},
	}

	cmdTop.Flags().IntVar(&opts.Limit, "limit", 10, "Number of tracks")

	return &cmdTop
}

func init() {
	rootCmd.AddCommand(NewCmdTop())
}
