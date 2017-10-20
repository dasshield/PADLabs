package common

import (
	"bufio"
	"encoding/gob"
	"log"
	"net"
	"errors"
	"fmt"
)

type broker interface {
	newRW() (*bufio.ReadWriter, error)
	Start() error
}

func sendMessage(rw *bufio.ReadWriter, message ServerMessage) error {
	fmt.Println("sendmessage")
	enc := gob.NewEncoder(rw)
	err := enc.Encode(&message)
	if err != nil {
		log.Println("Error while sending message.", message, err)
		return err
	}
	rw.Flush()
	return nil
}

func receiveMessage(rw *bufio.ReadWriter) (message ServerMessage, err error) {
	dec := gob.NewDecoder(rw)
	err = dec.Decode(&message)
	if err != nil {
		log.Println("Error decoding data", err)
		return message, err
	}
	rw.Flush()
	return message, nil
}

func Start() error {
	// listen on all interfaces
	ln, err := net.Listen("tcp", ":14460")

	if err != nil {
		return errors.New("Not able to accept connections.")
	}

	fmt.Println("Listening on port 14460...")
	// close the listener when the app closes
	defer ln.Close()

	go broadCaster()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection", err)
			continue
		}
		log.Println("Connection Accepted.")

		go handleConnection(conn)
	}
}

func broadCaster() {
	for {
		for _, queue := range pipeline {
			message := queue.Pop()
			if message.Type != "nil" {
				broadCastMessage(message)
			}
		}
	}
}

func broadCastMessage(message ServerMessage) (err error) {

	streams, ok := subscriptionMap[message.Topic]
	fmt.Println(subscriptionMap[message.Topic])
	log.Println(message, streams, ok)
	if !ok {
		return nil
	}

	for _, subscriber := range subscriptionMap[message.Topic] {
		fmt.Println("send message")
		sendMessage(subscriber, message)
	}
	return nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	for {
		message, err := receiveMessage(rw)
		if err != nil {
			log.Println("error in handleConnection", err)
			return
		}
		handleMessage(message, rw)
	}
}

func handleMessage(message ServerMessage, rw *bufio.ReadWriter) {
	fmt.Println("Received message type " + message.Type)
	switch message.Type {
	case newPublisher:
		fmt.Println("Publisher connected")
	case messagePublished:

		_, ok := pipeline[message.Topic]
		if !ok {
			pipeline[message.Topic] = NewQueue()
		}
		pipeline[message.Topic].Push(message)
		fmt.Println("Message published", message.Message)
	case newSubscriber:
		subscriptionMap[message.Topic] = append(subscriptionMap[message.Topic], rw)
		fmt.Println("Subscribed to " + message.Topic)
		fmt.Println(subscriptionMap[message.Topic])
		fmt.Println(subscriptionMap)
	default:
		sendMessage(rw, ServerMessage{
			Type:    unknownTypeError,
			Message: unknownTypeError,
		})
		log.Println("Unrecognised message.")
	}
}