package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/mrydengren/elvis/pkg/api/lastfm"
	"github.com/mrydengren/elvis/pkg/api/setlistfm"
	"github.com/mrydengren/elvis/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	SPOTIFY_ID     = "SPOTIFY_ID"
	SPOTIFY_SECRET = "SPOTIFY_SECRET"
)

var rootCmd = cobra.Command{
	Use:   "elvis",
	Short: "A Spotify playlist generator",
	Long:  "Elvis is a CLI for generating Spotify playlists from various web APIs.",
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

func init() {
	cobra.OnInitialize(initConfig)
}
