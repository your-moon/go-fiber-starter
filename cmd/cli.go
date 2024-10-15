package cmd

import (
	"fmt"

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
		return
	}

	viper.AddConfigPath("./config/")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().
		StringVar(&cfgFile, "config", "", "config file (default is $CURRENT_DIR/config/dev.yaml)")
}
