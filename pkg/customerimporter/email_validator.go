package customerimporter

import (
	"fmt"
	"net/mail"
	"strings"
)

// reports whether provided string is valid email adress
func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// extracts mailing domain from provided string and returns it.
// If provided string is not valid email throws an error
func getDomainFromEmail(email string) (string, error) {
	if !isValidEmail(email) {
		return "", fmt.Errorf("invalid email address format: %s", email)
	}
	parts := strings.Split(email, "@")
	return parts[1], nil
}
