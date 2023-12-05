package customerimporter

import (
	"fmt"
	"net/mail"
	"strings"
)

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func getDomainFromEmail(email string) (string, error) {
	if !isValidEmail(email) {
		return "", fmt.Errorf("invalid email address format: %s", email)
	}
	parts := strings.Split(email, "@")
	return parts[1], nil
}
