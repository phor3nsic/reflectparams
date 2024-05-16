package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/phor3nsic/reflectparams/pkg/globalvar" //Use the file with global variables
	"github.com/phor3nsic/reflectparams/pkg/read"
)

func main() {
	flag.StringVar(&globalvar.Parameter, "p", "", "Param to check. Ex: FUZZ")
	flag.StringVar(&globalvar.ListUrls, "l", "", "List of urls. Ex: /tmp/urls.txt")
	flag.StringVar(&globalvar.Proxy, "x", "", "Replay proxy requests. Ex: http://localhost:8080")
	flag.StringVar(&globalvar.Inject, "i", "", "Inject special chars on parameter. Ex: ['\"><s>s]")
	flag.IntVar(&globalvar.Concorrency, "c", 10, "Concorrency value.")
	flag.Parse()

	if globalvar.Parameter == "" {
		fmt.Println("It is necessary (-p) to set the parameter value...")
		os.Exit(1)
	}

	if globalvar.ListUrls != "" {
		read.ReadFile(globalvar.ListUrls) //If with list file
	} else {
		read.ReadStdin() //If by stdin
	}

}
