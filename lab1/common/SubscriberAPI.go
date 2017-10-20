package common

import (
	"bufio"
	"encoding/gob"
	"log"
	"fmt"
)

type subscriberApi interface {
	Subscribe(topic string, callback Func) error
}

type Func func(string)

type subscriber struct {
	rw          *bufio.ReadWriter
	callBackMap map[string]Func
}

func (sub *subscriber) Subscribe(topic string, callback Func) error {
	sub.callBackMap[topic] = callback

	message := ServerMessage{
		Type:  newSubscriber,
		Topic: topic,
	}

	enc := gob.NewEncoder(sub.rw)
	err := enc.Encode(&message)
	if err != nil {
		log.Println("Error while sending message.", message, err)
		return err
	}
	sub.rw.Flush()

	for {
		msg, err := receiveMessage(sub.rw)
		if err != nil {
			fmt.Println("Could not subscribe ")
			break
		}
		sub.callBackMap[topic](msg.Message)
	}

	return nil
}

func NewSub() subscriberApi {
	rw, _ := newRW()
	subscriberObj := &subscriber{
		rw,
		make(map[string]Func),
	}

	return subscriberObj
}

func (sub *subscriber) hasCallback(message ServerMessage) bool {
	_, ok := sub.callBackMap[message.Topic]
	return ok
}
