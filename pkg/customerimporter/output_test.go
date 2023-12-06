package customerimporter

// import (
// 	"bytes"
// 	"fmt"
// 	"os"
// 	"strings"
// 	"testing"

// 	constants "github.com/Aberlink/customerImporter/pkg/constants"
// 	log "github.com/sirupsen/logrus"
// 	"github.com/sirupsen/logrus/hooks/test"
// )

// type fileWriter interface {
// 	Create(name string) (*os.File, error)
// 	WriteFile(name string, data []byte, perm os.FileMode) error
// }

// type mockFileWriter struct {
// 	mockFile *os.File
// 	err      error
// }

// func (m *mockFileWriter) Create(name string) (*os.File, error) {
// 	return m.mockFile, m.err
// }

// func (m *mockFileWriter) WriteFile(name string, data []byte, perm os.FileMode) error {
// 	return m.err
// }

// func TestSaveDomainsToCSV(t *testing.T) {
// 	// Mock the file writer
// 	mockFile := &os.File{}
// 	mockFileWriter := &mockFileWriter{mockFile: mockFile}

// 	// Replace the os.Create function temporarily for testing
// 	oldOsCreate := os.Create
// 	defer func() {
// 		os.Create = oldOsCreate
// 	}()
// 	os.Create = mockFileWriter.Create

// 	domainCountSlice := countSlice{
// 		{"example.com", 3},
// 		{"test.com", 1},
// 		{"demo.com", 2},
// 	}

// 	var buf bytes.Buffer
// 	log.SetOutput(&buf)

// 	err := saveDomainsToCSV(domainCountSlice, "test.csv")

// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}

// 	// Verify the log output
// 	expectedLog := "File test.csv saved"
// 	if !strings.Contains(buf.String(), expectedLog) {
// 		t.Errorf("Expected log output to contain '%s', but got '%s'", expectedLog, buf.String())
// 	}
// }

// func TestSaveDomainsToCSV_ErrorCreatingFile(t *testing.T) {
// 	// Mock the file writer to simulate an error when creating a file
// 	mockFileWriter := &mockFileWriter{err: fmt.Errorf("error creating file")}

// 	// Replace the os.Create function temporarily for testing
// 	oldOsCreate := os.Create
// 	defer func() {
// 		os.Create = oldOsCreate
// 	}()
// 	os.Create = mockFileWriter.Create

// 	domainCountSlice := countSlice{
// 		{"example.com", 3},
// 		{"test.com", 1},
// 		{"demo.com", 2},
// 	}

// 	var buf bytes.Buffer
// 	log.SetOutput(&buf)

// 	err := saveDomainsToCSV(domainCountSlice, "test.csv")

// 	if err == nil {
// 		t.Error("Expected an error, but got none")
// 	}

// 	// Verify the log output
// 	expectedLog := "Error when saving file: error creating file"
// 	if !strings.Contains(buf.String(), expectedLog) {
// 		t.Errorf("Expected log output to contain '%s', but got '%s'", expectedLog, buf.String())
// 	}
// }

// func TestPrintDomains(t *testing.T) {
// 	// Capture log messages
// 	hook := test.NewGlobal()
// 	defer hook.Reset()

// 	domainCountSlice := countSlice{
// 		{"example.com", 3},
// 		{"test.com", 1},
// 		{"demo.com", 2},
// 	}

// 	printDomains(domainCountSlice)

// 	// Verify the log output
// 	expectedLogs := []string{
// 		"Domain: 'example.com' has '3' clients.",
// 		"Domain: 'test.com' has '1' clients.",
// 		"Domain: 'demo.com' has '2' clients.",
// 	}

// 	for i, entry := range hook.AllEntries() {
// 		if entry.Message != expectedLogs[i] {
// 			t.Errorf("Expected log output '%s', but got '%s'", expectedLogs[i], entry.Message)
// 		}
// 	}
// }

// func TestOutputDomains(t *testing.T) {
// 	// Mock the file writer
// 	mockFile := &os.File{}
// 	mockFileWriter := &mockFileWriter{mockFile: mockFile}

// 	// Replace the os.Create function temporarily for testing
// 	oldOsCreate := os.Create
// 	defer func() {
// 		os.Create = oldOsCreate
// 	}()
// 	os.Create = mockFileWriter.Create

// 	// Capture log messages
// 	hook := test.NewGlobal()
// 	defer hook.Reset()

// 	// Set up the domainCounts map
// 	domainCounts = map[string]int{
// 		"example.com": 3,
// 		"test.com":    1,
// 		"demo.com":    2,
// 	}

// 	// Run the function with save=true, print=true, and a valid filename
// 	OutputDomains(true, true, "test.csv", constants.Count)

// 	// Verify the log output for file save
// 	expectedLogSave := "File test.csv saved"
// 	if !strings.Contains(hook.LastEntry().Message, expectedLogSave) {
// 		t.Errorf("Expected log output to contain '%s', but got '%s'", expectedLogSave, hook.LastEntry().Message)
// 	}

// 	// Verify the log output for domain print
// 	expectedLogsPrint := []string{
// 		"Domain: 'example.com' has '3' clients.",
// 		"Domain: 'test.com' has '1' clients.",
// 		"Domain: 'demo.com' has '2' clients.",
// 	}

// 	for i, entry := range hook.AllEntries() {
// 		if entry.Message != expectedLogsPrint[i] {
// 			t.Errorf("Expected log output '%s', but got '%s'", expectedLogsPrint[i], entry.Message)
// 		}
// 	}
// }

// // Cleanup the domainCounts map after the tests
// func TestMain(m *testing.M) {
// 	defer func() {
// 		domainCounts = make(map[string]int)
// 	}()
// 	m.Run()
// }
