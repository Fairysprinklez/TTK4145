package fsm

import(
	"../config"
	"../driver"
)

func setAllLights() {
	//no clue how to do this yet.
}

func FsmInit(lift config.Lift) {
	ThisLift = 
}

func FsmOnInitBetweenFloors() {
	driver.SetMotorDirection(config.MD_Down)
	config.ThisLift.MotorDir = config.MD_Down
	config.Lift.LiftBehavior = config.LiftMoving 
	
}

func FsmOnRequestButtonPress(buttonEvent config.ButtonEvent) {
	switch config.Lift.LiftBehaviour {
	case LiftDoorOpen:
		//if currentFloor == requestedFloor { 
			//start timer
		//}else { 
			//requests[buttonEvent] = 1
		//}
	case LiftMoving:
		//requests[buttonEvent] = 1

	case LiftIdle:
		//if currentfloor = requestedfloor { 
			//driver.SetDoorOpenLamp(1)
			//start timer
			//config.Lift.Liftbehaviour = LiftDoorOpen
		//}else {
			//requests[buttonEvent] = 1
			//config.ThisLift.MotorDir = requestsChooseDirection
			//driver.SetMotorDirection(config.ThisLift.MotorDir)
			//config.Lift.Liftbehaviour = LiftMoving
		//}
	}
	//setAllLights() 

}

func FsmOnFloorArrival(newFloor int) {

}

func FsmOnDoorTimeout() {

}


