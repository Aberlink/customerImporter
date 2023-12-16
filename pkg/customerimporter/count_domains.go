// Package customerimporter provides ability to read CSV files and extract
// data from it. At current sate it reads file header (first line), and looks
// for 'email' column there. If found, it counts how many times each domain
// was spoted, and returns that values as prints or exports to new .csv file
package customerimporter

import (
	"sort"

	constants "github.com/Aberlink/customerImporter/pkg/constants"
	log "github.com/sirupsen/logrus"
)

// domainCounts is used during domains counting, each spoted domain is created as new key, during
// procesing value is increased each time domain is spoted again
var domainCounts = make(map[string]int)

// countSlice is used to store domains counting, where domain is name of domain and count
// is amount of how many times it was spotted across examined file. Holding
// this in this form allow easy sorting before outputting
type countSlice []struct {
	domain string
	count  int
}

// createCountSlice is a constructor for countSlice object, it takes map as argument,
// that key value pairs will be added to final slice.
// returns CountSlice
func createCountSlice(mapToParse map[string]int) countSlice {
	domainCountSlice := make(countSlice, len(mapToParse))
	domainCountSlice.mapToSlice(mapToParse)
	return domainCountSlice
}

// mapToSlice convers map object into slice, where key,value pairs are converted into
// {key, value} structs and appended to slice. It is helpful when it is
// needed to sort map content
func (cs *countSlice) mapToSlice(mapToParse map[string]int) {
	index := 0
	for key, value := range mapToParse {
		(*cs)[index] = struct {
			domain string
			count  int
		}{key, value}
		index++
	}
}

// sortByCount sorts given countSlice by count values across whole slice.
func (cs *countSlice) sortByCount() {
	sort.Slice(*cs, func(i, j int) bool {
		return (*cs)[i].count > (*cs)[j].count
	})
}

// sortByDomain sorts given countSlice alphabeticaly by domain values across whole slice.
func (cs *countSlice) sortByDomain() {
	sort.Slice(*cs, func(i, j int) bool {
		return ((*cs)[i].domain) < (*cs)[j].domain
	})
}

// sortDomains, based on chosen method, performs sorting of mailing data extracted from
// examined file. Original map is first converted do slice, that is than
// sorted. Methods declaration can be found in [pkg/constants/constants.go].
// In case no proper one as been chosen returns unsorted slice and logs warning
func sortDomains(sortBy string, domainCounts map[string]int) countSlice {
	domainCountSlice := createCountSlice(domainCounts)
	switch sortBy {
	case constants.Count:
		domainCountSlice.sortByCount()
	case constants.Domain:
		domainCountSlice.sortByDomain()
	default:
		log.Warnf("Wrong sorting method '%s', returning unsorted", sortBy)
	}
	return domainCountSlice
}
