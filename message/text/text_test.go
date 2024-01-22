package text

import (
	"fmt"
	"testing"
)

func TestText(t *testing.T) {
	var raw = "test-ok"

	var msg = New(raw)

	var buffer = msg.Marshal()
	println(fmt.Sprintf("%x", buffer))

	var msg2, err = Unmarshal(buffer)
	if err != nil {
		t.Error("could not unmarshal buffer: " + err.Error())
		return
	}

	if string(msg2.Text) == raw {
		t.Log("ok")
	} else {
		t.Error("could not parse data")
	}
}
