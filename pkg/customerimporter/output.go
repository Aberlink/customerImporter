package customerimporter

import (
	"encoding/csv"
	"fmt"
	"os"

	constants "github.com/Aberlink/customerImporter/pkg/constants"
	log "github.com/sirupsen/logrus"
)

var domainsHeader = []string{constants.Domain, constants.Count}

func saveDomainsToCSV(domainCountSlice countSlice, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(domainsHeader); err != nil {
		return err
	}

	for _, pos := range domainCountSlice {
		row := []string{pos.domain, fmt.Sprintf("%d", pos.count)}
		if err := writer.Write(row); err != nil {
			log.Warnf("Error when saving %v row: %v", row, err)
			continue
		}
	}
	return nil
}

func printDomains(domainCountSlice countSlice) {
	for _, pos := range domainCountSlice {
		log.Infof("Domain: '%s' has '%s' clients.", pos.domain, fmt.Sprintf("%d", pos.count))
	}

}

func OutputDomains(print, save bool, filename, sortBy string) {
	domainCountSlice := sortDomains(sortBy, domainCounts)
	if save {
		if err := saveDomainsToCSV(domainCountSlice, filename); err != nil {
			log.Errorf("Error when saving file: %v", err)
		} else {
			log.Infof("File %s saved", filename)
		}
	}
	if print {
		printDomains(domainCountSlice)
	}
}
