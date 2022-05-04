package request

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

func noRedirect(req *http.Request, via []*http.Request) error {
	return errors.New("Don't redirect!")
}

func Req(url string, param string) {

	var transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   20 * time.Second,
			KeepAlive: time.Second,
			DualStack: true,
		}).DialContext,
	}

	client := &http.Client{
		Transport: transport,
	}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return nil
	}

	resp, err := client.Get(url)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//out := fmt.Sprintf("{\"name\":\"Reflect Params\",\"severity\":\"info\",\"url\":\"%s\"}", url)
	sb := string(body)
	if strings.Contains(sb, param) {
		fmt.Println(url)
	}

}
