package core

import "net"

type RefreshChan = chan uint8

type Info struct {
	Version      uint8
	SupportType  uint16
	DeviceName   string
	DeviceType   string
	DiscoverTime int64
	timer        RefreshChan
	ip           *net.UDPAddr
}

type Response struct {
	Ip          int    `json:"ip,omitempty"`
	Port        int    `json:"port,omitempty"`
	Version     uint8  `json:"version,omitempty"`
	SupportType uint16 `json:"support_type,omitempty"`
	DeviceId    string `json:"device_id,omitempty"`
	DeviceName  string `json:"device_name,omitempty"`
	DeviceType  string `json:"device_type,omitempty"`
}

func createResponse(id string, i *Info) *Response {
	var ip = i.ip.IP.To4()

	return &Response{
		Ip:          int(ip[0])<<24 + int(ip[1])<<16 + int(ip[2])<<8 + int(ip[3]),
		Port:        i.ip.Port,
		Version:     i.Version,
		SupportType: i.SupportType,
		DeviceId:    uuidArrayToString([]byte(id)),
		DeviceName:  i.DeviceName,
		DeviceType:  i.DeviceType,
	}
}
