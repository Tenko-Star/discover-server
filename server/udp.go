package server

import (
	"go-discover-server/core"
	"go-discover-server/log"
	"go-discover-server/message"
	"net"
)

func RunUDPServer(bind net.IP, port int) {
	var listen, err = net.ListenUDP("udp", &net.UDPAddr{
		IP:   bind,
		Port: port,
	})

	if err != nil {
		panic("could not listen udp port: " + err.Error())
	}

	go func() {
		for {
			var data = make([]byte, 1024)
			var addr *net.UDPAddr
			var readLen = 0
			readLen, addr, err = listen.ReadFromUDP(data)
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
	}()
}
