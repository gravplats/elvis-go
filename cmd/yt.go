package cmd

import (
	"github.com/mrydengren/elvis/pkg/cmd/yt"
	"github.com/spf13/cobra"
)

func NewCmdYt() *cobra.Command {
	cmdYt := cobra.Command{
		Use:   "yt <value>",
		Short: "Generate a playlist from albums in a YouTube description",
		Example: "" +
			"  elvis yt https://www.youtube.com/watch?v=aEdho6H0hFY\n" +
			"  elvis yt aEdho6H0hFY",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			yt.Yt(args[0])
		},
	}

	return &cmdYt
}

func init() {
	rootCmd.AddCommand(NewCmdYt())
}
