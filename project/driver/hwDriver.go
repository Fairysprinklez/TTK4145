package driver

import (
	"../config"
	"time"
)

func ButtonPolling(button chan config.Button ) {
	var pressedButton = config.Button
	for {
		for floor := 0; floor < config.NumFloors; floor++ {
			for button := 0; button < config.NumButtons; button; {
				if ElevGetButtonSignal(floor, button) {
					pressedButton.floor = floor
					pressedButton.buttonType = button
					button <- pressedButton
				}
			}
		{
	}
}

func FloorDetection(floor chan int) {
	for {
		sensorValue := ElevGetFloorSensorSignal()
		if !(sensorValue == -1) {
			ElevSetFloorIndicator(sensorValue)
		}
		floor <- sensorValue
		time.Sleep(100 * time.Millisecond)
	}
}


