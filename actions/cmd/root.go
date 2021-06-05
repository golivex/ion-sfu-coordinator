package cmd

import (
	"fmt"

	log "github.com/pion/ion-log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "action",
	Short: "action is a collection of small utility to interface with ion-sfu using ion-sdk-go",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	log.Init("info")
	cobra.OnInitialize(initConfig)

	//TODO heere we should have option between cluster or sfu direct and it will be used in all internal commands also
}

func Execute() error {
	return rootCmd.Execute()
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
