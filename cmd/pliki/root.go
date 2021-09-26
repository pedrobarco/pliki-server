package main

import (
	"fmt"
	"pliki-server/internal/version"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	semver = version.GetVersion()

	cmd = &cobra.Command{
		Version: semver,
		Use:     "pk",
		Short:   "A box for ephemeral files",
		Long: `Pliki-server is a temporary file storage server built in Go and pliki is its CLI version.
        Complete documentation is available at https://github.com/pedrobarco/pliki-server.`,
	}
)

func init() {
	cobra.OnInitialize(loadAppConfig)

	// set version template
	cmd.SetVersionTemplate(fmt.Sprintf("v%s\n", semver))

	// start
}

func loadAppConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/pliki-server")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error creating config: %s", err))
	}

	/*
	   var acfg *configs.AppConfig
	   if err := viper.Unmarshal(&acfg); err != nil {
	       panic(fmt.Errorf("Fatal error loading config: %s \n", err))
	   }
	*/
}
