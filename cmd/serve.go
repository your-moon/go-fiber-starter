package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/your-moon/go-fiber-starter/api"
	"github.com/your-moon/go-fiber-starter/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Prod bool

var serveCmd = &cobra.Command{
	Use: "serve",
	RunE: func(cmd *cobra.Command, args []string) error {
		UseProdConfig()
		if Prod {
			dbConfigUrl := fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				viper.GetString("db.host"),
				viper.GetString("db.port"),
				viper.GetString("db.user"),
				viper.GetString("db.password"),
				viper.GetString("db.name"),
			)
			fmt.Println("ViperConfig: ", dbConfigUrl)

			db, err := gorm.Open(postgres.Open(dbConfigUrl), &gorm.Config{
				PrepareStmt:                              true,
				SkipDefaultTransaction:                   true,
				DisableForeignKeyConstraintWhenMigrating: true,
			})
			if err != nil {
				panic(err.Error())
			}

			services.DB = db

			app := api.Init()
			app.Listen(viper.GetString("server.host") + ":" + viper.GetString("server.port"))
		} else {
			fmt.Println("Running in development mode")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolVarP(&Prod, "prod", "p", false, "Run in production mode")
}

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
