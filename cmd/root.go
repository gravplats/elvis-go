package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/mrydengren/elvis/pkg/api/lastfm"
	"github.com/mrydengren/elvis/pkg/api/setlistfm"
	"github.com/mrydengren/elvis/pkg/config"
	"github.com/mrydengren/elvis/pkg/debug"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
	"time"
)

type RootOptions struct {
	Debug bool
}

const (
	SPOTIFY_ID     = "SPOTIFY_ID"
	SPOTIFY_SECRET = "SPOTIFY_SECRET"
)

var rootOpts = RootOptions{
	Debug: false,
}

var rootCmd = cobra.Command{
	Use:   "elvis",
	Short: "A Spotify playlist generator",
	Long:  "Elvis is a CLI for generating Spotify playlists from various web APIs.",
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		debug.DumpInput()
		if dir, ok := debug.GetDebugDir(); ok {
			fmt.Printf("\nWrote debug information to %s.\n", dir)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.SetConfigName(".elvis")
	viper.SetConfigType("json")

	viper.AddConfigPath(home)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Unable to read config file.")
	}

	var cfg config.Config
	viper.Unmarshal(&cfg)

	// The lastfm API wrapper reads these, if present.
	os.Setenv(lastfm.ApiKey, cfg.Lastfm.Key)
	os.Setenv(lastfm.ApiSecret, cfg.Lastfm.Secret)
	// The setlistfm API wrapper reads these, if present.
	os.Setenv(setlistfm.ApiKey, cfg.Setlistfm.Key)
	// The Spotify API wrapper reads these, if present.
	os.Setenv(SPOTIFY_ID, cfg.Spotify.Id)
	os.Setenv(SPOTIFY_SECRET, cfg.Spotify.Secret)
}

func initDebug() {
	os.Setenv(debug.DEBUG, strconv.FormatBool(rootOpts.Debug))
	os.Setenv(debug.DEBUG_SESSION_ID, strconv.FormatInt(time.Now().Unix(), 10))
}

func init() {
	cobra.OnInitialize(
		initConfig,
		initDebug,
	)

	rootCmd.PersistentFlags().BoolVar(&rootOpts.Debug, "debug", false, "Output debug information")
	rootCmd.PersistentFlags().MarkHidden("debug")
}
