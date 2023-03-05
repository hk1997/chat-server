package impl

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	rabbitmq "github.com/wagslane/go-rabbitmq"
)

type RabbitMqBrokerImpl struct {
	ConnectionUri         string
	connectionEstablished bool
	publisher             *rabbitmq.Publisher
}

func (rb RabbitMqBrokerImpl) getConnection() *rabbitmq.Conn {
	conn, err := rabbitmq.NewConn(
		rb.ConnectionUri,
		rabbitmq.WithConnectionOptionsLogging,
	)

	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func (rb RabbitMqBrokerImpl) getPublisher() *rabbitmq.Publisher {

	conn := rb.getConnection()

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("events"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}

	return publisher
}

// Publishes a message to the queue.
func (rb RabbitMqBrokerImpl) Publish(routingKeys []string, message []byte) {

	if rb.publisher == nil {
		rb.publisher = rb.getPublisher()
	}

	err := rb.publisher.Publish(message, routingKeys)

	if err != nil {
		log.Println(err)
	}

	log.Println("Published message", message, "to routing keys", routingKeys)
}

// Registers a consumer to a queue with the given message handler.
func (rb RabbitMqBrokerImpl) Consumer(queue string, handler func(body []byte)) {
	conn := rb.getConnection()
	defer conn.Close()

	consumer, err := rabbitmq.NewConsumer(
		conn,
		func(d rabbitmq.Delivery) rabbitmq.Action {
			handler(d.Body)
			// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue
			return rabbitmq.Ack
		},
		queue,
		rabbitmq.WithConsumerOptionsRoutingKey("my_routing_key"),
		rabbitmq.WithConsumerOptionsExchangeName("events"),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)

	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	// block main thread - wait for shutdown signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("stopping consumer")
}
