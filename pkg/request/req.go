package request

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Req(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	sb := string(body)
	log.Printf(sb)
}
