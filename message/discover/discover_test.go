package discover

import (
	"github.com/google/uuid"
	"go-discover-server/message"
	"testing"
)

func TestNewMessage(t *testing.T) {
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
		return
	}

	var discover = New(version, supportType, deviceName, deviceType, id.String())

	t.Logf(
		"message info (Version=%d, TextSupport=%t, FileSupport=%t, DeviceName=%s, DeviceType=%s)",
		discover.Version,
		discover.isSupportText(),
		discover.isSupportFile(),
		discover.DeviceName,
		discover.DeviceType,
	)
}

func TestMessageMarshal(t *testing.T) {
	var (
		version     = message.V1
		supportType = message.SupportText | message.SupportFile
		deviceName  = "test-name"
		deviceType  = "test-type"
		err         error
		id          uuid.UUID
		buffer      []byte
	)

	id, err = uuid.NewUUID()
	if err != nil {
		t.Errorf("create uuid fail: %s", err)
		return
	}

	var discover = New(version, supportType, deviceName, deviceType, id.String())

	buffer, err = discover.Marshal()
	if err != nil {
		t.Errorf("Marshal error: %s", err)
		return
	}

	var bufferLen = len(buffer)
	t.Logf("buffer len: %d", bufferLen)

	var discover2 *Message
	discover2, err = Unmarshal(buffer)
	if err != nil {
		t.Errorf("Unmarshal error: %s", err)
		return
	}

	t.Logf(
		"discover info (Version=%d, TextSupport=%t, FileSupport=%t, DeviceName=%s, DeviceType=%s)",
		discover.Version,
		discover.isSupportText(),
		discover.isSupportFile(),
		discover.DeviceName,
		discover.DeviceType,
	)

	t.Logf(
		"discover info (Version=%d, TextSupport=%t, FileSupport=%t, DeviceName=%s, DeviceType=%s)",
		discover2.Version,
		discover2.isSupportText(),
		discover2.isSupportFile(),
		discover2.DeviceName,
		discover2.DeviceType,
	)
}
