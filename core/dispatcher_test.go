package core

import (
	"encoding/json"
	"github.com/google/uuid"
	"go-discover-server/message"
	"net"
	"testing"
	"time"
)

func TestDispatcher(t *testing.T) {
	var id = addDevice(t)
	addDevice(t)
	addDevice(t)
	addDevice(t)
	addDevice(t)

	var info = FindDevice(id)
	if info == nil {
		t.Errorf("could not find device")
		return
	}

	time.Sleep(time.Second * 11)

	info = FindDevice(id)
	if info != nil {
		t.Errorf("could not remove device")
		return
	}
}

func TestGetAllDevice(t *testing.T) {
	addDevice(t)
	addDevice(t)
	addDevice(t)
	addDevice(t)

	var jsonData, err = json.Marshal(GetAllDevice())
	if err != nil {
		t.Errorf("could not get all devices: %s", err)
		return
	}

	t.Logf("All info: %s", string(jsonData))
}

func addDevice(t *testing.T) string {
	var err error
	var id string
	var _uuid uuid.UUID
	var bUuid []byte

	_uuid, err = uuid.NewUUID()
	if err != nil {
		t.Errorf("create uuid fail: %s", err)
		return ""
	}
	bUuid, err = _uuid.MarshalBinary()
	if err != nil {
		t.Errorf("create uuid fail: %s", err)
		return ""
	}

	var addr = &net.UDPAddr{
		IP:   net.IP([]byte{0x7f, 0x0, 0x0, 0x1}),
		Port: 18088,
	}
	var m = message.New(message.V1, message.SupportText, "test-device", "test-type", bUuid)
	id, err = AddDevice(m, addr)
	if err != nil {
		t.Errorf("add device error: %s", err)
		return ""
	}

	return id
}
