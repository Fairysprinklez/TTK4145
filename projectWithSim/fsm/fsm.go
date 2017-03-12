package fsm

import(
	"../config"
	"../driver"
)

//Do we really want function calls, and not prompts on channels?
//Using function calls, certainly simplifies things as it's more similar
//to the template provided in the project repo.

func setAllLights() {
	//no clue how to do this yet.
}

func FsmInit(lift config.Lift) {
	//is this really necessary???
	//behaviour is set in main.initializeLiftData
}

func FsmOnInitBetweenFloors() {
	driver.SetMotorDirection(config.MD_Down)
	// fsm won't have access to "ThisLift" unless argument in func call from main
	config.ThisLift.MotorDir = config.MD_Down
	config.Lift.LiftBehavior = config.LiftMoving 
	
}

func FsmOnRequestButtonPress(buttonEvent config.ButtonEvent) {
	switch ThisLift.Behaviour {
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


