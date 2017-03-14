//TODO: package

import (
	"./config"
)

/*TODO: determine channels:
	
	
*/

var nodeMap config.NodeMap
var myID string
var change bool

func nodeMapCompiler(	recieveMessage <-chan config.Message,
						sendMap chan<- config.NodeMap,
						recieveLift <-chan config.Lift,
						lostPeers <-chan []string,
						){
	
	thisLift := <- recieveLift
	myID = ThisLift.ID
	myMap = make(config.NodeMap)
	myMap[myID] = thisLift
	change = false
	
	for{	
		select{
		case message := <-recieveMessage:
			senderID := message.ID
			//Adds lift in map, if it doesn't exist in it
			for key := range message.nodeMap {
				exists := myMap[key]
				if !exists {
					myMap[key] := message.nodeMap[key]
					change = true
				}
			}
			
			//Changes myMap[senderID] to message.nodemap[senderID] if they are different
			if (message.nodeMap[senderID] != myMap[senderID]){
				myMap[senderID] := message.nodeMap[senderID]
				change = true
			}
			
			//Adopts appropriate values of message.nodeMap[senderID] into myMap[myID]
			for floors := 0; floors < config.NumFloors; floors++ {
				for buttons := 0; buttons < 2; buttons++{
					if (message.nodeMap[senderID].Requests[floors][buttons] == false && message.nodeMap[senderID].Behaviour == config.DoorOpen){
					
					}
				}
			}
			
			
			
			
				
				if {
				for /*elements in map*/{
					/*compare values*/
					/*What conditions to set? 0->1?, 1->0?, overwrite every time? Need some sort of sync*/
					if /*senderID = mapElement.key*/{
						/*update values*/
					}
				}
			//TODO: Figure out how the hell to manage the maps....
			//testing yielded no conclusive results, NEED HELP!!!
			
			/*sendchannel*/<-nodeMap
	
		case incLift = <-recieveLift:
			thisLift = myMap[myID]
			if thisLift != incLift{
				thisLift = incLift
				myMap[myID] = thisLift
				sendMap <- myMap
			}
			
		}
	}