package customerimporter

import (
	"encoding/csv"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func ProcesFile(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	columnsMap, err := getColumnIndex(file)
	if err != nil {
		log.Fatalf("Error when reading header: %v", err)
	}
	if !hasNeededColumns(columnsMap) {
		return
	}
	columnIterator(file, columnsMap)
}

func getColumnIndex(file *os.File) (map[string]int, error) {
	var columnsMap = make(map[string]int)
	setToFileStart(file)

	reader := csv.NewReader(file)
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

func columnIterator(file *os.File, columnsMap map[string]int) {
	setToFileStart(file)

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		log.Errorf("Error reading CSV: %v", err)
	} //skip first line as it is header
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

func setToFileStart(file *os.File) error {
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("failed to seek the file: %v", err)
	}
	return nil

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
