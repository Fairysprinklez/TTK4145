package nodeMapCompiler

import (
	"../config"
)


var myMap config.NodeMap
var myID string
var myMapChanged bool
var thisLift config.Lift
var senderLift config.Lift
var liftUpdate config.LiftUpdate

func NodeMapCompiler(	recieveMessage <-chan config.Message,
						sendMap chan<- config.NodeMap,
						recieveLift <-chan config.LiftUpdate,
						lostPeers <-chan []string,
						){
	
	myMap = make(config.NodeMap)

	liftUpdate := <- recieveLift
	thisLift := liftUpdate.Lift
	myID = thisLift.ID
	myMap[myID] = thisLift
	sendMap <- myMap
	
	for{	

		myMapChanged = false

		select{
			case message := <-recieveMessage:
				senderID := message.ID
				senderLift := message.NodeMap[senderID]
				thisLift := myMap[myID]

				//Adds lift in map, if it doesn't exist in it
				for key := range message.NodeMap {
					 
					if _, exists := message.NodeMap[key]; !exists{
						newLift := message.NodeMap[key]
						myMap[key] = newLift
						myMapChanged = true
					}
				}
			
				if (senderLift.Alive && !(myMap[senderID].Alive)){
				
					for floor := 0; floor < config.NumFloors; floor++ {
						for button := 0; button < config.NumButtons; button++{
							val := (senderLift.Requests[floor][button] || myMap[senderID].Requests[floor][button])
			
							senderLift.Requests[floor][button] = val
			
						}
					}
					myMap[senderID] = senderLift
					myMapChanged = true

				}else if (thisLift.Alive && !(message.NodeMap[myID].Alive)){
				
					for floor := 0; floor < config.NumFloors; floor++ {
						for button := 0; button < config.NumButtons; button++{
							
							val := (thisLift.Requests[floor][button] || message.NodeMap[myID].Requests[floor][button])
							thisLift.Requests[floor][button] = val
			
						}
					}
					myMap[myID] = thisLift
					myMapChanged = true
				}
			
				//Adopts appropriate values of message.nodeMap[senderID] into myMap[myID]
				for floor := 0; floor < config.NumFloors; floor++ {
					for button := 0; button < 2; button++{
						thisLift = myMap[myID]
					
						if (senderLift.Requests[floor][button] != thisLift.Requests[floor][button]){
							if ((!senderLift.Requests[floor][button]) && (senderLift.Behaviour == config.LiftDoorOpen) && (senderLift.LastKnownFloor == floor)){
									thisLift.Requests[floor][button] = false

							}else if (senderLift.Requests[floor][button] && !(thisLift.Behaviour == config.LiftDoorOpen && thisLift.LastKnownFloor == floor)){

								thisLift.Requests[floor][button] = true
							
							}
						}
					}
					if thisLift != myMap[myID] {

					myMap[myID] = thisLift
					myMapChanged = true

					}
				}
						
				//Changes myMap[senderID] to message.nodemap[senderID] if they are different
				if (senderLift != myMap[senderID]){
					myMap[senderID] = senderLift
					myMapChanged = true
				}
	
			case liftUpdate = <-recieveLift:
				incomingLift := liftUpdate.Lift
				source := liftUpdate.Source
				thisLift = myMap[myID]
				if thisLift != incomingLift{
					if source == config.FSM{
				
						thisLift.Behaviour = incomingLift.Behaviour
						thisLift.MotorDir = incomingLift.MotorDir
						for floor := 0; floor < config.NumFloors; floor++ {
							for button := 0; button < config.NumButtons; button++ {
								if (!incomingLift.Requests[floor][button] && incomingLift.Behaviour == config.LiftDoorOpen && incomingLift.LastKnownFloor == floor) {

									thisLift.Requests[floor][button] = incomingLift.Requests[floor][button]

								}
							}
						}
					} else if source == config.Cost {

						thisLift.TargetFloor = incomingLift.TargetFloor

					}else if source == config.Button_Poll {
						
						for floor := 0; floor < config.NumFloors; floor++ {
							for button := 0; button < config.NumButtons; button++{
								if incomingLift.Requests[floor][button]{
									thisLift.Requests[floor][button] = incomingLift.Requests[floor][button]
								}
							}
						}
					}else if source == config.Floor_Poll {

						thisLift.LastKnownFloor = incomingLift.LastKnownFloor

					}
				}
				if thisLift != myMap[myID] {
					
					myMap[myID] = thisLift
					myMapChanged = true
				}

			case disconnectedNodes := <- lostPeers:
				thisLift := myMap[myID]
				for _,peer := range disconnectedNodes {
					disconnectedLift := myMap[peer]
					disconnectedLift.Alive = false
					for floor := 0; floor < config.NumFloors; floor++ {
						for button := 0; button < 2; button++{
							thisLift.Requests[floor][button] = (thisLift.Requests[floor][button] || disconnectedLift.Requests[floor][button])

						}
					}
					myMap[peer] = disconnectedLift
				}
				myMap[myID] = thisLift
				myMapChanged = true



			if myMapChanged {
				sendMap <- myMap
				
			}

		}
	}
}