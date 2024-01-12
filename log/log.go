package log

import (
	"fmt"
	"time"
)

type LogLevel = int

const timeFormat = "15:04:05"

const (
	LevelError = 0b0001
	LevelWarn  = 0b0011
	LevelInfo  = 0b0111
	LevelDebug = 0b1111
)

const (
	maskError = 1 << 0
	maskWarn  = 1 << 1
	maskInfo  = 1 << 2
	maskDebug = 1 << 3
)

var level = LevelInfo

func SetLevel(l LogLevel) {
	level = l
}

func D(format string, args ...any) {
	if level&maskDebug == 0 {
		return
	}

	var t = time.Now().Format(timeFormat)
	var header = "[DBUG][" + t + "] " + format + "\n"

	fmt.Printf(header, args...)
}

func I(format string, args ...any) {
	if level&maskInfo == 0 {
		return
	}

	var t = time.Now().Format(timeFormat)
	var header = "[INFO][" + t + "] " + format + "\n"

	fmt.Printf(header, args...)
}

func W(format string, args ...any) {
	if level&maskWarn == 0 {
		return
	}

	var t = time.Now().Format(timeFormat)
	var header = "[WARN][" + t + "] " + format + "\n"

	fmt.Printf(header, args...)
}

func E(format string, args ...any) {
	if level&maskError == 0 {
		return
	}

	var t = time.Now().Format(timeFormat)
	var header = "[ERRO][" + t + "] " + format + "\n"

	fmt.Printf(header, args...)
}
