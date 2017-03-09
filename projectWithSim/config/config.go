package config

//TODO: correct imports, move numFloors, numButtons etc.


const NumFloors int = 4
const NumButtons int = 3
const MotorSpeed int = 2800



type Lift struct{
	ID string
	Alive bool
	LastKnownFloor int
	MotorDir int
	//hallMatrix [4][2]bool
	//cabMatrix [4]bool
}


type NodeMap map[string]Lift

type Message struct {
        //NodeMap NodeMap
        ID string
	Iter int
}
	

	


