package broker

type Broker interface {
	// Publishes a message to the queue.
	Publish(routingKeys []string, message []byte)
	// Registers a consumer to a queue with the given message handler.
	Consumer(queue string, handler func(body []byte))
}
