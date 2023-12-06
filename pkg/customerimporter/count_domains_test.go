package customerimporter

import (
	"bytes"
	"os"
	"reflect"
	"strings"
	"testing"

	constants "github.com/Aberlink/customerImporter/pkg/constants"
	log "github.com/sirupsen/logrus"
)

var testdomainCounts = map[string]int{
	"example.com": 3,
	"test.com":    1,
	"demo.com":    2,
}

func TestSortByCount(t *testing.T) {
	expectedResult := countSlice{
		{"example.com", 3},
		{"demo.com", 2},
		{"test.com", 1},
	}

	result := sortByCount(mapToSlice(testdomainCounts))

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Sorting by count failed. Expected %v, got %v", expectedResult, result)
	}
}

func TestSortByDomain(t *testing.T) {
	expectedResult := countSlice{
		{"demo.com", 2},
		{"example.com", 3},
		{"test.com", 1},
	}

	result := sortByDomain(mapToSlice(testdomainCounts))

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Sorting by domain failed. Expected %v, got %v", expectedResult, result)
	}
}

func TestSortDomains_ValidSortByCount(t *testing.T) {
	expectedResult := countSlice{
		{"example.com", 3},
		{"demo.com", 2},
		{"test.com", 1},
	}

	result := sortDomains(constants.Count, testdomainCounts)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Sorting domains by count failed. Expected %v, got %v", expectedResult, result)
	}
}

func TestSortDomains_ValidSortByDomain(t *testing.T) {
	expectedResult := countSlice{
		{"demo.com", 2},
		{"example.com", 3},
		{"test.com", 1},
	}

	result := sortDomains(constants.Domain, testdomainCounts)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Sorting domains by domain failed. Expected %v, got %v", expectedResult, result)
	}
}

func TestSortDomains_InvalidSortMethod(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	defer func() {
		log.SetOutput(os.Stderr)
	}()
	log.SetOutput(&LogBuffer{buffer: &buf})

	sortDomains("invalid", testdomainCounts)

	logOutput := buf.String()
	expectedFragment := "Wrong sorting method"
	if !strings.Contains(logOutput, expectedFragment) {
		t.Errorf("Expected log output to contain '%s', but got '%s'", expectedFragment, logOutput)
	}
}
