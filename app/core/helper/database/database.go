package database

import (
	"camping-backend-with-go/config"
	"fmt"
	"log"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(config *config.Config) *gorm.DB {
	gormConfig := newGormConfig()

	values := url.Values{}
	values.Add("charset", "utf8mb4")
	values.Add("parseTime", "true")
	query := values.Encode()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		config.Infra.Db.User,
		config.Infra.Db.Password,
		config.Infra.Db.Host,
		config.Infra.Db.Port,
		config.Infra.Db.DBName,
		query,
	)

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	// database
	// db.AutoMigrate()

	return db
}

func newGormConfig() *gorm.Config {
	return &gorm.Config{}
}
