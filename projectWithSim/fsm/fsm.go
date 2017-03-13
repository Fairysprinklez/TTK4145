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

func FsmOnInitBetweenFloors(LiftCh chan config.Lift) {
	thisLift := <- LiftCh	
	driver.SetMotorDirection(config.MD_Down)
	thisLift.MotorDir = config.MD_Down
	thisLift.LiftBehavior = config.LiftMoving
	LiftCh <- thisLift
	
}
 
func FsmOnRequestButtonPress(liftIn chan config.Lift, liftOut chan config.Lift, buttonEvent config.ButtonEvent) {
	timerIsActive := false
	timer := time.NewTimer(3*time.Second)	
	for {
	
		thisLift := <- liftin

		switch thisLift.Behaviour {
		case LiftDoorOpen:
			if thisLift.LastKnownFloor == thisLift.TargetFloor {
				driver.SetDoorOpenLamp(1)	
				if !timerExist {
					timer.Reset(time.Second * config.DoorOpenDuration)
					timerIsActive = true
					for b := 0; b < config.NumButtons; b++ {
						thisLift.Requests[thislift.LastKnownFloor][b] = false
					}
			}
			select {
			case <-timer.C: {
				driver.SetDoorOpenLamp(0)
				thisLift.Behaviour = config.LiftIdle
				timerIsActive = false				
				}
			default:
			}

		case LiftMoving:
			if thisLift.LastKnownFloor == thisLift.TargetFloor {
				thisLift.Behaviour = config.LiftDoorOpen
			}
	
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
		liftOut <- thisLift
	}

}

func FsmOnFloorArrival(newFloor int) {

}

func FsmOnDoorTimeout() {

}


