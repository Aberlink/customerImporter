package main

import (
	"flag"

	constants "github.com/Aberlink/customerImporter/pkg/constants"
	ci "github.com/Aberlink/customerImporter/pkg/customerimporter"
	v "github.com/Aberlink/customerImporter/pkg/inputvalidator"
	log "github.com/sirupsen/logrus"
)

var print, save bool
var sortBy, inputPath, outputPath string

func main() {
	flag.BoolVar(&print, "print", true, "Print program output.")
	flag.BoolVar(&save, "save", true, "Save program output to .csv file.")
	flag.StringVar(&sortBy, "sortby", constants.Count, "Choose to sort domains by 'count' or 'domain'.")
	flag.StringVar(&inputPath, constants.Input, "customers.csv", "Input .csv file path where first row is a header.")
	flag.StringVar(&outputPath, constants.Output, "sorted_domains.csv", "Output .csv file path.")

	flag.Parse()
	err := v.ValidateFlags(inputPath, outputPath, sortBy)
	if err != nil {
		log.Fatal(err)
	}

	ci.ProcesFile(inputPath)
	ci.OutputDomains(print, save, outputPath, sortBy)

}
