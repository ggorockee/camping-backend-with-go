package main

import (
	"camping-backend-with-go/api/routes"
	"camping-backend-with-go/pkg/auth"
	"camping-backend-with-go/pkg/config"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/healthcheck"
	"camping-backend-with-go/pkg/spot"
	"camping-backend-with-go/pkg/user"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	gitmysql "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := databaseConnection()

	healthcheckService := healthcheck.NewService()

	userRepo := user.NewRepo(db)
	userService := user.NewService(userRepo)

	authRepo := auth.NewRepo(db)
	authService := auth.NewService(authRepo)

	spotRepo := spot.NewRepo(db)
	spotService := spot.NewService(spotRepo)

	app := fiber.New()
	app.Use(cors.New())

	v1 := app.Group("/v1")

	routes.UserRouter(v1, userService)
	routes.AuthRouter(v1, authService)

	routes.SpotRouter(v1, spotService)
	routes.HealthCheckRouter(v1, healthcheckService)
	log.Fatal(app.Listen(":3000"))
}

func databaseConnection() *gorm.DB {

	PROXY := os.Getenv("PROXY")
	var dsn string
	if PROXY == "true" {
		// CA 인증서 로드
		caCert, err := os.ReadFile(config.Config("CA_FILE"))
		if err != nil {
			log.Fatalf("Could not read CA certificate: %v", err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		// 클라이언트 인증서 및 키 로드
		cert, err := tls.LoadX509KeyPair(config.Config("CERT_FILE"), config.Config("KEY_FILE"))
		if err != nil {
			log.Fatalf("Could not load client key pair: %v", err)
		}

		// TLS 설정
		tlsConfig := &tls.Config{
			RootCAs:      caCertPool,
			Certificates: []tls.Certificate{cert},
		}

		// TLS 설정 등록
		err = gitmysql.RegisterTLSConfig("custom", tlsConfig)
		if err != nil {
			log.Fatalf("cannot rester TLSConfig %v\n", err)
		}

		//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		dsn = fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=custom",
			config.Config("DB_USER"),
			config.Config("DB_HOST"),
			config.Config("DB_PORT"),
			config.Config("DB_NAME"),
		)
	}

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	err = db.AutoMigrate(
		&entities.Spot{},
		&entities.User{},
	)
	if err != nil {
		log.Println(err.Error())
	}

	return db
}
