package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/swagger"
	// docs are generated by Swag CLI, you have to import them.
	// replace with your own docs folder, usually "github.com/username/reponame/docs"
	_ "github.com/gulizay91/template-go-api/docs"
)

type Router struct {
	appRouter *fiber.App
}

func NewRouter(appRouter *fiber.App) *Router {
	return &Router{
		appRouter: appRouter,
	}
}

func (router *Router) AddRouter() {
	// Middleware
	router.appRouter.Use(recover.New())
	router.appRouter.Use(cors.New())

	// Routes
	router.appRouter.Get("/swagger/*", swagger.HandlerDefault)
	router.appRouter.Get("/health", HealthCheck)
	router.appRouter.Get("/ready", ReadyCheck)

	// Create routes group.
	route := router.appRouter.Group("/api/v1")

	route.Get("/template", GetTemplate)
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"status": "✅ Server is up and running!",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}

// ReadyCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /ready [get]
func ReadyCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"status": "✅ Server is ready!",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}

type Template struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

var templateArray []Template

// GetTemplate godoc
// @Summary get template
// @Description get template
// @Tags templates
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/template [get]
func GetTemplate(c *fiber.Ctx) error {

	templateArray = append(templateArray, Template{ID: 1, Name: "Name1", Message: "Message1"})
	templateArray = append(templateArray, Template{ID: 2, Name: "Name2", Message: "Message2"})
	templateArray = append(templateArray, Template{ID: 3, Name: "Name3", Message: "Message3"})

	return c.JSON(templateArray)
}
