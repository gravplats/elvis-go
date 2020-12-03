package cmd

import (
	elvisCmd "github.com/mrydengren/elvis/pkg/cmd"
	"github.com/spf13/cobra"
	"strings"
)

func NewCmdTop() *cobra.Command {
	cmdTop := cobra.Command{
		Use:     "top <artist>",
		Short:   "Generate a playlist from the last.fm top tracks for an artist",
		Example: "  elvis top opeth",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			artist := strings.Join(args, " ")
			// TODO: limit should be a CLI flag
			limit := 10

			elvisCmd.Top(artist, limit)
		},
	}

	return &cmdTop
}

func init() {
	rootCmd.AddCommand(NewCmdTop())
}
