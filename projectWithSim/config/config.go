package config

//TODO: correct imports, move numFloors, numButtons etc.


const NumFloors int = 4
const NumButtons int = 3
const MotorSpeed int = 2800



type Lift struct{
	ID string
	LastKnownFloor int
	MotorDir int
	//hallMatrix [4][2]bool
	//cabMatrix [4]bool
}


	

	


