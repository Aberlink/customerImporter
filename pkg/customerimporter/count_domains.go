package customerimporter

import (
	"sort"

	constants "github.com/Aberlink/customerImporter/pkg/constants"
	log "github.com/sirupsen/logrus"
)

var domainCounts = make(map[string]int)

type countSlice []struct {
	domain string
	count  int
}

func mapToSlice(mapToParse map[string]int) countSlice {
	var domainCountSlice countSlice

	for key, value := range mapToParse {
		domainCountSlice = append(domainCountSlice, struct {
			domain string
			count  int
		}{key, value})
	}
	return domainCountSlice
}

func sortByCount(domainCountSlice countSlice) countSlice {
	sort.SliceStable(domainCountSlice, func(i, j int) bool {
		return domainCountSlice[i].count > domainCountSlice[j].count
	})
	return domainCountSlice
}

func sortByDomain(domainCountSlice countSlice) countSlice {
	sort.SliceStable(domainCountSlice, func(i, j int) bool {
		return (domainCountSlice[i].domain) < domainCountSlice[j].domain
	})
	return domainCountSlice
}

func sortDomains(sortBy string) countSlice {
	domainCountSlice := mapToSlice(domainCounts)
	switch sortBy {
	case constants.Count:
		domainCountSlice = sortByCount(domainCountSlice)
	case constants.Domain:
		domainCountSlice = sortByDomain(domainCountSlice)
	default:
		log.Warnf("Wrong sorting method '%s', returning unsorted", sortBy)
	}
	return domainCountSlice
}
