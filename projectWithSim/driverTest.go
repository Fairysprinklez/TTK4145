package main

import(
	"./driver"
	 "./config"

)

func main() {
	driver.Init(config.ET_Simulation) //change to ET_Simulation to run on Ander's sim 

	driver.SetMotorDirection(config.MD_Up)
	
	for {
		if driver.GetFloorSensorSignal() == 3 {
			driver.SetMotorDirection(config.MD_Down)
		} else if driver.GetFloorSensorSignal() == 0 {
			driver.SetMotorDirection(config.MD_Up)
		}

		if driver.GetStopSignal() == 1 {
			driver.SetMotorDirection(config.MD_Stop)
		}		
	}

	
	
	

}
