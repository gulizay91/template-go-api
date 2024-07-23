package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	configs "github.com/gulizay91/template-go-api/config"
	"github.com/gulizay91/template-go-api/internal/models"
	"github.com/gulizay91/template-go-api/pkg/utils"
	"github.com/streadway/amqp"
)

// TemplateHandler struct for handling template-related requests
type TemplateHandler struct {
	AppConfig   *configs.Config
	AmqpConn    *amqp.Connection
	AmqpChannel *amqp.Channel
}

// NewTemplateHandler creates a new instance of TemplateHandler
func NewTemplateHandler(config *configs.Config, amqpConn *amqp.Connection, amqpChannel *amqp.Channel) *TemplateHandler {
	return &TemplateHandler{
		AppConfig:   config,
		AmqpConn:    amqpConn,
		AmqpChannel: amqpChannel,
	}
}

var templateArray []models.Template

// GetTemplate godoc
// @Summary get template
// @Description get template
// @Tags templates
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Template
// @Router /api/v1/template [get]
func (h *TemplateHandler) GetTemplate(c *fiber.Ctx) error {

	templateArray = append(templateArray, models.Template{ID: 1, Name: "Name1", Message: "Message1"})
	templateArray = append(templateArray, models.Template{ID: 2, Name: "Name2", Message: "Message2"})
	templateArray = append(templateArray, models.Template{ID: 3, Name: "Name3", Message: "Message3"})

	return c.JSON(templateArray)
}

// SendTemplateMessage godoc
// @Summary send template message
// @Description send message to rabbitmq queue
// @Tags templates
// @Accept json
// @Produce json
// @Param   message body  models.Message true "Message" example({ "body": "{\"message\": \"template message body\"}" })
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/template/message [post]
func (h *TemplateHandler) SendTemplateMessage(c *fiber.Ctx) error {
	var msg models.Message

	if err := c.BodyParser(&msg); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Validate JSON format
	if err := utils.ValidateJSON(msg.Body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	err := h.AmqpChannel.Publish(
		string(models.TemplateExchangeName), // exchange
		"",                                  // routing key
		false,                               // mandatory
		false,                               // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(msg.Body),
		})
	if err != nil {
		log.Errorf("Failed to publish a message: %s", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send message"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Message sent successfully"})
}
