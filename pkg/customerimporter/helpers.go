package customerimporter

import (
	"bytes"
)

type LogBuffer struct {
	buffer *bytes.Buffer
}

func (l *LogBuffer) Write(p []byte) (n int, err error) {
	return l.buffer.Write(p)
}
