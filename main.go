package main

import (
	"flag"

	"github.com/phor3nsic/reflectparams/pkg/read"
)

var url string
var param string
var listToTest string

func main() {
	flag.StringVar(&param, "p", "", "Param to check")
	flag.StringVar(&listToTest, "l", "", "List Of Urls")
	flag.Parse()
	if listToTest != "" {
		read.ReadFile(listToTest, param)
	} else {
		read.ReadStdin(param)
	}

}
