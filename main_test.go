package main

import (
	"github.com/google/uuid"
	"go-discover-server/message"
	"go-discover-server/message/discover"
	"net"
	"strings"
	"sync"
	"testing"
	"time"
)

var group = sync.WaitGroup{}

func TestAll(t *testing.T) {
	group.Add(2)

	go testUDP(t)
	time.Sleep(time.Second)
	go testGroup(t)

	group.Wait()
}

func testUDP(t *testing.T) {
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
	t.Logf("create device: %s", string(m.DeviceId))

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

	group.Done()
}

func testGroup(t *testing.T) {
	var addr = net.IPv4(239, 233, 1, 1)
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
	t.Logf("create device: %s", string(m.DeviceId))

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

	group.Done()
}

func createDevice(t *testing.T) *discover.Message {
	var (
		version     = message.V1
		supportType = message.SupportText | message.SupportFile
		deviceName  = "test-name"
		deviceType  = "test-type"
		err         error
		id          uuid.UUID
	)

	id, err = uuid.NewUUID()
	if err != nil {
		t.Errorf("create uuid fail: %s", err)
		return nil
	}
	idStr := strings.Replace(id.String(), "-", "", -1)

	return discover.New(version, supportType, deviceName, deviceType, idStr)
}

func TestHex2Str(t *testing.T) {
	var arr = []byte{0x64, 0x39, 0x33, 0x31, 0x65, 0x30, 0x63, 0x31, 0x35, 0x65, 0x61, 0x33, 0x34, 0x31, 0x39, 0x65}

	t.Logf("result byte array: %s", string(arr))

	var str = "1e5dca2eb37f11ee80e418c04d4e3a9a"

	t.Logf("result str: %s", string([]byte(str)))
}
