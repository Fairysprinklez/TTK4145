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
	for {
		iter++
		fmt.Println(iter)
		time.Sleep(3*time.Second)
	}

	go func (){
		sigs := make(chan os.Signal)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		fmt.Println("I'm dying, will reincarnate")
		backup := exec.Command("gnome-terminal", "-x", "sh", "-c", "./martinTest.go")
		backup.Run()
		os.Exit(0)
	}()
}

