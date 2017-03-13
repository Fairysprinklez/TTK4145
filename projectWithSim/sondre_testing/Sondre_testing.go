package main

import (
	"../config"
	"../network/localip"
	"fmt"
)

func initializeLiftData() config.Lift {
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

//Tried this and "var nodeMap config.NodeMap"
var nodeMap map[string]config.Lift

func main() {
	
	//"panic: assignment to entry in nil map" no matter what I try....
	nodeMap["testLift"] = initializeLiftData()
	
	fmt.Println(nodeMap["testLift"].Requests[0][0])

}

/*Maps and printing can go to hell.......*/