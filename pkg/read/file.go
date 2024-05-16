package read

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/phor3nsic/reflectparams/pkg/globalvar"
	"github.com/phor3nsic/reflectparams/pkg/request"
)

func ReadFile(path string) {
	var wg sync.WaitGroup

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	maxConcurrency := globalvar.Concorrency    //max concorrency
	sem := make(chan struct{}, maxConcurrency) //semafore

	for scanner.Scan() {
		line := scanner.Text()
		sem <- struct{}{}

		go func(url string) {
			defer func() {
				<-sem
			}()

			if globalvar.Inject != "" {
				param := globalvar.Parameter + globalvar.Inject
				changedUrl := strings.Replace(url, globalvar.Parameter, param, -1)
				request.Req(changedUrl, param)
			} else {
				request.Req(url, globalvar.Parameter)
			}
		}(line)
	}

	for i := 0; i < maxConcurrency; i++ {
		sem <- struct{}{}
	}

	wg.Wait()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func ReadStdin() {
	var wg sync.WaitGroup
	//var line string
	scanner := bufio.NewScanner(os.Stdin)
	maxConcurrency := globalvar.Concorrency
	sem := make(chan struct{}, maxConcurrency)

	for scanner.Scan() {
		line := scanner.Text()
		sem <- struct{}{}

		go func(url string) {
			defer func() {
				<-sem
			}()

			if globalvar.Inject != "" {
				param := globalvar.Parameter + globalvar.Inject
				changedUrl := strings.Replace(url, globalvar.Parameter, param, -1)
				request.Req(changedUrl, param)
			} else {
				request.Req(url, globalvar.Parameter)
			}
		}(line)
	}

	for i := 0; i < maxConcurrency; i++ {
		sem <- struct{}{}
	}

	wg.Wait()

}
