package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gulizay91/template-go-api/internal/routers"
)

var serverError chan error

func InitFiber() *fiber.App {
	fiberConfig := fiber.Config{
		BodyLimit: 30 * 1024 * 1024, // this is the default limit of 4MB
	}
	app := fiber.New(fiberConfig)
	app.Use(recover.New())

	registerRouters(app)

	//log.Fatal(app.Listen(":" + config.Service.Port))
	//if err := app.Listen(":" + config.Service.Port); err != nil {
	//	log.Panic(err)
	//}
	serverError = make(chan error, 1)
	go func() {
		if err := app.Listen(":" + config.Service.Port); err != nil {
			log.Panic(err)
			serverError <- err
		}
	}()

	return app
}

func registerRouters(app *fiber.App) {
	// Create a new Router with RabbitMQ connection
	router := routers.NewRouter(app, config, AmqpConn, AmqpChannel)

	// Register routes
	router.AddRouter()

	log.Debug("Routers Registered.")
}
