package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/your-moon/go-fiber-starter/config"
	"github.com/your-moon/go-fiber-starter/internal/api"
	"github.com/your-moon/go-fiber-starter/internal/services"
)

var Prod bool

var serveCmd = &cobra.Command{
	Use: "serve",
	RunE: func(cmd *cobra.Command, args []string) error {

		if config.InitCfgFile != "" {
			if err := run(); err != nil {
				return err
			}
		}

		if Prod {
			if err := runProd(); err != nil {
				return err
			}
		} else {
			if err := runDev(); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolVarP(&Prod, "prod", "p", false, "Run in production mode")
}

func run() error {
	fmt.Println("Config mode")

	if err := services.InitDB(); err != nil {
		return err
	}

	app := api.Init()
	app.Listen(viper.GetString("server.host") + ":" + viper.GetString("server.port"))
	return nil
}

func runDev() error {
	config.UseConfig("dev")
	fmt.Println("Running in development mode")

	if err := services.InitDB(); err != nil {
		return err
	}

	app := api.Init()
	app.Listen(viper.GetString("server.host") + ":" + viper.GetString("server.port"))
	return nil
}

func runProd() error {
	config.UseConfig("prod")

	fmt.Println("Running in production mode")

	if err := services.InitDB(); err != nil {
		return err
	}

	app := api.Init()
	app.Listen(viper.GetString("server.host") + ":" + viper.GetString("server.port"))
	return nil
}

// name cases: prod, dev, test
