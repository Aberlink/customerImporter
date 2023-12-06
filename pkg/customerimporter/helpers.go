package customerimporter

import (
	"bytes"
)

//	LogBuffer is helper buffer to redirect logs output from console to it. Used
//
// in testing, to check loging behavior
type logBuffer struct {
	buffer *bytes.Buffer
}

func (l *logBuffer) Write(p []byte) (n int, err error) {
	return l.buffer.Write(p)
}
