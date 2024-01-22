package text

import (
	"encoding/binary"
	"errors"
	"go-discover-server/message"
)

type Message struct {
	magic  uint16
	length uint16
	Text   []byte
}

func New(text string) *Message {
	return &Message{
		magic:  message.TextMagic,
		length: uint16(len(text)),
		Text:   []byte(text),
	}
}

func (t *Message) Marshal() []byte {
	var length = t.length + 4
	var buffer = make([]byte, 0, length)

	buffer = binary.BigEndian.AppendUint16(buffer, t.magic)
	buffer = binary.BigEndian.AppendUint16(buffer, t.length)
	buffer = append(buffer, t.Text...)

	return buffer
}

func Unmarshal(buffer []byte) (*Message, error) {
	var bufferLength = len(buffer)
	if bufferLength == 0 {
		return nil, errors.New("empty buffer")
	}

	var magic = binary.BigEndian.Uint16(buffer[0:2])
	if magic != message.TextMagic {
		return nil, errors.New("not a text message")
	}

	var textLength = binary.BigEndian.Uint16(buffer[2:4])
	if int(textLength+4) != bufferLength {
		return nil, errors.New("incorrect packet: length error")
	}

	return &Message{
		magic:  message.TextMagic,
		length: textLength,
		Text:   buffer[4:],
	}, nil
}
