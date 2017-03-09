//TODO: package

import (
	"./config"
	"./network/localip"
)

/*TODO: determine channels:
	
	
*/
func universeCompiler(/*channels*/){
	//Initializing the map of the universe
	myID := localip.Localip()
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
	}
	
	for{	
		select{
		case external := <- /*recievechannel*/:
			for key, 
			
			/*sendchannel*/<-nodeMap
	
		case myFloor := <-/*internalFloor*/:
			nodeMap[myID].LastKnownFloor = myFloor
			/*sendchannel*/ <-nodeMap
			
		case myButton := <- /*intBtn*/:
			nodeMap[myID].requestMatrix[myButton.floor][myButton.button] = true
			/*sendchannel*/<-nodeMap
			
		case myMD	:= <- /*intMD*/:
			nodeMap[myID].MotorDir = myMD
			/*sendchannel*/<-nodeMap
		
		}
	}