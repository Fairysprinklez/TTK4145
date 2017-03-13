package main

import (
	"fmt"
	"./config"
	"./network/localip"
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
		-1,
		config.MD_Stop,
		config.LiftIdle,
		requests}
	
	}
	
	return lift
}

func main() {
	ThisLift := initializeLiftData()
	
	ThisLift.Behaviour = config.LiftDoorOpen
	fmt.Println(ThisLift)
}

