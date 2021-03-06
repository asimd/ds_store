package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/asimd/ds_store"
)

func main() {

	var fileName string

	flag.StringVar(&fileName, "i", "", "DS_Store input file")
	flag.Parse()

	if fileName == "" {
		if len(os.Args) >= 2 {
			fileName = os.Args[1]
		} else {
			fileName = ".DS_Store"
		}
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		panic("No file given or it does not exist:" + fileName)
	}

	dat, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	a, err := ds_store.NewAllocator(dat)

	filenames, err := a.TraverseFromRootNode()
	if err != nil {
		panic(err)
	}

	for _, f := range filenames {
		fmt.Printf("File: %s \n", f)
	}
}
