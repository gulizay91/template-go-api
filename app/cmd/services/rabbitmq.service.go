package services

import (
	"github.com/gulizay91/template-go-api/internal/models"
	"github.com/gulizay91/template-go-api/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

var AmqpConn *amqp.Connection
var AmqpChannel *amqp.Channel

func InitRabbitMQ() {
	// Connect to RabbitMQ
	AmqpConn = rabbitmq.ConnectRabbitMQ(config.RabbitMQ.Hosts, config.RabbitMQ.Username, config.RabbitMQ.Password)

	// Create a channel from the connection
	AmqpChannel = rabbitmq.CreateChannel(AmqpConn)

	// Declare exchange
	rabbitmq.DeclareExchange(AmqpChannel, string(models.TemplateExchangeName), string(rabbitmq.Fanout))

	// Delete queues if they exist and are empty
	//rabbitmq.DeleteQueueIfExistsAndEmpty(AmqpChannel, string(models.TemplateMessageQueue))
	//rabbitmq.DeleteQueueIfExistsAndEmpty(AmqpChannel, string(models.TemplateMessageQueue)+DeadLetterQueueSuffix)

	// Declare queues
	rabbitmq.DeclareQueue(AmqpChannel, string(models.TemplateMessageQueue)+rabbitmq.DeadLetterQueueSuffix, true, false, false, false, nil, false)
	rabbitmq.DeclareQueue(AmqpChannel, string(models.TemplateMessageQueue), true, false, false, false, nil, true)

	// Bind queue to exchange
	rabbitmq.BindQueue(AmqpChannel, string(models.TemplateMessageQueue), "", string(models.TemplateExchangeName))
}
