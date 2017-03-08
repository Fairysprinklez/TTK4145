package config

//TODO: correct imports, move numFloors, numButtons etc.


const NumFloors int = 4
const NumButtons int = 3
const MotorSpeed int = 2800

type Button struct{
	floor int
	buttonType int 
}

type Lift struct{
	ID int
	lastKnownFloor int
	motorDir int
	hallMatrix [4][2]bool
	cabMatrix [4]bool
}

//type CompleteNodeStatus struct{
	//TODO: Make map

	


