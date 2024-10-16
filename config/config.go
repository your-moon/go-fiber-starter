package config

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func UseProdConfig() {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("prod")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	color.Green("Using config file: %s", viper.ConfigFileUsed())
}

func UseTestConfig() {
	viper.AddConfigPath("/Users/munkherdeneerdenebat/project/go-fiber-starter/config/")
	viper.SetConfigName("test")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	color.Green("Using config file: %s", viper.ConfigFileUsed())
}
