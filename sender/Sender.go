package main

import (
	"net"
	"encoding/gob"
	"fmt"
	"PAD1/vars"
)


func main() {
	conn, _ := net.Dial("tcp", "0.0.0.0:14460")

	msg := &vars.ServerMessage{
		Type: "client",
		Message: "simple",
	}

	fmt.Println(msg)
	enc := gob.NewEncoder(conn)
	enc.Encode(msg)

}
