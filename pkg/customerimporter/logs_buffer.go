package customerimporter

import (
	"bytes"
	"io"
)

type LogBuffer struct {
	buffer *bytes.Buffer
}

func (l *LogBuffer) Write(p []byte) (n int, err error) {
	return l.buffer.Write(p)
}

type mockCSVReader struct {
	data  [][]string
	index int
}

func (m *mockCSVReader) Read() (record []string, err error) {
	if m.index >= len(m.data) {
		return nil, io.EOF
	}
	record = m.data[m.index]
	m.index++
	return record, nil
}
