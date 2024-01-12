package message

import (
	"github.com/google/uuid"
	"testing"
)

func TestNewMessage(t *testing.T) {
	var (
		version     = V1
		supportType = SupportText | SupportFile
		deviceName  = "test-name"
		deviceType  = "test-type"
		err         error
		id          uuid.UUID
		binId       []byte
	)

	id, err = uuid.NewUUID()
	if err != nil {
		t.Errorf("create uuid fail: %s", err)
		return
	}
	binId, err = id.MarshalBinary()
	if err != nil {
		t.Errorf("create uuid fail: %s", err)
		return
	}

	var message = New(version, supportType, deviceName, deviceType, binId)

	t.Logf(
		"message info (Version=%d, TextSupport=%t, FileSupport=%t, DeviceName=%s, DeviceType=%s)",
		message.Version,
		message.isSupportText(),
		message.isSupportFile(),
		message.DeviceName,
		message.DeviceType,
	)
}

func TestMessageMarshal(t *testing.T) {
	var (
		version     = V1
		supportType = SupportText | SupportFile
		deviceName  = "test-name"
		deviceType  = "test-type"
		err         error
		id          uuid.UUID
		binId       []byte
		buffer      []byte
	)

	id, err = uuid.NewUUID()
	if err != nil {
		t.Errorf("create uuid fail: %s", err)
		return
	}
	binId, err = id.MarshalBinary()
	if err != nil {
		t.Errorf("create uuid fail: %s", err)
		return
	}

	var message = New(version, supportType, deviceName, deviceType, binId)

	buffer, err = message.Marshal()
	if err != nil {
		t.Errorf("Marshal error: %s", err)
		return
	}

	var bufferLen = len(buffer)
	t.Logf("buffer len: %d", bufferLen)

	var message2 *Message
	message2, err = Unmarshal(buffer)
	if err != nil {
		t.Errorf("Unmarshal error: %s", err)
		return
	}

	t.Logf(
		"message info (Version=%d, TextSupport=%t, FileSupport=%t, DeviceName=%s, DeviceType=%s)",
		message.Version,
		message.isSupportText(),
		message.isSupportFile(),
		message.DeviceName,
		message.DeviceType,
	)

	t.Logf(
		"message2 info (Version=%d, TextSupport=%t, FileSupport=%t, DeviceName=%s, DeviceType=%s)",
		message2.Version,
		message2.isSupportText(),
		message2.isSupportFile(),
		message2.DeviceName,
		message2.DeviceType,
	)
}
