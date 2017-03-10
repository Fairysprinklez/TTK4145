package fsm

import(
	"../config"
	"../driver"
)

type LiftBehaviour int const(
	LiftIdle = iota
	LiftDoorOpen
	LifMoving	
)

func FsmOnInitBetweenFloors() {
	
}

func FsmOnRequestButtonPress(buttonEvent config.ButtonEvent) {

}

func FsmOnFloorArrival(newFloor int) {

}

func FsmOnDoorTimeout() {

}


