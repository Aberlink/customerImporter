package customerimporter

import (
	"bytes"
)

// helper buffer to redirect logs output from console to it. Used
// in testing, to check loging behavior
type LogBuffer struct {
	buffer *bytes.Buffer
}

func (l *LogBuffer) Write(p []byte) (n int, err error) {
	return l.buffer.Write(p)
}
