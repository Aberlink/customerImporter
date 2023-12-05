package customerimporter

import (
	log "github.com/sirupsen/logrus"
)

var handlers = map[string]func(
	cell string, optionalArgs ...interface{}){
	"email": handleEmail,
}

func handleRow(row []string, columnsMap map[string]int) {
	for column, handler := range handlers {
		if index, ok := columnsMap[column]; ok {
			handler(row[index])
		}
	}
}

func handleEmail(email string, optionalArgs ...interface{}) {
	domain, err := getDomainFromEmail(email)
	if err != nil {
		log.Warnf("Failed to extract domain from '%s': %v", email, err)
		return
	} else {
		domainCounts[domain]++
	}
}
