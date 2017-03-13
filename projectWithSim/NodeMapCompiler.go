//TODO: package

import (
	"./config"
	"./network/localip"
)

/*TODO: determine channels:
	
	
*/
func nodeMapCompiler(	externalMsg chan config.Message,
						sendMsg chan config.Message,
						lostPeers chan []string,
						/*intFloor*/,
						/*intBtn*/,
						/*intMD*/,
						/*intBehav*/
						){
	//Initializing the map of the universe elsewhere
	/*myID := localip.Localip()
	nodeMap := make(map[string]config.Lift)
	nodeMap[myID] = Lift{
		myID,
		true,
		-1, 
		0, 
		{{false, false, false},
		 {false, false, false},
		 {false, false, false},
		 {false, false, false}}		
	}*/
	
	for{	
		select{
		case external := <- /*recievechannel*/:
			for /*each key*/{
				if /*key not in map*/{
					/*put lift in own copy of nodeMap*/
				}
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
	
		case myFloor := <-/*intFloor*/:
			nodeMap[myID].LastKnownFloor = myFloor
			/*sendchannel*/ <-nodeMap
			
		case myButton := <- /*intBtn*/:
			nodeMap[myID].requestMatrix[myButton.floor][myButton.button] = true
			/*sendchannel*/<-nodeMap
			
		case myMD := <- /*intMD*/:
			nodeMap[myID].MotorDir = myMD
			/*sendchannel*/<-nodeMap
		
		case myBehaviour := /*intBehav*/
			nodeMap[myID].Behaviour = myBehaviour
			/*sendchannel*/
		
		}
	}