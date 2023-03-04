package broker

type Message struct {
	id         int64
	senderId   int64
	receiverId int64
	timeStamp  int64
	payload    []byte
}

type Acknowledgement struct {
	id        int64
	messageId int64
	status    Status
	timestamp int64
}

type Status int8

const (
	DELIVERED Status = 0
	READ      Status = 1
)
