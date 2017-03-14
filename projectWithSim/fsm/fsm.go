package fsm

import (
	"../config"
	"../driver"
	"time"
)

/*func FsmInit(lift config.Lift) {
	//is this really necessary???
	//behaviour is set in main.initializeLiftData
}*/

func FsmOnInitBetweenFloors(LiftCh chan config.Lift) {
	thisLift := <-LiftCh
	driver.SetMotorDirection(config.MD_Down)
	thisLift.MotorDir = config.MD_Down
	thisLift.Behaviour = config.LiftMoving
	LiftCh <- thisLift

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
			if thisLift.LastKnownFloor == thisLift.TargetFloor {
				thisLift.Behaviour = config.LiftDoorOpen
			}
			//what if FsmOnInitBetweenFloor get's called?

		case config.LiftIdle:
			//The cost-function set's this Lift.MotorDir, we just pass it on to hardware
			//could be something wonky here if MotorDir is 0, since it will set the hardware repeatedly
			driver.SetMotorDirection(thisLift.MotorDir)
			if thisLift.MotorDir != config.MD_Stop {
				thisLift.Behaviour = config.LiftMoving
			}
		}
		liftOut <- thisLift
	}

}
