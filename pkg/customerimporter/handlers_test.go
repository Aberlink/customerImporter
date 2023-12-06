package customerimporter

import (
	"bytes"
	"os"
	"reflect"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestHandleRow(t *testing.T) {
	mockRow := []string{"john@example.com", "30"}
	mockColumnsMap := map[string]int{"email": 0, "age": 1}
	var processedEmails []string

	oldHandleEmail := handlers["email"]
	defer func() {
		handlers["email"] = oldHandleEmail
	}()

	handlers["email"] = func(email string, optionalArgs ...interface{}) {
		processedEmails = append(processedEmails, email)
	}
	handleRow(mockRow, mockColumnsMap)
	expectedProcessedEmails := []string{"john@example.com"}
	if !reflect.DeepEqual(processedEmails, expectedProcessedEmails) {
		t.Errorf("Processed emails do not match. Expected %v, got %v", expectedProcessedEmails, processedEmails)
	}
}

func TestHandleEmail_Success(t *testing.T) {
	mockEmail := "john@example.com"
	olddomainCounts := domainCounts
	domainCounts = make(map[string]int)
	defer func() {
		domainCounts = olddomainCounts
	}()

	if !reflect.DeepEqual(domainCounts["example.com"], 0) {
		t.Errorf("domainCount holds invalid number. Expected %v, got %v", 0, domainCounts["example.com"])
	}

	handleEmail(mockEmail)

	if !reflect.DeepEqual(domainCounts["example.com"], 1) {
		t.Errorf("domainCount holds invalid number. Expected %v, got %v", 1, domainCounts["example.com"])
	}
}

func TestHandleEmail_Failure(t *testing.T) {
	mockEmail := "invalid-email"

	var buf bytes.Buffer
	log.SetOutput(&buf)

	defer func() {
		log.SetOutput(os.Stderr)
	}()
	log.SetOutput(&LogBuffer{buffer: &buf})

	handleEmail(mockEmail)

	logOutput := buf.String()
	expectedFragment := "warning"
	if !strings.Contains(logOutput, expectedFragment) {
		t.Errorf("Expected log output to contain '%s', but got '%s'", expectedFragment, logOutput)
	}
}
