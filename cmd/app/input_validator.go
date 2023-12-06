package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	constants "github.com/Aberlink/customerImporter/pkg/constants"
)

// helper variable that contain options for output sorting
var sortOptions = []string{constants.Count, constants.Domain}

// error template returned if flags validation fails
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error in %s: %s", e.Field, e.Msg)
}

// checks if given path leads to .csv file. Field argment is added just for
// logging, to make easier to track what kind of file is missing (input or
// output). Throws error if path is not valid
func csvFile(path, field string) error {
	extension := strings.ToLower(filepath.Ext(path))
	if extension != ".csv" {
		return &ValidationError{Field: field, Msg: "File must have a .csv extension"}
	}
	return nil
}

// checks if file behing given path exists. Throws error if not
func fileExist(path, field string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return &ValidationError{Field: field, Msg: "File does not exist"}
	}
	return err
}

// checks if provided sorting method is valid (inside defined sortOptions).
// If not throws an error that shows valid options
func validateSortFlag(flag string) error {
	for _, value := range sortOptions {
		if value == flag {
			return nil
		}
	}
	msg := fmt.Sprintf("Sorting options are '%s', '%s'", constants.Count, constants.Domain)
	return &ValidationError{Field: "sortBy", Msg: msg}
}

// acumulation function, that takes string format flags and validates. Throws error
// if any validation fails
func validateFlags(inputPath, outputPath, sortBy string) error {
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
