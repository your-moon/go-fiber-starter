package cmd

import (
	"github.com/spf13/cobra"

	"github.com/your-moon/go-fiber-starter/config"
)

var rootCmd = &cobra.Command{}

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().
		StringVar(&config.InitCfgFile, "config", "", "config file (default is $CURRENT_DIR/config/dev.yaml)")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	config.UseInitConfig()
}
