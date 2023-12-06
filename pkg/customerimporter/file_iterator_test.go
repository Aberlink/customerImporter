package customerimporter

import (
	"io"
	"reflect"
	"testing"
)

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

var mockFileContent = [][]string{
	{"name", "email", "age"},
	{"John", "john@example.com", "25"},
	{"Alice", "alice@example.com", "22"},
	{"Tomas", "tomas@example.com", "12"},
}
var mockColumnIndexes = map[string]int{"name": 0, "email": 1, "age": 2}

func TestGetColumnIndex(t *testing.T) {

	mockReader := &mockCSVReader{data: mockFileContent}
	result, _ := getColumnIndex(mockReader)

	if !reflect.DeepEqual(mockColumnIndexes, result) {
		t.Errorf("Invalid column mapping. Expected %v, got %v", mockColumnIndexes, result)
	}
}

func TestColumnIterator(t *testing.T) {
	var oldHandlers = handlers
	handlers = make(map[string]func(cell string, optionalArgs ...interface{}))
	handlers["email"] = func(cell string, optionalArgs ...interface{}) {}

	defer func() {
		handlers = oldHandlers
	}()

	mockReader := &mockCSVReader{data: mockFileContent}

	columnIterator(mockReader, mockColumnIndexes)
}

func TestHasNeededColumns(t *testing.T) {
	tests := []struct {
		name     string
		setup    func()
		expected bool
	}{
		{
			name:     "Pass",
			setup:    func() { handlers["email"] = func(cell string, optionalArgs ...interface{}) {} },
			expected: true,
		},
		{
			name:     "Fail",
			setup:    func() { handlers["gender"] = func(cell string, optionalArgs ...interface{}) {} },
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var oldHandlers = handlers
			handlers = make(map[string]func(cell string, optionalArgs ...interface{}))

			defer func() {
				handlers = oldHandlers
			}()

			test.setup()
			has := hasNeededColumns(mockColumnIndexes)
			if has != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, has)
			}
		})
	}
}
