package customerimporter

import (
	"testing"
)

func TestIsValidEmail_ValidEmail(t *testing.T) {
	email := "john.doe@example.com"
	result := isValidEmail(email)

	if !result {
		t.Errorf("Expected email to be valid, but it was considered invalid.")
	}
}

func TestIsValidEmail_InvalidEmail(t *testing.T) {
	email := "invalid-email"
	result := isValidEmail(email)

	if result {
		t.Errorf("Expected email to be invalid, but it was considered valid.")
	}
}

func TestGetDomainFromEmail_ValidEmail(t *testing.T) {
	email := "john.doe@example.com"
	expectedDomain := "example.com"

	domain, err := getDomainFromEmail(email)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if domain != expectedDomain {
		t.Errorf("Expected domain %s, but got %s", expectedDomain, domain)
	}
}

func TestGetDomainFromEmail_InvalidEmail(t *testing.T) {
	email := "invalid-email"

	_, err := getDomainFromEmail(email)

	if err == nil {
		t.Errorf("Expected error for invalid email, but no error was returned.")
	}
}
