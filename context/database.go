package context

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDbContext() (*gorm.DB, *sql.DB, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, nil, err
	}

	database, dbError := gorm.Open(mysql.Open(config.ConnectionString), &gorm.Config{})
	if dbError != nil {
		return nil, nil, dbError
	}

	sqlDB, err := database.DB()
	if err != nil {
		panic(err)
	}

	return database, sqlDB, nil
}
