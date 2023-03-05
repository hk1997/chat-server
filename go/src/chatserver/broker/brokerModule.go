package broker

import (
	rabbitMqBrokerImpl "chatserver/broker/impl"
)

func GetMessagingBroker(connectionUri string) Broker {
	return rabbitMqBrokerImpl.RabbitMqBrokerImpl{
		ConnectionUri: connectionUri,
	}
}
