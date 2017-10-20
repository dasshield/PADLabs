package common

import "bufio"

type ServerMessage struct {
	Id      int
	Type    string
	Topic   string
	Message string
}

const (
	newSubscriber    = "New subscriber"
	newPublisher     = "New publisher"
	unknownTypeError = "Unknown message type"
	messagePublished = "Message published"
)

var (
	subscriptionMap map[string][]*bufio.ReadWriter = make(map[string][]*bufio.ReadWriter)
	pipeline        map[string]*MQueue = make(map[string]*MQueue)
	publisherId	int
)
