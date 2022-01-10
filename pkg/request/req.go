package request

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Req(url string, param string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	sb := string(body)
	if strings.Contains(sb, param) {
		fmt.Println(url)
	}

}
