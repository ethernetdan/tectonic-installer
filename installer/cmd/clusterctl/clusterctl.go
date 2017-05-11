package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

/*
	clusterctl [-s] <file>

	clusterctl deploys clusters using the Tectonic installer executable.

	Arguments
	- Files describing clusters to be deployed or specs to be generated into clusters.

	-g .
*/

var (
	flags = struct {
		generate bool
	}{}

	// file to be parsed
	files []string
)

func init() {
	flag.BoolVar(&flags.generate, "g", false, "output a set of clusters generated using a spec")
	flag.Parse()
}

func main() {
	file := flag.Arg(0)
	if len(file) < 1 {
		fmt.Println("a file must be specified")
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("could not read file: %v", err)
		os.Exit(1)
	}

	if flags.generate {
		if err := generate(data); err != nil {
			fmt.Printf("Error: %v", err)
		}
	}
}
