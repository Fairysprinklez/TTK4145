package main

import (
	
	"./localip"
	"./peers"
	"./bcast"
	"flag"
	"fmt"
	"time"
	"os"
	"../config"
)




func network() {
	
	var universe map[string]config.Lift

	// Our id can be anything. Here we pass it on the command line, using
	//  `go run main.go -id=our_id`
	var id string
	flag.StringVar(&id, "identity", "", "id of this peer")
	flag.Parse()
	

	if id == "" {
		localIP, err := localip.LocalIP()
		if err != nil {
			fmt.Println(err)
			localIP = "DISCONNECTED"
		}
		id = fmt.Sprintf("%s-%d", localIP, os.Getpid())
		fmt.Printf("%s", id)
		
		
	}	
	
	
	// We make a channel for receiving updates on the id's of the peers that are
	//  alive on the network
	peerUpdateCh := make(chan peers.PeerUpdate)
	// We can disable/enable the transmitter after it has been started.
	// This could be used to signal that we are somehow "unavailable".
	peerTxEnable := make(chan bool)
	go peers.Transmitter(20188, id, peerTxEnable)
	go peers.Receiver(20188, peerUpdateCh)

	universe = make(map[string]config.Lift)
	messageTx := make(chan map[string]config.Lift)
	messageRx := make(chan map[string]config.Lift)

	go bcast.Transmitter(20088, messageTx)
	go bcast.Receiver(20088, messageRx)

	//testing
	universe[id] = config.Lift{id, 2, 3}
	
	go func() {
		for {
			messageTx <- universe
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select {
		
		case p := <-peerUpdateCh:
			fmt.Printf("Peer update:\n")
			fmt.Printf("  Peers:    %q\n", p.Peers)
			fmt.Printf("  New:      %q\n", p.New)
			fmt.Printf("  Lost:     %q\n", p.Lost)
		
		case a := <- messageRx:
			for k := range a {
				//used for testing
				fmt.Printf("Received from %s: %s %d %d\n", k, a[k].ID, a[k].LastKnownFloor, a[k].MotorDir)
			}			
		}
	}
}


//for testing
func main() {
	
	network()
}
