package main

import(
	"./network"
	"./network/localip"
	"./config"
	"fmt"
	"time"
	
)

func initializeLiftData() config.Lift {
	//TODO: this is a hackjob, but could be useful...
	var lift config.Lift
	var requests [config.NumFloors][config.NumButtons]bool
	id, err := localip.LocalIP()
	if err != nil {
		for f := 0; f < config.NumFloors; f++ {
			for b := 0; b < config.NumButtons; b++ {
				requests[f][b] = false		
			}
		}
	
		lift = config.Lift{id,
		true,
		-1,
		config.MD_Stop,
		config.LiftIdle,
		requests}
	
	}
	
	return lift
}



func main() {

	//we need to initialize an instance of elevator here I think -Martin
	ThisLift := initializeLiftData()
	//fmt.Println(ThisLift)

	if driver.GetFloorSensorSignal() == -1 {
		fsm.FsmOnInitBetweenFloors()
	}

	

	
	
	send := make(chan config.Message)
	recieve := make(chan config.Message)
	lostPeers := make(chan []string)
	
	go network.Network(send, recieve, lostPeers)
	//go nodeMapCompiler(send, recieve, lostPeers, /*intFloor*/, /*intBtn*/, /*intMD*/, /*intBehav*/)
	
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
