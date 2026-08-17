package stopwatch

import (
	"runtime"
	"time"
)

type snapshot struct {
	timestamp      time.Time
	lineNumber     int
	joinedMessages string
}

func newSnapshot(message string) snapshot {
	return snapshot{
		timestamp:      time.Now(),
		lineNumber:     getLineNumber(),
		joinedMessages: message,
	}
}

func getLineNumber() int {
	_, _, line, ok := runtime.Caller(3)
	if !ok {
		panic("idk")
	}

	return line
}
