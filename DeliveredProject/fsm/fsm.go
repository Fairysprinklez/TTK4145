package fsm

import (
	"../config"
	"../driver"
	"time"
)


func FsmOnInitBetweenFloors(thisLift config.Lift) (config.Lift) {
	driver.SetMotorDirection(config.MD_Down)
	thisLift.MotorDir = config.MD_Down
	thisLift.Behaviour = config.LiftMoving
	for {
		floor := driver.GetFloorSensorSignal()
		if floor != -1 {
			thisLift.LastKnownFloor = floor
			driver.SetMotorDirection(config.MD_Stop)
			thisLift.MotorDir = config.MD_Stop
			thisLift.Behaviour = config.LiftIdle
			return thisLift
		}
	}
	
}

func FsmOnInitInFloor(thisLift  config.Lift) (config.Lift) {
	floor := driver.GetFloorSensorSignal()
	thisLift.LastKnownFloor = floor
	driver.SetMotorDirection(config.MD_Stop)
	thisLift.MotorDir = config.MD_Stop
	thisLift.Behaviour = config.LiftIdle
	return thisLift
	
}

func FsmLoop(liftIn chan config.Lift, liftOut chan config.Lift) {
	timerIsActive := false
	timer := time.NewTimer(3 * time.Second)
	for {

		thisLift := <-liftIn

		switch thisLift.Behaviour {
		case config.LiftDoorOpen:
			if thisLift.LastKnownFloor == thisLift.TargetFloor {
				driver.SetDoorOpenLamp(1)
				if !timerIsActive {
					timer.Reset(time.Second * 3)
					timerIsActive = true
					for b := 0; b < config.NumButtons; b++ {
						thisLift.Requests[thisLift.LastKnownFloor][b] = false
					}
				}
			}
			select {
			case <-timer.C:
				{
					driver.SetDoorOpenLamp(0)
					thisLift.Behaviour = config.LiftIdle
					timerIsActive = false
				}
			default:
			}

		case config.LiftMoving:
			if (driver.GetFloorSensorSignal() == thisLift.TargetFloor) && (thisLift.TargetFloor != -1) {
				thisLift.Behaviour = config.LiftDoorOpen
			}

		case config.LiftIdle:

			if(thisLift.TargetFloor > thisLift.LastKnownFloor && thisLift.TargetFloor != -1) {
				thisLift.MotorDir = config.MD_Up
			}else if (thisLift.TargetFloor > thisLift.LastKnownFloor && thisLift.TargetFloor != -1){
				thisLift.MotorDir = config.MD_Down
			}else if (thisLift.TargetFloor == thisLift.LastKnownFloor){
				thisLift.Behaviour = config.LiftDoorOpen
			}			
				
			driver.SetMotorDirection(thisLift.MotorDir)
			if thisLift.MotorDir != config.MD_Stop {
				thisLift.Behaviour = config.LiftMoving
			}
		}
		liftOut <- thisLift
	}

}
