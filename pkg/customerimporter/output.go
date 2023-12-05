package customerimporter

import (
	"encoding/csv"
	"fmt"
	"os"
)

var domainsHeader = []string{"domain", "count"}

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
			fmt.Printf("Error when saving %v row: %v\n", row, err)
			continue
		}
	}
	return nil
}

func printDomains(domainCountSlice countSlice) {
	for _, pos := range domainCountSlice {
		fmt.Printf("Domain: '%s' has '%s' clients.\n", pos.domain, fmt.Sprintf("%d", pos.count))
	}

}

func OutputDomains(print, save bool, filename, sortBy string) {
	domainCountSlice := sortDomains(sortBy)
	if save {
		if err := saveDomainsToCSV(domainCountSlice, filename); err != nil {
			fmt.Printf("Error when saving file: %v\n", err)
		} else {
			fmt.Printf("File %s saved\n", filename)
		}
	}
	if print {
		printDomains(domainCountSlice)
	}
}
