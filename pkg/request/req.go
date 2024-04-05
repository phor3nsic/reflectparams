package request

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/phor3nsic/reflectparams/pkg/globalvar"
)

func noRedirect(req *http.Request, via []*http.Request) error {
	return errors.New("Don't redirect!")
}

func Req(urlStr, param string, proxyAddr ...string) {
	var transport *http.Transport
	if len(proxyAddr) > 0 {
		proxyURL, err := url.Parse(proxyAddr[0])
		if err != nil {
			log.Fatal("Error to check proxy addr:", err)
			return
		}
		transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			DialContext: (&net.Dialer{
				Timeout:   20 * time.Second,
				KeepAlive: time.Second,
				DualStack: true,
			}).DialContext,
		}
	} else {
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			DialContext: (&net.Dialer{
				Timeout:   20 * time.Second,
				KeepAlive: time.Second,
				DualStack: true,
			}).DialContext,
		}
	}

	client := &http.Client{
		Transport: transport,
	}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return nil
	}

	resp, err := client.Get(urlStr)
	if err != nil {
		log.Fatal("Error to make request::", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error to read response:", err)
		return
	}

	sb := string(body)
	if strings.Contains(sb, param) {
		fmt.Println(urlStr)
		if globalvar.Proxy != "" {
			Req(urlStr, "q2q4uk69dd91x0sgvmzD7f7S", globalvar.Proxy)
		}

	}
}
