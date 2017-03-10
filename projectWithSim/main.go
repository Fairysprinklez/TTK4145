package main

import(
	"./network"
	"./config"
	"fmt"
	"time"
	
)

func main() {
	send := make(chan config.Message)
	recieve := make(chan config.Message)
	lostPeers := make(chan []string)

	go network.Network(send, recieve, lostPeers)
	
	go func() {	
		test := config.Message{"Kåre", 0}
		for {
			send <- test
			test.Iter++
			time.Sleep(1*time.Second)
			fmt.Println("sending")
		}
	}()
	
	for {
		select {
		case p := <-lostPeers:
			fmt.Println(p)
		case r := <-recieve:
			fmt.Println("recieved: ", r)
		}
	}	
}
