package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gulizay91/template-go-api/routers"
)

var serverError chan error

func InitFiber() *fiber.App {
	fiberConfig := fiber.Config{
		BodyLimit: 30 * 1024 * 1024, // this is the default limit of 4MB
	}
	app := fiber.New(fiberConfig)
	app.Use(recover.New())

	registerRouters(app)

	log.Debugf("Server Now Listen %s:%s", config.Server.Addr, config.Server.Port)
	//log.Fatal(app.Listen(":" + config.Server.Port))
	//if err := app.Listen(":" + config.Server.Port); err != nil {
	//	log.Panic(err)
	//}
	serverError = make(chan error, 1)
	go func() {
		if err := app.Listen(":" + config.Server.Port); err != nil {
			log.Panic(err)
			serverError <- err
		}
	}()

	return app
}

func registerRouters(app *fiber.App) {
	routers.NewRouter(app).AddRouter()

	log.Debug("Routers Registered.")
}
