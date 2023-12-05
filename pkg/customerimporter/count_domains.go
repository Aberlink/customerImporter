package customerimporter

import (
	"fmt"
	"sort"
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
	if sortBy == "count" {
		domainCountSlice = sortByCount(domainCountSlice)
	} else if sortBy == "domain" {
		domainCountSlice = sortByDomain(domainCountSlice)
	} else {
		fmt.Printf("Wrong sorting method '%s', returning unsorted", sortBy)
	}
	return domainCountSlice
}
