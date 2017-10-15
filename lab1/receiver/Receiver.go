package main

import (
	"net"
	"bufio"
	"fmt"
	"PAD1/vars"
	"encoding/gob"
)

func main() {
	conn, _ := net.Dial("tcp", "0.0.0.0:14460")

	msg := &vars.ServerMessage{
		Type: "client",
	}

	enc := gob.NewEncoder(conn)
	enc.Encode(msg)

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message != "" {
			fmt.Println("Message from brocker:", message)
			break
		}

	}
}
