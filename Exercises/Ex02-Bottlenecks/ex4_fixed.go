package main
//found semaphore solution here:http://www.golangpatterns.info/concurrency/semaphores
import (
	. "fmt"
	"runtime"
	"time"
)

var Global = 0

type empty struct{}
type semaphore chan empty

// acquire n resources
func (s semaphore) P(n int) {
    e := empty{}
    for i := 0; i < n; i++ {
        s <- e
    }
}

// release n resources
func (s semaphore) V(n int) {
    for i := 0; i < n; i++ {
        <-s
    }
}

func routine1(s semaphore) {
	for i := 0; i < 1000000; i++ {
		s.P(1)
		Global++
		//Println(Global)
		s.V(1)
	}

}

func routine2(s semaphore) {
	for i := 0; i < 1000001; i++ {
		s.P(1)
		Global--
		//Println(Global)
		s.V(1)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	sem := make(semaphore, 1)
	go routine1(sem)
	go routine2(sem)

	time.Sleep(2000*time.Millisecond)
	Println(Global)
}
