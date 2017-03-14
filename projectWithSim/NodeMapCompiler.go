//TODO: package

import (
	"./config"
)


var nodeMap config.NodeMap
var myID string
var myMapChanged bool
var thisLift config.Lift
var senderLift config.Lift

//TODO: return type?
func mergeRequests(recievedRequests []bool, myCopy []bool){
	mergedRequests := myCopy
	for floor := 0; floors < config.NumFloors; floor++ {
		for button := 0; button < config.NumButtons; button++{
			
			mergedRequests[floor][button] = (recievedRequests[floor][button] || myRequests[floor][button])
			
		}
	}
	
	return mergedRequests
}

func nodeMapCompiler(	recieveMessage <-chan config.Message,
						sendMap chan<- config.NodeMap,
						recieveLift <-chan config.LiftUpdate,
						lostPeers <-chan []string,
						){
	
	myMap = make(config.NodeMap)

	thisLift := <- recieveLift
	myID = thisLift.ID
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
			
			if (senderLift.Alive && !(myMap[senderID].Alive)){
				mergedRequests := mergeRequests(senderLift.Requests, myMap[senderID].Requests)
				senderLift.Requests = mergedRequests
				myMap[senderID] = senderLift
				myMapChanged = true

			}else if (thisLift.Alive && !(message.nodeMap[myID])){
				mergedRequests := mergeRequests(thisLift.Requests, message.nodemap[myID].Requests)
				thisLift.Requests = mergedRequests
				myMap[myID] = thisLift
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
					}
					if thisLift != myMap[myID] {

					myMap[myID] = thisLift
					myMapChanged = true

					}
				}
			}
			
			//Changes myMap[senderID] to message.nodemap[senderID] if they are different
			if (senderLift != myMap[senderID]){
				myMap[senderID] := senderLift
				myMapChanged = true
			}
	
		case liftUpdate = <-recieveLift:
			incomingLift = liftUpdate.Lift
			source = liftUpdate.Source
			thisLift = myMap[myID]
			if thisLift != incLift{
				if source == config.FSM{
					thisLift.Behaviour = incomingLift.Behaviour
					for floor := 0; floors < config.NumFloors; floor++ {
						for button := 0; button < config.NumButtons; button++{
							if (!incomingLift.Requests[floor][Button]
								&& incomingLift.Behaviour == config.LiftDoorOpen
								&& incomingLift.LastKnownFloor == floor)
							{

								thisLift.Requests[floor][button] = incomingLift.Requests[floor][Button]

							}
						}
					}
				} else if source == config.Cost {

					thisLift.TargetFloor = incomingLift.TargetFloor

				}else if source == config.Button_Poll {
					for floor := 0; floors < config.NumFloors; floor++ {
						for button := 0; button < config.NumButtons; button++{
							if incomingLift.Requests[floor][Button]{
								thisLift.Requests[floor][button] = incomingLift.Requests[floor][Button]
							}
						}
					}
				}else if source == config.Floor_Poll {

					thisLift.LastKnownFloor = incomingLift.LastKnownFloor

				}
			}
			if thisLift != myMap[myId] {
				myMap[myID] = thisLift
				myMapChanged = true
			}

		case disconnectedNodes := <- lostPeers:
			thisLift := myMap[myID]
			for _,peer := range disconnectedNodes {
  				disconnectedLift := myMap[peer]
  				disconnectedLift.Alive = false
  				for floor := 0; floors < config.NumFloors; floor++ {
					for button := 0; button < 2; button++{
						thisLift.Requests[floor][button] = (thisLift.Requests[floor][button] || disconnectedLift.Requests[floor][button])

					}
				}
				myMap[peer] = disconnectedLift
			}
			myMap[myID]
			myMapChanged = true



		if myMapChanged {
			sendMap <- myMap
		}

	}
}