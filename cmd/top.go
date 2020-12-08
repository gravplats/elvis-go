package cmd

import (
	"github.com/mrydengren/elvis/pkg/cmd/top"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return pflag.ErrHelp
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			artist := strings.Join(args, " ")
			return top.Top(artist, opts.Limit)
		},
	}

	cmdTop.Flags().IntVar(&opts.Limit, "limit", 10, "Number of tracks")

	return &cmdTop
}

func init() {
	rootCmd.AddCommand(NewCmdTop())
}
