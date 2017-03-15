package network

import (
	"../config"
	"./bcast"
	"./localip"
	"./peers"
	"fmt"
	"os"
)

func Network(messageTx chan config.Message, messageRx chan config.Message, lostPeers chan []string) {

	localIP, err := localip.LocalIP()
	if err != nil {
		fmt.Println(err)
		localIP = "DISCONNECTED"
	}
	var ID string = fmt.Sprintf("peer-%s-%d", localIP, os.Getpid())

	peerUpdateCh := make(chan peers.PeerUpdate)

	peerTxEnable := make(chan bool)
	go peers.Transmitter(20188, ID, peerTxEnable)
	go peers.Receiver(20188, peerUpdateCh)

	go bcast.Transmitter(20088, messageTx)
	go bcast.Receiver(20088, messageRx)
	fmt.Println("started network")

	for {
		p := <-peerUpdateCh
		if len(p.Lost) != 0 {
			lostPeers <- p.Lost
		}
	}
}
