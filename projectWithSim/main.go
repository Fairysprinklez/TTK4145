package main

import(
	"./network"
	"./config"
	"fmt"
	"time"
	
)
/*
func initializeLift() config.Lift {
	//TODO: this must be fixed
	lift := config.Lift{localip.Localip,
		1,
		-1,
		config.MD_Stop,
		LiftIDle,

*/

func main() {

	//we need to initialize an instance of elevator here I think -Martin

	
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
