package main

import (
	"flag"
	"os"
	"strings"
)

type options struct {
	archivePaths []string
	destination  string
	selector     string
}

func parseCliOptions() (o options) {
	var concatenatedPaths string

	//flag.StringVar(&concatenatedPaths, "archives", "", "paths of mpq archives to search in order, comma delimited")
	flag.StringVar(&concatenatedPaths, "a", "", "paths of mpq archives to search in order, comma delimited")

	//flag.StringVar(&o.destination, "destination", "", "destination directory to extract files into")
	flag.StringVar(&o.destination, "d", "", "destination directory to extract files into")

	//flag.StringVar(&o.selector, "extract", "", "internal path of file to extract from mpq archives")
	flag.StringVar(&o.selector, "x", "", "internal path of file to extract from mpq archives")

	flag.Parse()

	if concatenatedPaths == "" {
		flag.Usage()
		os.Exit(0)
	}

	o.archivePaths = strings.Split(concatenatedPaths, ",")

	if o.destination == "" {
		o.destination = "."
	}

	return
}
