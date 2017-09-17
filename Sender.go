package main

import (
	"net"
	"fmt"
)

func main() {
	conn, _ := net.Dial("tcp", "0.0.0.0:14460")
	for {
		fmt.Fprintf(conn, "valet\n")
	}

}
