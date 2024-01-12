package main

import (
	"github.com/google/uuid"
	"go-discover-server/message"
	"net"
	"testing"
	"time"
)

func TestUDP(t *testing.T) {
	var addr = net.IPv4(127, 0, 0, 1)
	var port = 9972
	var udpAddr = &net.UDPAddr{
		IP:   addr,
		Port: port,
	}
	var err error
	var conn net.Conn

	conn, err = net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		t.Errorf("could not dial udp: %s", err.Error())
		return
	}

	var data []byte
	var m = createDevice(t)
	data, err = m.Marshal()
	if err != nil {
		t.Errorf("could marshal data: %s", err.Error())
		return
	}

	var timer = time.NewTimer(time.Second * 3)
	var counter = 0
	for {
		if counter > 10 {
			break
		}

		_, err = conn.Write(data)
		if err != nil {
			t.Errorf("could not send data: %s", err.Error())
			break
		}

		t.Logf("send success")
		<-timer.C
		timer.Reset(time.Second * 3)
		counter++
	}
}

func createDevice(t *testing.T) *message.Message {
	var (
		version     = message.V1
		supportType = message.SupportText | message.SupportFile
		deviceName  = "test-name"
		deviceType  = "test-type"
		err         error
		id          uuid.UUID
		binId       []byte
	)

	id, err = uuid.NewUUID()
	if err != nil {
		t.Errorf("create uuid fail: %s", err)
		return nil
	}
	binId, err = id.MarshalBinary()
	if err != nil {
		t.Errorf("create uuid fail: %s", err)
		return nil
	}

	return message.New(version, supportType, deviceName, deviceType, binId)
}
