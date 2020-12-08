package cmd

import (
	"github.com/mrydengren/elvis/pkg/cmd/yt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"log"
)

func NewCmdYt() *cobra.Command {
	cmdYt := cobra.Command{
		Use:   "yt <value>",
		Short: "Generate a playlist from albums in a YouTube description",
		Example: "" +
			"  elvis yt https://www.youtube.com/watch?v=aEdho6H0hFY\n" +
			"  elvis yt aEdho6H0hFY",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return pflag.ErrHelp
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 1 {
				log.Printf("Ignoring extra arguments: %+v\n", args[1:])
			}
			return yt.Yt(args[0])
		},
	}

	return &cmdYt
}

func init() {
	rootCmd.AddCommand(NewCmdYt())
}
