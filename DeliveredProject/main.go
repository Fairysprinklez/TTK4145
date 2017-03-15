package main

import (
	"./config"
	"./driver"
	"./fsm"
	"./network"
	"./network/localip"
	"./nodeMapCompiler"
	"./targetFloorAssigner"
	"fmt"
	"time"
	"os"
	"syscall"
	"os/exec"
	"os/signal"
)

func initializeLiftData() config.Lift {
	var lift config.Lift
	id, err := localip.LocalIP()
	if err == nil {
		lift = config.Lift{
			ID: id,
			Alive: true,
			LastKnownFloor: -1,
			TargetFloor: -1,
			MotorDir: config.MD_Stop,
			Behaviour: config.LiftUninitialized}

	}

	return lift
}

var message config.Message
var localLift config.LiftUpdate
var outboundMap config.NodeMap

func main() {

	//##### Hardware Init #####
	driver.Init(config.ET_Comedi)


	//network channels
	send := make(chan config.Message)
	recieve := make(chan config.Message)
	lostPeers := make(chan []string)

	//compiler channels
	recievedMsg := make(chan config.Message)
	sendMap := make(chan config.NodeMap)
	liftToCompiler := make(chan config.LiftUpdate)
	disconnectedNodes := make(chan []string)
	

	//polling channels
	polledButton := make(chan config.ButtonEvent)
	polledFloorSensor := make(chan int)

	//FSM channels
	liftToFsm := make(chan config.Lift)
	liftFromFsm := make(chan config.Lift)

	//COST channels
	mapToTargetFloorAssigner := make(chan config.NodeMap)
	liftFromTargetFloorAssigner := make(chan config.Lift)
	finishedReadingMap := make(chan bool)


	thisLift := initializeLiftData()
	myID := thisLift.ID

	//##### FSM Init #####
	if driver.GetFloorSensorSignal() == -1 {
		thisLift = fsm.FsmOnInitBetweenFloors(thisLift)
	}else {
		thisLift = fsm.FsmOnInitInFloor(thisLift)
	}

	//Starting threads
	go network.Network(send, recieve, lostPeers)
	go nodeMapCompiler.NodeMapCompiler(recievedMsg, sendMap, liftToCompiler, disconnectedNodes)
	go driver.PollButtons(polledButton)
	go driver.PollFloorSensor(polledFloorSensor)
	go fsm.FsmLoop(liftToFsm, liftFromFsm)
	go targetFloorAssigner.TargetFloorAssigner(mapToTargetFloorAssigner, liftFromTargetFloorAssigner, finishedReadingMap, myID)


	localLift.Lift = thisLift
	localLift.Source = config.FSM

	

	liftToCompiler <- localLift

	

	go func (){
		sigs := make(chan os.Signal)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		fmt.Println("Rebooting process")
		backup := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run main.go")
		backup.Run()
		os.Exit(0)
	}()

	//Scetch of main loop
	for{
		
		select{
			case p := <-lostPeers:
				disconnectedNodes <- p 
				

			case message := <- recieve:
				if message.ID != myID {
					recievedMsg <- message
				}
				

			case outboundMap := <- sendMap:
				thisLift = outboundMap[myID]
				message.ID = myID
				message.NodeMap = outboundMap
				send <- message
				mapToTargetFloorAssigner <- outboundMap
				<- finishedReadingMap
				liftToFsm <- thisLift
				
			case button := <- polledButton:
				liftData := initializeLiftData()
				liftData.Requests[button.Floor][button.Button] = true
				localLift.Lift = liftData
				localLift.Source = config.Button_Poll
				liftToCompiler <- localLift
				

			case floor := <- polledFloorSensor:
				liftData := initializeLiftData()
				liftData.LastKnownFloor = floor
				localLift.Lift = liftData
				localLift.Source = config.Floor_Poll
				liftToCompiler <- localLift


			case liftData := <- liftFromFsm:
				localLift.Lift = liftData
				localLift.Source = config.FSM
				liftToCompiler <- localLift

			case liftData := <- liftFromTargetFloorAssigner:
				go driver.SetAllButtonLamps(liftData.Requests)
				localLift.Lift = liftData
				localLift.Source = config.Cost
				liftToCompiler <- localLift

		}
		time.Sleep(5*time.Millisecond)

	}
}
