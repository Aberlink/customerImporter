package customerimporter

import "fmt"

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
		fmt.Printf("Error extracting domain from '%s': %v, skiping\n", email, err)
		return
	} else {
		domainCounts[domain]++
	}
}
