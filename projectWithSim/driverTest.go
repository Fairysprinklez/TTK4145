package main

import(
	"./driver"
	. "./config"

)

func main() {
	driver.Init(ET_Comedi) //change to ET_Simulation to run on Ander's sim 

	driver.SetMotorDirection(MD_Up)
	
	for {
		if driver.GetFloorSensorSignal() == 3 {
			driver.SetMotorDirection(MD_Down)
		} else if driver.GetFloorSensorSignal() == 0 {
			driver.SetMotorDirection(MD_Up)
		}

		if driver.GetStopSignal() == 1 {
			driver.SetMotorDirection(MD_Stop)
		}		
	}

	
	
	

}
