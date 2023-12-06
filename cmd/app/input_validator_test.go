package main

import (
	"os"
	"testing"

	constants "github.com/Aberlink/customerImporter/pkg/constants"
)

func TestCSVFile_ValidExtension(t *testing.T) {
	path := "example.csv"
	field := "someField"

	err := csvFile(path, field)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestCSVFile_InvalidExtension(t *testing.T) {
	path := "example.txt"
	field := "someField"

	err := csvFile(path, field)

	if err == nil {
		t.Error("Expected an error, but got none")
	} else if vErr, ok := err.(*ValidationError); ok {
		expectedMsg := "File must have a .csv extension"
		if vErr.Msg != expectedMsg {
			t.Errorf("Expected error message '%s', but got '%s'", expectedMsg, vErr.Msg)
		}
	} else {
		t.Errorf("Expected ValidationError, but got %T", err)
	}
}

func TestFileExist_FileExists(t *testing.T) {
	path := "existing_file.txt"
	field := "someField"

	// Create an empty file for testing
	_, err := os.Create(path)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(path)

	err = fileExist(path, field)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestFileExist_FileDoesNotExist(t *testing.T) {
	path := "nonexistent_file.csv"
	field := "someField"

	err := fileExist(path, field)

	if err == nil {
		t.Error("Expected an error, but got none")
	} else if vErr, ok := err.(*ValidationError); ok {
		expectedMsg := "File does not exist"
		if vErr.Msg != expectedMsg {
			t.Errorf("Expected error message '%s', but got '%s'", expectedMsg, vErr.Msg)
		}
	} else {
		t.Errorf("Expected ValidationError, but got %T", err)
	}
}

func TestValidateSortFlag_ValidFlag(t *testing.T) {
	flag := constants.Count

	err := validateSortFlag(flag)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestValidateSortFlag_InvalidFlag(t *testing.T) {
	flag := "invalidFlag"

	err := validateSortFlag(flag)

	if err == nil {
		t.Error("Expected an error, but got none")
	} else if vErr, ok := err.(*ValidationError); ok {
		expectedMsg := "Sorting options are 'count', 'domain'"
		if vErr.Msg != expectedMsg {
			t.Errorf("Expected error message '%s', but got '%s'", expectedMsg, vErr.Msg)
		}
	} else {
		t.Errorf("Expected ValidationError, but got %T", err)
	}
}
