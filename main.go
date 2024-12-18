package main

import (
	_ "camping-backend-with-go/docs"
	"camping-backend-with-go/internal/infrastructure/server"
)

const (
	WEB_PORT int = 3000
)

// @title ggocamping App
// @version 1.0
// @description This is an API for ggocamping Application
// @contact.name ggorockee
// @contact.email ggorockee@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	server.Start(WEB_PORT)
}
