package customerimporter

import (
	"encoding/csv"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// CSVReader is interface to abstract  csv.Reader, usefull when
// mocking file is needed during testing
type CSVReader interface {
	Read() (record []string, err error)
}

// Opens file with path provided in inputPath. Builds reader that
// moves over it line by line. On first iteration data about header
// is collected and saved, than it moves to proper content. Throws fatal
// error if it is not possible to open file or read its content
func ProcesFile(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	columnsMap, err := getColumnIndex(reader)
	if err != nil {
		log.Fatalf("Error when reading header: %v", err)
	}
	if !hasNeededColumns(columnsMap) {
		return
	}
	columnIterator(reader, columnsMap)
}

// Reads one line from provided file, assuming this is first one
// and might be treated as header. Based on that data, cretes map,
// where key is name of the column and value is its index in the file
// so for line 'name,email,age' it will output [name:0, email:1, age:2]
// throws an error if it is not possible to read line
func getColumnIndex(reader CSVReader) (map[string]int, error) {
	var columnsMap = make(map[string]int)

	header, err := reader.Read()
	if err != nil {
		return columnsMap, err
	}
	for i, name := range header {
		name = strings.ToLower(name)
		columnsMap[name] = i
	}
	return columnsMap, nil
}

// moves over file rows, assuming they are data rows and header was skipped
// earlier. For each row starts processing of its data by handlers that are defined
// in [pkg/customerimporter/handlers.go]
func columnIterator(reader CSVReader, columnsMap map[string]int) {
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf("Error reading CSV: %v", err)
			continue
		}
		handleRow(row, columnsMap)
	}

}

// based on handlers definitions that might be found in [pkg/customerimporter/handlers.go],
// checks whether provided file contains all columns that handlers waiting for. Comparison is based
// on names, so for each handler definition there should be column with strictly the same name.
// If event one column is missing validation is considered as failed
func hasNeededColumns(columnsMap map[string]int) bool {
	for handler := range handlers {
		hasColumn := false
		for column := range columnsMap {
			if handler == column {
				hasColumn = true
			}
		}
		if !hasColumn {
			log.Errorf("Column '%s' not found, cant parse data", handler)
			return false
		}
	}
	return true
}
