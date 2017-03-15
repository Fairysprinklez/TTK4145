package main

import (
	"./config"
	"./network/localip"
	"./nodeMapCompiler"
	"fmt"
)

func initializeLiftData() config.Lift {
	var lift config.Lift
	id, err := localip.LocalIP()
	if err == nil {
		lift = config.Lift{
			ID: id,
			Alive: true,
			LastKnownFloor: -1,
			TargetFloor: -1,
			MotorDir: config.MD_Stop,
			Behaviour: config.LiftIdle}

	}

	return lift
}

var nodeMap config.NodeMap
var localLift config.LiftUpdate
var testLift config.Lift

func main() {

	recievedMsg := make(chan config.Message)
	sendMap := make(chan config.NodeMap)
	liftToCompiler := make(chan config.LiftUpdate)
	disconnectedNodes := make(chan []string)
	
	go nodeMapCompiler.NodeMapCompiler(recievedMsg, sendMap, liftToCompiler, disconnectedNodes)
	
	testLift = initializeLiftData()
	testLift2 := initializeLiftData()
	testLift2.ID = "hei"
	
	nodeMap = make(config.NodeMap)
	nodeMap[testLift.ID] = testLift
	nodeMap[testLift2.ID] = testLift2
	
	fmt.Println("testLift: ", nodeMap)
	localLift.Lift = testLift
	localLift.Source = config.FSM
	//fmt.Println("localLift: ", localLift)
	//liftToCompiler <- localLift
	
	

}

/*Maps and printing can go to hell.......*/
//LOL
//http://stackoverflow.com/questions/40578646/golang-i-have-a-map-of-structs-why-cant-i-directly-modify-a-field-in-a-struct
//LOOK AT THIS SONDRE:
//https://play.golang.org/p/ecdUU30FQT
