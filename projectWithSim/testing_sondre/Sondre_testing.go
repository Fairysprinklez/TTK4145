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
			-1,
			config.MD_Stop,
			config.LiftIdle,
			requests}

	}

	return lift
}

var nodeMap config.NodeMap

func main() {

	Lift1 := initializeLiftData()
	Lift2 := initializeLiftData()
	Lift2.LastKnownFloor = 1

	fmt.Println(Lift1 == Lift2)

	nodeMap = make(config.NodeMap)

	//"panic: assignment to entry in nil map" no matter what I try....
	nodeMap["testLift"] = initializeLiftData()

	val, ok := nodeMap["testLift1"]
	if !ok {
		fmt.Println("testLift key in map")
		fmt.Println(val.Requests[0][0])
	}

	//fmt.Println(nodeMap["testLift"].Requests[0][0])

}

/*Maps and printing can go to hell.......*/
//LOL
//http://stackoverflow.com/questions/40578646/golang-i-have-a-map-of-structs-why-cant-i-directly-modify-a-field-in-a-struct
//LOOK AT THIS SONDRE:
//https://play.golang.org/p/ecdUU30FQT
