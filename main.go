package main

import (
	"go-discover-server/log"
	"go-discover-server/server"
	"net"
)

func main() {
	var udpServerIp = net.IPv4(0, 0, 0, 0)
	var udpServerPort = 9972
	var httpServerAddr = "0.0.0.0:9973"

	//log.SetLevel(log.LevelWarn)

	log.W("Server start.")

	log.I("Start udp server.")
	server.RunUDPServer(udpServerIp, udpServerPort)

	log.I("Start Http server.")
	server.RunHttpServer(httpServerAddr)
}
