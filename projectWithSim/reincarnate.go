package main

import (
	"fmt"
	"time"
	"os"
	"syscall"
	"os/exec"
	"os/signal"
)

func main() {
	var iter int
	iter = 0

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func (){
		<-sigs
		fmt.Println("I'm dying, will reincarnate")
		backup := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run martinTest.go")
		backup.Run()
		os.Exit(0)

	}()

	for {
		iter++
		fmt.Println(iter)
		time.Sleep(3*time.Second)
	}
}

