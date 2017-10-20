package common

import (
	"bufio"
	"net"
)

type publisherApi interface {
	Publish(topic string, message string) error
}

type publisher struct {
	id int
	rw *bufio.ReadWriter
}

func newRW() (*bufio.ReadWriter, error) {

	conn, err := net.Dial("tcp", "0.0.0.0:14460")
	if err != nil {
		return nil, err
	}
	rw := bufio.NewReadWriter(bufio.NewReader(conn),
		bufio.NewWriter(conn))
	return rw, nil
}

func NewPub() publisherApi {
	rw, _ := newRW()
	publisherObj := publisher{
		publisherId,
		rw,
	}

	publisherId++
	return publisherObj
}

func (publisherObj publisher) Publish(topic string, message string) error {
	err := sendMessage(publisherObj.rw, ServerMessage{
		Id:      publisherObj.id,
		Type:    messagePublished,
		Topic:   topic,
		Message: message,
	})
	return err
}
