package main

import (
	"flag"
	"fmt"

	ci "github.com/Aberlink/customerImporter/pkg/customerimporter"
	v "github.com/Aberlink/customerImporter/pkg/inputvalidator"
)

var print bool
var save bool
var sortBy string
var inputPath string
var outputPath string

func main() {
	flag.BoolVar(&print, "print", true, "Print program output.")
	flag.BoolVar(&save, "save", true, "Save program output to .csv file.")
	flag.StringVar(&sortBy, "sortby", "count", "Choose to sort domains by 'count' or 'domain'.")
	flag.StringVar(&inputPath, "input", "customers.csv", "Input .csv file path where first row is a header.")
	flag.StringVar(&outputPath, "output", "sorted_domains.csv", "Output .csv file path.")

	flag.Parse()
	err := v.ValidateFlags(inputPath, outputPath, sortBy)
	if err != nil {
		fmt.Println(err)
		return
	}

	ci.ProcesFile(inputPath)
	ci.OutputDomains(print, save, outputPath, sortBy)

}
