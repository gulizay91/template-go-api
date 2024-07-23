package services

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/gulizay91/template-go-api/internal/models"
	"github.com/gulizay91/template-go-api/pkg/rabbitmq"
)

// BrokerManager Global channel manager instance
var BrokerManager *rabbitmq.MessageBrokerManager

func InitRabbitMQ() {
	var err error
	BrokerManager, err = rabbitmq.NewMessageBrokerManager(config.RabbitMQ.Hosts, config.RabbitMQ.Username, config.RabbitMQ.Password)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ connection: %v", err)
	}
	log.Info("RabbitMQ initialized.")

	// Start monitoring the RabbitMQ connection
	go BrokerManager.MonitorConnection()

	// Declare exchanges, queues
	RegisterQueue(string(rabbitmq.Fanout), string(models.TemplateExchangeName), "", string(models.TemplateMessageQueue), true)
}

// RegisterQueue Declare queue
func RegisterQueue(exchangeType string, exchangeName string, routingKey string, queueName string, hasDlq bool) {
	amqpChannel, err := BrokerManager.GetChannel(exchangeName)
	if err != nil {
		log.Fatalf("Failed to get RabbitMQ channel for exchange %s: %v", exchangeName, err)
	}
	rabbitmq.DeclareExchange(amqpChannel, exchangeName, exchangeType)

	// Declare queue
	if hasDlq {
		//rabbitmq.DeleteQueueIfExistsAndEmpty(AmqpChannel, queueName+DeadLetterQueueSuffix)
		rabbitmq.DeclareQueue(amqpChannel, queueName+rabbitmq.DeadLetterQueueSuffix, true, false, false, false, nil, false)
	}
	//rabbitmq.DeleteQueueIfExistsAndEmpty(AmqpChannel, queueName)
	rabbitmq.DeclareQueue(amqpChannel, queueName, true, false, false, false, nil, hasDlq)

	// Bind queue to exchange
	rabbitmq.BindQueue(amqpChannel, queueName, routingKey, exchangeName)

	log.Infof("Registered Publisher for %s exchange, %s queue", exchangeName, queueName)
}
