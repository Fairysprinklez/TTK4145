package main

import(
	"./driver"

)

func main() {
	driver.ElevInitSim()

	driver.ElevSetMotorDirection(-1)
	
	for {
		if driver.ElevGetFloorSensorSignal() == 3 {
			driver.ElevSetMotorDirection(-1)
		} else if driver.ElevGetFloorSensorSignal() == 0 {
			driver.ElevSetMotorDirection(1)
		}

		if driver.ElevGetStopSignal() {
			driver.ElevSetMotorDirection(0)
		}		
	}

	
	
	

}
