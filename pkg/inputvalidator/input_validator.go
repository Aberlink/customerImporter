package inputvalidator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	constants "github.com/Aberlink/customerImporter/pkg/constants"
)

var sortOptions = []string{constants.Count, constants.Domain}

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
	msg := fmt.Sprintf("Sorting options are '%s', '%s'", constants.Count, constants.Domain)
	return &ValidationError{Field: "sortBy", Msg: msg}
}

func ValidateFlags(inputPath, outputPath, sortBy string) error {
	if err := fileExist(inputPath, constants.Input); err != nil {
		return err
	}
	if err := csvFile(inputPath, constants.Input); err != nil {
		return err
	}
	if err := csvFile(outputPath, constants.Output); err != nil {
		return err
	}
	if err := validateSortFlag(sortBy); err != nil {
		return err
	}
	return nil
}
