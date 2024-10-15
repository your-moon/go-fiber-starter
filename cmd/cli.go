package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var rootCmd = &cobra.Command{}

func Execute() {
	rootCmd.Execute()
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./config/")
		switch {
		case Prod:
			viper.SetConfigName("prod")
		case Test:
			viper.SetConfigName("test")
		default:
			viper.SetConfigName("dev")
		}
		viper.SetConfigType("yaml")

		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		viper.OnConfigChange(func(e fsnotify.Event) {
			color.Yellow("Config file changed: %s", e.Name)
		})

		viper.WatchConfig()

		color.Green("Using config file: %s", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().
		StringVar(&cfgFile, "config", "", "config file (default is $CURRENT_DIR/config/dev.yaml)")
}
