package config

//TODO: correct imports, move numFloors, numButtons etc.

const NumFloors int = 4
const NumButtons int = 3
const MotorSpeed int = 2800

const DoorOpenDuration int = 3 //seconds

type LiftBehaviour int

const (
	LiftIdle = iota
	LiftDoorOpen
	LiftMoving
)

//Added for simplified logic in compiler -Sondre
type Source int
const (
	FSM = iota
	Cost 
	Button_Poll
	Floor_Poll
)

type Lift struct {
	ID             string
	Alive          bool
	LastKnownFloor int
	TargetFloor    int
	MotorDir       MotorDirection
	Behaviour      LiftBehaviour //state
	Requests       [NumFloors][NumButtons]bool
}

type NodeMap map[string]Lift

type Message struct {
	NodeMap NodeMap
	ID      string
	//Iter    int //for testing
}

//Added for simplified logic in compiler -Sondre
type LiftUpdate struct {
	Lift Lift
	Source Source
}