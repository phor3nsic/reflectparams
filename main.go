package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/phor3nsic/reflectparams/pkg/globalvar"
	"github.com/phor3nsic/reflectparams/pkg/read"
)

func main() {
	flag.StringVar(&globalvar.Parameter, "p", "", "Param to check")
	flag.StringVar(&globalvar.ListUrls, "l", "", "List Of Urls")
	flag.StringVar(&globalvar.Proxy, "x", "", "Replay proxy requests")
	flag.StringVar(&globalvar.Inject, "i", "", "Inject special chars on parameter...")
	flag.Parse()

	if globalvar.Parameter == "" {
		fmt.Println("It is necessary (-p) to set the parameter value...")
		os.Exit(1)
	}

	if globalvar.ListUrls != "" {
		read.ReadFile(globalvar.ListUrls)
	} else {
		read.ReadStdin()
	}

}
