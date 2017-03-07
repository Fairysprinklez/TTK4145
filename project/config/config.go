package config
//TODO: correct imports, move numFloors, numButtons etc.


type Lift struct{
	ID int
	lastKnownFloor int
	motorDir int
	hallMatrix [4][2]bool
	cabMatrix [4]bool
}

type CompleteNodeStatus struct{
	//TODO: Make map

	


