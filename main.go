package main

import "camping-backend-with-go/internal/infrastructure/server"

const (
	WEB_PORT int = 3000
)

func main() {
	server.Start(WEB_PORT)
}
