package core

import (
	"go-discover-server/log"
	"go-discover-server/message/discover"
	"net"
	"sync"
	"time"
)

var deviceMap = make(map[string]*Info)
var locker = sync.Mutex{}

func AddDevice(m *discover.Message, addr *net.UDPAddr) (string, error) {
	locker.Lock()
	defer locker.Unlock()

	var id = string(m.DeviceId)
	if info, ok := deviceMap[id]; ok {
		select {
		case info.timer <- 1:
			log.D("device refreshed: %s", string(m.DeviceId))
			return id, nil

		default:
			// write to a closed channel
			return createInfo(m, addr)
		}
	}

	return createInfo(m, addr)
}

func RemoveDevice(id string) {
	locker.Lock()
	defer locker.Unlock()

	delete(deviceMap, id)
	log.I("device removed: %s", id)
}

func FindDevice(id string) *Info {
	locker.Lock()
	defer locker.Unlock()

	if info, ok := deviceMap[id]; ok {
		return info
	}

	return nil
}

type FullInfo struct {
	id   string
	info *Info
}

func GetAllDevice() []*Response {
	locker.Lock()
	var devices = make([]FullInfo, 0)
	for id, info := range deviceMap {
		devices = append(devices, FullInfo{
			id:   id,
			info: info,
		})
	}
	locker.Unlock()

	var responses = make([]*Response, 0)
	for _, device := range devices {
		responses = append(responses, createResponse(device.id, device.info))
	}

	return responses
}

func createTimer(id string) RefreshChan {
	var refreshChan = make(chan uint8)
	var timer = time.NewTimer(time.Second * 10)

	go func() {
		for {
			select {
			case <-timer.C:
				RemoveDevice(id)
				close(refreshChan)
				return

			case <-refreshChan:
				timer.Reset(time.Second * 10)
			}
		}
	}()

	return refreshChan
}

func createInfo(m *discover.Message, addr *net.UDPAddr) (string, error) {
	var id = string(m.DeviceId)
	var info = &Info{
		Version:      m.Version,
		SupportType:  m.SupportType,
		DeviceName:   m.DeviceName,
		DeviceType:   m.DeviceType,
		DiscoverTime: time.Now().Unix(),
		timer:        createTimer(id),
		ip:           addr,
	}

	deviceMap[id] = info
	log.I("device added: %s", string(m.DeviceId))

	return id, nil
}
