package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/your-moon/go-fiber-starter/api"
	"github.com/your-moon/go-fiber-starter/config"
	"github.com/your-moon/go-fiber-starter/services"
)

var Prod bool

var serveCmd = &cobra.Command{
	Use: "serve",
	RunE: func(cmd *cobra.Command, args []string) error {

		if config.InitCfgFile != "" {
			run()
		}

		if Prod {
			runProd()
		} else {
			runDev()
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolVarP(&Prod, "prod", "p", false, "Run in production mode")
}

func run() {
	fmt.Println("Config mode")

	services.InitDB()

	app := api.Init()
	app.Listen(viper.GetString("server.host") + ":" + viper.GetString("server.port"))
}

func runDev() {
	config.UseConfig("dev")
	fmt.Println("Running in development mode")

	services.InitDB()

	app := api.Init()
	app.Listen(viper.GetString("server.host") + ":" + viper.GetString("server.port"))
}

func runProd() {
	config.UseConfig("prod")

	fmt.Println("Running in production mode")

	services.InitDB()

	app := api.Init()
	app.Listen(viper.GetString("server.host") + ":" + viper.GetString("server.port"))
}

// name cases: prod, dev, test
