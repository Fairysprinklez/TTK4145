//TODO: package?

import . "time"

func timer(timerChan chan <- bool){
	timer := Newtimer(Second * doorOpenDuration) 
	//TODO: define doorOpenDuration somewhere
	<- timer.C
	timerChan <- true
}
	