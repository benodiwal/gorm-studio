package database

import (
	"fmt"
	"log"

	"github.com/benodiwal/gorm-studio/pkg/env"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	DB *gorm.DB
	logger *log.Logger
}

func Connect(logger *log.Logger) * Database {
	db_driver := env.Read(env.DB_DRIVER)
	db_host := env.Read(env.DB_HOST)
	db_user := env.Read(env.DB_USER)
	db_password := env.Read(env.DB_PASSWORD)
	db_name := env.Read(env.DB_NAME)
	db_port := env.Read(env.DB_PORT)
	ssl_mode := env.Read(env.DB_SSLMODE)

	db_url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", db_host, db_user, db_password, db_name, db_port, ssl_mode)
	DB, err := gorm.Open(db_driver, db_url)

	if err != nil {
		logger.Panic("database connection error: ", err)
	} else {
		logger.Println("database connection successful: ", db_driver)
	}

	return &Database{
		DB: DB,
		logger: logger,
	}
}
