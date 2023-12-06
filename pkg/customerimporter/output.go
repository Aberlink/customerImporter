package customerimporter

import (
	"encoding/csv"
	"fmt"
	"os"

	constants "github.com/Aberlink/customerImporter/pkg/constants"
	log "github.com/sirupsen/logrus"
)

// helper variable that contain header for outputh csv file
var domainsHeader = []string{constants.Domain, constants.Count}

// allows save data about domains and their occurance to csv file. FIrst it
// create new file with provided name, that writes defined header and starts
// to iterete over countSlice, appending its rows to file. Throws an error if
// it is not possible to create file or write to it
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

// Iteates over provided countSlice and logs in output row by row, informing
// about detected domains and custommers that are using it
func printDomains(domainCountSlice countSlice) {
	for _, pos := range domainCountSlice {
		log.Infof("Domain: '%s' has '%s' clients.", pos.domain, fmt.Sprintf("%d", pos.count))
	}

}

// based on print and save flags, outputs data collected during file analysis. Filename
// is name of file that will be saved, sortBy allows to chose way of sorting, (by count
// or in alphabetic order)
func OutputDomains(print, save bool, filename, sortBy string) {
	domainCountSlice := sortDomains(sortBy, domainCounts)
	if print {
		printDomains(domainCountSlice)
	}
	if save {
		if err := saveDomainsToCSV(domainCountSlice, filename); err != nil {
			log.Errorf("Error when saving file: %v", err)
		} else {
			log.Infof("File %s saved", filename)
		}
	}
}
