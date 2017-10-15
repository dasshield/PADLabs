package main

import (
	"net"
	"fmt"
	"bufio"
	"encoding/gob"
	"PAD1/vars"
	"sync"
	"errors"
)

type MQueue struct {
	lock sync.Mutex
	stack []string
}


func NewMQueue() *MQueue {
	return &MQueue {stack: []string{}}
}

func (q *MQueue) Push(s string) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.stack = append(q.stack, s)
}

func (q *MQueue) Pop() (string, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	length := len(q.stack)
	if length == 0 {
		return "", errors.New("Empte queue")
	}

	res := q.stack[length - 1]
	q.stack = q.stack[:length - 1]
	return res, nil
}

var Queue = NewMQueue()


func main()  {
	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":14460")

	fmt.Println("Listening on port 14460...")
	// close the listener when the app closes

	for {
		// accept connection on port
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
		}

		go handleConnection(conn)
	}

	defer ln.Close()
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	message := readMessage(*rw)
	handleMessage(message, *rw)
/*	for {
		message := readMessage(rw)
		fmt.Println(message)

	}*/
}

func readMessage(rw bufio.ReadWriter) (message vars.ServerMessage) {
	dec := gob.NewDecoder(rw)
	dec.Decode(&message)
	return message
}

func handleMessage(message vars.ServerMessage, rw bufio.ReadWriter)  {
	switch message.Type {
	case "client":
		Queue.Push(message.Message)
	case "receiver":
		mes, err := Queue.Pop()
		if err != nil {
			fmt.Println(err)
			break
		}
		rw.WriteString(mes)
	default:
		fmt.Println("Invalid mesage type")
	}
}
