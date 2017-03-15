// We often want to execute Go code at some point in the
// future, or repeatedly at some interval. Go's built-in
// _timer_ and _ticker_ features make both of these tasks
// easy. We'll look first at timers and then
// at [tickers](tickers).

package main

import "time"
import "fmt"

func main() {
	timerExist := false
	timer := time.NewTimer(3*time.Second)	
	for {		
		
		if !timerExist {
			timer.Reset(3*time.Second)
			timerExist = true
		}
		select {
		case <-timer.C:
			fmt.Println("we waited 3 seconds")
			timerExist = false
		default:
		}
		fmt.Println("SPAM")
		time.Sleep(200*time.Millisecond)
	}


}
