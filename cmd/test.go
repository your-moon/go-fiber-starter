package cmd

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/your-moon/go-fiber-starter/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Test bool

var testCmd = &cobra.Command{
	Use: "test",
	RunE: func(cmd *cobra.Command, args []string) error {

		if Prod {
			return errors.New(color.RedString("Cannot run test in production mode"))
		}

		config.UseTestConfig()

		fmt.Printf(color.GreenString("Test mode"))

		//test db connection
		dbConfigUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			viper.GetString("db.host"),
			viper.GetString("db.port"),
			viper.GetString("db.user"),
			viper.GetString("db.password"),
			viper.GetString("db.name"),
		)

		_, err := gorm.Open(postgres.Open(dbConfigUrl), &gorm.Config{
			PrepareStmt:                              true,
			SkipDefaultTransaction:                   true,
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			panic(err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
