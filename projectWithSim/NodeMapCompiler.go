//TODO: package

import (
	"./config"
)

/*TODO: determine channels:
	
	
*/

var nodeMap config.NodeMap
var myID string
var myMapChanged bool
var thisLift config.Lift
var senderLift config.Lift

func nodeMapCompiler(	recieveMessage <-chan config.Message,
						sendMap chan<- config.NodeMap,
						recieveLift <-chan config.Lift,
						lostPeers <-chan []string,
						){
	
	myMap = make(config.NodeMap)

	thisLift := <- recieveLift
	myID = ThisLift.ID
	myMap[myID] = thisLift
	myMapChanged = false
	
	for{	

		myMapChanged = false

		select{
		case message := <-recieveMessage:
			senderID := message.ID
			senderLift := message.nodeMap[senderID]
			thisLift := myMap[myID]

			//Adds lift in map, if it doesn't exist in it
			for key := range message.nodeMap {
				exists := myMap[key]
				if !exists {
					newLift := message.nodeMap[key]
					myMap[key] := newLift
					myMapChanged = true
				}
			}
			
			//Changes myMap[senderID] to message.nodemap[senderID] if they are different
			if (senderLift != myMap[senderID]){
				myMap[senderID] := senderLift
				myMapChanged = true
			}
			
			//Adopts appropriate values of message.nodeMap[senderID] into myMap[myID]
			for floor := 0; floors < config.NumFloors; floor++ {
				for button := 0; button < 2; button++{
					thisLift = myMap[myID]
					if (senderLift.Requests[floor][button] != thisLift.Requests[floor][button]){
						if (!senderLift.Requests[floor][button] 
							&& senderLift.Behaviour == config.LiftDoorOpen
							&& senderLift.LastKnownFloor == floor)
						{

							thisLift.Requests[floors][button] = false

						}else if (senderLift.Requests[floor][button]
							&& !(thisLift.Behaviour == config.LiftDoorOpen && thisLift.LastKnownFloor == floor))
						{

							thisLift.Requests[floor][button] = true
							
						}
					}
					if thisLift != myMap[myID] {

					myMap[myID] = thisLift
					myMapChanged = true

					}
				}
			}
			
			if myMapChanged {
				sendMap <- myMap
			}
	
		case incLift = <-recieveLift:
			thisLift = myMap[myID]
			if thisLift != incLift{

				thisLift.Behaviour = incLift.Behaviour
				thisLift.MotorDir = incLift.MotorDir
				thisLift.TargetFloor = incLift.TargetFloor
				if thisLift.Requests != incLift.Requests {
					for floor := 0; floors < config.NumFloors; floor++ {
						for button := 0; button < 2; button++{
							if thisLift.Requests[floor][button] != incLift.Requests[floor][button] {
								
							}	
						}
					}
				}
			}
		}
	}