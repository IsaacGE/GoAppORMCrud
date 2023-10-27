package context

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbContext *gorm.DB

func SetupDbContext() (*sql.DB, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	database, dbError := gorm.Open(mysql.Open(config.ConnectionString), &gorm.Config{})
	if dbError != nil {
		return nil, dbError
	}

	sqlDB, err := database.DB()
	if err != nil {
		panic(err)
	}

	DbContext = database
	return sqlDB, nil
}
