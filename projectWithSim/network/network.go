package network //TODO: remember to change this to network later

import (
	"./bcast"
	"./localip"
	"./peers"
	"fmt"
	"os"
)



type Message struct {
	Msg string
	Iter int
}




func ConnectToNetwork(messageTx chan Message, messageRx chan Message) {
	//our id is the local ip of the computer (and the processID is appendend for testing)
	var id string	
	localIP, err := localip.LocalIP()
	if err != nil {
		fmt.Println(err)
		localIP = "DISCONNECTED"
	}
	id = fmt.Sprintf("peer-%s-%d", localIP, os.Getpid())
	
	// We make a channel for receiving updates on the id's of the peers that are
	//  alive on the network
	peerUpdateCh := make(chan peers.PeerUpdate)
	// We can disable/enable the transmitter after it has been started.
	// This could be used to signal that we are somehow "unavailable".
	peerTxEnable := make(chan bool)
	go peers.Transmitter(20088, id, peerTxEnable)
	go peers.Receiver(20088, peerUpdateCh)

	// We make channels for sending and receiving our custom data types
	
	// ... and start the transmitter/receiver pair on port 20088
	// These functions can take any number of channels! It is also possible to
	//  start multiple transmitters/receivers on the same port.
	go bcast.Transmitter(20088, messageTx)
	go bcast.Receiver(20088, messageRx)

}
	



