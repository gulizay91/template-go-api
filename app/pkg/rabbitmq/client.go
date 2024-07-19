package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	stdLog "log"
)

type ExchangeType string

const (
	Direct ExchangeType = "direct"
	Fanout ExchangeType = "fanout"
	Topic  ExchangeType = "topic"
)

const DeadLetterQueueSuffix = "_error"

func ConnectRabbitMQ(hosts []string, username string, password string) *amqp.Connection {
	var err error
	var conn *amqp.Connection
	for _, host := range hosts {
		rabbitmqURL := fmt.Sprintf("amqp://%s:%s@%s/",
			username,
			password,
			host,
		)
		conn, err = amqp.Dial(rabbitmqURL)
		if err == nil {
			stdLog.Printf("Connected to RabbitMQ at %s", host)
			break
		}
		stdLog.Printf("Failed to connect to RabbitMQ at %s: %v", host, err)
	}
	if err != nil {
		stdLog.Fatalf("Failed to connect to any RabbitMQ server: %v", err)
	}
	return conn
}

func CreateChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		stdLog.Fatalf("Failed to open a channel: %v", err)
	}
	return ch
}

func DeclareExchange(ch *amqp.Channel, name, kind string) {
	err := ch.ExchangeDeclare(
		name,  // name of the exchange
		kind,  // type
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		stdLog.Printf("Failed to declare exchange %s: %v", name, err)
	}
}

func DeclareQueue(ch *amqp.Channel, name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table, hasDlq bool) {
	if hasDlq {
		if args == nil {
			args = amqp.Table{}
		}
		args["x-dead-letter-exchange"] = ""                              // default exchange
		args["x-dead-letter-routing-key"] = name + DeadLetterQueueSuffix // dead-letter queue name
	}
	_, err := ch.QueueDeclare(
		name,       // name
		durable,    // durable
		autoDelete, // delete when unused
		exclusive,  // exclusive
		noWait,     // no-wait
		args,       // arguments
	)
	if err != nil {
		stdLog.Printf("Failed to declare queue %s: %v", name, err)
	}
}

func BindQueue(ch *amqp.Channel, queueName, routingKey, exchangeName string) {
	err := ch.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		stdLog.Printf("Failed to bind queue %s to exchange %s: %v", queueName, exchangeName, err)
	}
}

func DeleteQueueIfExistsAndEmpty(ch *amqp.Channel, queueName string) {
	queue, err := ch.QueueInspect(queueName)
	if err != nil {
		stdLog.Printf("Queue %s does not exist or cannot be inspected: %v", queueName, err)
		return
	}

	if queue.Messages == 0 {
		_, err := ch.QueueDelete(queueName, false, false, false)
		if err != nil {
			stdLog.Printf("Failed to delete queue %s: %v", queueName, err)
		} else {
			stdLog.Printf("Queue %s deleted", queueName)
		}
	} else {
		stdLog.Printf("Queue %s is not empty and cannot be deleted", queueName)
	}
}
