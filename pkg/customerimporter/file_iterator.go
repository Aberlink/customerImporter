package customerimporter

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func ProcesFile(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	columnsMap, err := getColumnIndex(file)
	if err != nil {
		fmt.Printf("Error when reading header: %v\n", err)
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
		fmt.Printf("Error reading CSV: %v\n", err)
	} //skip first line as it is header
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading CSV: %v\n", err)
			continue
		}
		handleRow(row, columnsMap)
	}

}

func setToFileStart(file *os.File) {
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		fmt.Printf("failed to seek the file: %v\n", err)
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
			fmt.Printf("Column '%s' not found, cant parse data\n", handler)
			return false
		}
	}
	return true
}
