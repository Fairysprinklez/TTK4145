package config

//TODO: correct imports, move numFloors, numButtons etc.


const NumFloors int = 4
const NumButtons int = 3
const MotorSpeed int = 2800



type Lift struct{
	ID string
	LastKnownFloor int
	MotorDir int
	requestMatrix [NumFloors][NumButtons]bool	
}


type NodeMap map[string]Lift

type Message struct {
        NodeMap NodeMap
        ID string
}
	

	


