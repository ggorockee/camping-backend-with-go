package database

import (
	"camping-backend-with-go/internal/domain/entity"
	"net/url"

	"camping-backend-with-go/pkg/config"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	MYSQL := os.Getenv("GO_MYSQL")
	var dsn string
	var dbSetting string
	var db *gorm.DB
	var err error

	params := url.Values{}
	params.Add("charset", "utf8mb4")
	params.Add("parseTime", "True")
	params.Add("loc", "Asia/Seoul")

	if MYSQL == "" {
		log.Fatalf("database setting failed...")
	} else {
		dbSetting = "mysql"
		log.Printf("[INFO] %s DB setting enabled...\n", dbSetting)
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
			url.QueryEscape(config.Config("DB_USER")),
			url.QueryEscape(config.Config("DB_PASSWORD")),
			config.Config("DB_HOST"),
			config.Config("DB_PORT"),
			config.Config("DB_NAME"),
			params.Encode(),
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	err = db.AutoMigrate(
		&entity.User{},
		&entity.Category{},
		&entity.Amenity{},
		&entity.Review{},
		&entity.Spot{},
		&entity.WishList{},
	)
	if err != nil {
		log.Println(err.Error())
	}

	//db = db.Debug()

	return db
}
