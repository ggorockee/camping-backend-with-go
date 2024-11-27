package proxy

import (
	"camping-backend-with-go/pkg/config"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	gitmysql "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var dsn string

func GetProxyDatabase() string {
	enabled := os.Getenv("PROXY")
	if enabled == "true" {
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
		log.Println("[INFO] Proxy setting enabled...")
		dsn = fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=custom",
			config.Config("DB_USER"),
			config.Config("DB_HOST"),
			config.Config("DB_PORT"),
			config.Config("DB_NAME"),
		)
		return dsn
	}
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	return dsn
}
