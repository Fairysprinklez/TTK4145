package driver

import "time"
import . "../config"

/*
#cgo CFLAGS: -std=gnu11
#cgo LDFLAGS: -lpthread -lcomedi -lm
#include "elev.h"
*/
import "C"

const pollRate = 20 * time.Millisecond

func Init(elevatorType ElevatorType) {
	C.elev_init(C.elev_type(elevatorType))
}

func SetMotorDirection(dirn MotorDirection) {
	C.elev_set_motor_direction(C.elev_motor_direction_t(dirn))
}
func SetButtonLamp(button int, floor int, value int) {
	C.elev_set_button_lamp(C.elev_button_type_t(button), C.int(floor), C.int(value))
}
func SetFloorIndicator(floor int) {
	C.elev_set_floor_indicator(C.int(floor))
}
func SetDoorOpenLamp(value int) {
	C.elev_set_door_open_lamp(C.int(value))
}
func SetStopLamp(value int) {
	C.elev_set_stop_lamp(C.int(value))
}

func GetButtonSignal(button int, floor int) int {
	return int(C.elev_get_button_signal(C.elev_button_type_t(button), C.int(floor)))
}
func GetFloorSensorSignal() int {
	return int(C.elev_get_floor_sensor_signal())
}
func GetStopSignal() int {
	return int(C.elev_get_stop_signal())
}
func GetObstructionSignal() int {
	return int(C.elev_get_obstruction_signal())
}

func SetAllButtonLamps(lightMatrix [config.NumFloors][config.NumButtons]bool){
	for f := 0; f < config.NumFloors; f++{
		for b := 0; b < config.NumButtons; b++{
			if lightMatrix[f][b] {
				SetButtonLamp(b, f, 1)
			}else{
				SetButtonLamp(b, f, 0)
			}
		}
	}
}

func PollButtons(receiver chan<- ButtonEvent) {

	previous := make([][3]int, NumFloors)

	for {
		time.Sleep(pollRate)
		for f := 0; f < NumFloors; f++ {
			for b := 0; b < 3; b++ {
				v := GetButtonSignal(b, f)
				if v != previous[f][b] && v != 0 {
					receiver <- ButtonEvent{f, ButtonType(b)}
				}
				previous[f][b] = v
			}
		}

	}
}

func PollFloorSensor(receiver chan<- int) {

	previous := -1

	for {
		time.Sleep(pollRate)
		floor := GetFloorSensorSignal()
		if floor != previous && floor != -1 {
			SetFloorIndicator(floor)
			receiver <- floor
		}
		previous = floor
	}
}
