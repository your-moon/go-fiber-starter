package integrations

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {

	dbConfigUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.name"),
	)

	db, err := gorm.Open(postgres.Open(dbConfigUrl), &gorm.Config{
		PrepareStmt:                              true,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return err
	}

	DB = db
	return nil
}
