package customerimporter

import (
	log "github.com/sirupsen/logrus"
)

// handlers stores data parsing functions as values, bounded to columns
// names as keys. Might be extended with additional pairs in case that
// more columns needs to be checked during one program run
var handlers = map[string]func(
	cell string, optionalArgs ...interface{}){
	"email": handleEmail,
}

// moves over handlers definition, checks if column that certain handler
// parse exist in file mapping. If so, it is processed by function associated
// to this column
func handleRow(row []string, columnsMap map[string]int) {
	for column, handler := range handlers {
		if index, ok := columnsMap[column]; ok {
			handler(row[index])
		}
	}
}

// handler that is used to count how many times each domain occured in given file.
// it extracts emain domain from given string, and increase its counter in domainCounts
// map, that could be found in [pkg/customerimporter/count_domains.go]
func handleEmail(email string, optionalArgs ...interface{}) {
	domain, err := getDomainFromEmail(email)
	if err != nil {
		log.Warnf("Failed to extract domain from '%s': %v", email, err)
		return
	} else {
		domainCounts[domain]++
	}
}
