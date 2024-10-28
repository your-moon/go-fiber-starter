package config

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

var InitCfgFile string

func UseConfig(name string) {
	if InitCfgFile != "" {
		return
	}

	viper.AddConfigPath("./config/")
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	color.Green("Using config file: %s", viper.ConfigFileUsed())
}

func UseInitConfig() {
	if InitCfgFile != "" {
		viper.SetConfigFile(InitCfgFile)
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		color.Green("Using config file: %s", viper.ConfigFileUsed())
		return
	}
}
