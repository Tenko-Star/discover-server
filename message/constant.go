package message

const (
	V1 = 1
)

const Magic = 0x1f

const (
	SupportText = 1 << 0
	SupportFile = 1 << 1
)

const (
	TextMagic = Magic<<8 + 0x1
	FileMagic = Magic<<8 + 0x2
)
