package network

import (
	"../config"
	"./bcast"
	"./localip"
	"./peers"
	"fmt"
)

func Network(messageTx chan config.Message, messageRx chan config.Message, lostPeers chan []string) {

	localIP, err := localip.LocalIP()
	if err != nil {
		fmt.Println(err)
		localIP = "DISCONNECTED"
	}
	var ID string = fmt.Sprintf(localIP)

	peerUpdateCh := make(chan peers.PeerUpdate)

	peerTxEnable := make(chan bool)
	go peers.Transmitter(20188, ID, peerTxEnable)
	go peers.Receiver(20188, peerUpdateCh)

	go bcast.Transmitter(20088, messageTx)
	go bcast.Receiver(20088, messageRx)

	for {
		p := <-peerUpdateCh
		if len(p.Lost) != 0 {
			lostPeers <- p.Lost
		}
	}
}
