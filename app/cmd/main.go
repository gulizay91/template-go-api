package main

import (
	"github.com/gulizay91/template-go-api/cmd/services"
)

// @title Template Go API
// @version 1.0
// @description Template Go Api - RESTful
// @termsOfService https://swagger.io/terms/

// @contact.name Güliz AY
// @contact.url https://github.com/gulizay91
// @contact.email gulizay91@gmail.com

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @BasePath /
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Bearer-Token
func main() {
	services.Run()
}
