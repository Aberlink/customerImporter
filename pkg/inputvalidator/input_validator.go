package inputvalidator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var sortOptions = []string{"count", "domain"}

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error in %s: %s", e.Field, e.Msg)
}

func csvFile(path, field string) error {
	extension := strings.ToLower(filepath.Ext(path))
	if extension != ".csv" {
		return &ValidationError{Field: field, Msg: "File must have a .csv extension"}
	}
	return nil
}

func fileExist(path, field string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return &ValidationError{Field: field, Msg: "File does not exist"}
	}
	return err
}

func validateSortFlag(flag string) error {
	for _, value := range sortOptions {
		if value == flag {
			return nil
		}
	}
	return &ValidationError{Field: "sortBy", Msg: "Sorting options are 'count', 'domain'"}
}

func ValidateFlags(inputPath, outputPath, sortBy string) error {
	if err := fileExist(inputPath, "input"); err != nil {
		return err
	}
	if err := csvFile(inputPath, "input"); err != nil {
		return err
	}
	if err := csvFile(outputPath, "output"); err != nil {
		return err
	}
	if err := validateSortFlag(sortBy); err != nil {
		return err
	}
	return nil
}
