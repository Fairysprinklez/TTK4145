// We often want to execute Go code at some point in the
// future, or repeatedly at some interval. Go's built-in
// _timer_ and _ticker_ features make both of these tasks
// easy. We'll look first at timers and then
// at [tickers](tickers).

package main

import "time"
import "fmt"

func main() {

    // Timers represent a single event in the future. You
    // tell the timer how long you want to wait, and it
    // provides a channel that will be notified at that
    // time. This timer will wait 2 seconds.
    timer1 := time.NewTimer(time.Second * 2)

    // The `<-timer1.C` blocks on the timer's channel `C`
    // until it sends a value indicating that the timer
    // expired.
    <-timer1.C
    fmt.Println("Timer 1 expired")
}
