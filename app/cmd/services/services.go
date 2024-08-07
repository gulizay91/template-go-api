package services

import stdLog "log"

func Run() {
	InitConfig()
	stdLog.Printf("Configuration Initialized for %s", config.Service.Name)

	InitRabbitMQ()

	app := InitFiber()
	RegisterGracefulShutdown(app)
}
