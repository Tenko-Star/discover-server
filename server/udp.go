package server

import (
	"go-discover-server/core"
	"go-discover-server/log"
	"go-discover-server/message"
	"net"
)

func RunUDPServer(bind net.IP, port int) {
	go runUnicast(bind, port)
	//go runMulticast(port) // have some problem
}

func runUnicast(bind net.IP, port int) {
	var listenNormal, err = net.ListenUDP("udp", &net.UDPAddr{
		IP:   bind,
		Port: port,
	})
	if err != nil {
		panic("could not listen unicast: " + err.Error())
	}

	for {
		var data = make([]byte, 1024)
		var addr *net.UDPAddr
		var readLen = 0
		readLen, addr, err = listenNormal.ReadFromUDP(data)
		if err != nil {
			log.E("could not read data from udp: %s", err.Error())
		}

		if readLen == 0 {
			continue
		}

		var m *message.Message
		m, err = message.Unmarshal(data[:readLen])
		if err != nil {
			log.W("could not parse data: %s", err.Error())
			continue
		}

		_, err = core.AddDevice(m, addr)
		if err != nil {
			log.W("could not add device: %s", err.Error())
			continue
		}
	}
}

func runMulticast(port int) {
	var listenGroup, err = net.ListenMulticastUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(239, 233, 1, 1),
		Port: port,
	})
	if err != nil {
		panic("could not listen multicast: " + err.Error())
	}

	for {
		var data = make([]byte, 1024)
		var addr *net.UDPAddr
		var readLen = 0
		readLen, addr, err = listenGroup.ReadFromUDP(data)
		if err != nil {
			log.E("could not read data from udp group: %s", err.Error())
		}

		if readLen == 0 {
			continue
		}

		var m *message.Message
		m, err = message.Unmarshal(data[:readLen])
		if err != nil {
			log.W("could not parse data: %s", err.Error())
			continue
		}

		_, err = core.AddDevice(m, addr)
		if err != nil {
			log.W("could not add device: %s", err.Error())
			continue
		}
	}
}
