package main

import (
	"./network"
	"flag"
	"fmt"
	"os"
	"time"
)

type Message struct {
	Msg string
	Iter int
}


func main() {


	messageTx := make(chan Message)
	messageRx := make(chan Message)
	// Our id can be anything. Here we pass it on the command line, using
	//  `go run main.go -id=our_id`
	var identity string
	flag.StringVar(&identity, "identity", "", "id of this peer")
	flag.Parse()


	go network.ConnectToNetwork(messageTx, messageRx)

	go func() {
		testMsg := Message{"Hello from " + identity, 0}
		for {
			testMsg.Iter++
			go network.BroadcastMessage(testMsg, messageTx)
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		a := network.MessageRecieved(messageRx)
			fmt.Printf("Received: %#v\n", a)
	}
	
}
