package main

import (
	"./network"
	"flag"
	"fmt"
	"time"
)




func main() {


	messageTx := make(chan network.Message)
	messageRx := make(chan network.Message)
	// Our id can be anything. Here we pass it on the command line, using
	//  `go run main.go -id=our_id`
	// need to add some automatic way here to assign id as IP
	var identity string
	flag.StringVar(&identity, "identity", "", "id of this peer")
	flag.Parse()
	

	network.ConnectToNetwork(messageTx, messageRx)

	go func() {
		testMsg := network.Message{"Hello from " + identity, 0}
		for {
			testMsg.Iter++
			messageTx <- testMsg
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		a := <- messageRx
			fmt.Printf("Received: %#v\n", a)
	}
	
}
