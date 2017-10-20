package main

import (
	"pad/PADLabs/lab1/common"
	"bufio"
	"os"
	"fmt"
	"time"
)

func main()  {
	sub := common.NewSub()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input topic name: ")

	topic, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Cannot read inputed topic")
	}

	go sub.Subscribe(topic, func(msg string) {
		fmt.Println("Received message: ", msg)
	})

	time.Sleep(time.Second * 10000000)
}
