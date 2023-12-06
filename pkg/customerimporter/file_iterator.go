package customerimporter

import (
	"encoding/csv"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type CSVReader interface {
	Read() (record []string, err error)
}

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
