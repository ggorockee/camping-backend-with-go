package database

import (
	categoryentity "camping-backend-with-go/internal/domain/entity/category"
	userentity "camping-backend-with-go/internal/domain/entity/user"
	"camping-backend-with-go/pkg/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connection() *gorm.DB {
	MYSQL := os.Getenv("GO_MYSQL")
	var dsn string
	var dbSetting string
	var db *gorm.DB
	var err error

	if MYSQL == "" {
		log.Fatalf("database setting failed...")
	} else {
		dbSetting = "mysql"
		log.Printf("[INFO] %s DB setting enabled...\n", dbSetting)
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Config("DB_USER"),
			config.Config("DB_PASSWORD"),
			config.Config("DB_HOST"),
			config.Config("DB_PORT"),
			config.Config("DB_NAME"),
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	err = db.AutoMigrate(
		&userentity.User{},
		&categoryentity.Category{},
	)
	if err != nil {
		log.Println(err.Error())
	}

	return db
}
