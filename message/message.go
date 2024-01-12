package message

import (
	"encoding/binary"
	"errors"
)

type Message struct {
	magic         uint8
	Version       uint8
	SupportType   uint16
	deviceNameLen uint32
	deviceTypeLen uint32
	DeviceId      []byte
	DeviceName    string
	DeviceType    string
}

func New(version int, supportType int, deviceName, deviceType string, deviceId []byte) *Message {
	return &Message{
		magic:         Magic,
		Version:       uint8(version),
		SupportType:   uint16(supportType),
		deviceNameLen: uint32(len(deviceName)),
		deviceTypeLen: uint32(len(deviceType)),
		DeviceId:      deviceId,
		DeviceName:    deviceName,
		DeviceType:    deviceType,
	}
}

func Unmarshal(b []byte) (*Message, error) {
	var dataLen = len(b)
	var remain = dataLen
	if dataLen < 12 {
		return nil, errors.New("incorrect package")
	}
	remain = remain - 12

	var (
		magic         = b[0]
		version       = b[1]
		supportType   = binary.BigEndian.Uint16(b[2:4])
		deviceNameLen = binary.BigEndian.Uint32(b[4:8])
		deviceTypeLen = binary.BigEndian.Uint32(b[8:12])
	)

	if magic != Magic {
		return nil, errors.New("incorrect package")
	}

	remain = remain - 16
	if remain < 0 {
		return nil, errors.New("device id is incorrect")
	}
	var deviceId = b[12:28]

	remain = remain - int(deviceNameLen)
	if remain < 0 {
		return nil, errors.New("device name length is incorrect")
	}

	var deviceNameEnd = 28 + deviceNameLen
	var deviceName = string(b[28:deviceNameEnd])

	remain = remain - int(deviceTypeLen)
	if remain < 0 {
		return nil, errors.New("device type length is incorrect")
	}

	var deviceTypeEnd = deviceNameEnd + deviceTypeLen
	var deviceType = string(b[deviceNameEnd:deviceTypeEnd])

	return &Message{
		magic:         magic,
		Version:       version,
		SupportType:   supportType,
		deviceNameLen: deviceNameLen,
		deviceTypeLen: deviceTypeLen,
		DeviceId:      deviceId,
		DeviceName:    deviceName,
		DeviceType:    deviceType,
	}, nil
}

func (m *Message) Marshal() ([]byte, error) {
	var dataLen = 12 + len(m.DeviceName) + len(m.DeviceType) + len(m.DeviceId)
	var buffer = make([]byte, 0, dataLen)

	buffer = append(buffer, m.magic)
	buffer = append(buffer, m.Version)
	buffer = binary.BigEndian.AppendUint16(buffer, m.SupportType)
	buffer = binary.BigEndian.AppendUint32(buffer, uint32(len(m.DeviceName)))
	buffer = binary.BigEndian.AppendUint32(buffer, uint32(len(m.DeviceType)))
	buffer = append(buffer, m.DeviceId...)
	buffer = append(buffer, []byte(m.DeviceName)...)
	buffer = append(buffer, []byte(m.DeviceType)...)

	if len(buffer) != dataLen {
		return nil, errors.New("incorrect data len")
	}

	return buffer, nil
}

func (m *Message) isSupportText() bool {
	return m.SupportType&SupportText > 0
}

func (m *Message) isSupportFile() bool {
	return m.SupportType&SupportFile > 0
}
