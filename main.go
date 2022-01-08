package main

import (
	"flag"
)

var url string
var listToTest string

func main() {
	//flag.StringVar(&url, "u", "", "Url to test")
	flag.StringVar(&listToTest, "l", "", "List Of Urls")
	flag.Parse()

}
