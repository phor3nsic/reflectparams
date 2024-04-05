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

	var line string

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			if globalvar.Inject != "" {
				param := globalvar.Parameter + globalvar.Inject
				changedUrl := strings.Replace(url, globalvar.Parameter, param, -1)
				request.Req(changedUrl, param)
			} else {
				request.Req(url, globalvar.Parameter)
			}

		}(line)

	}

	wg.Wait()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func ReadStdin() {
	var wg sync.WaitGroup
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			if globalvar.Inject != "" {
				param := globalvar.Parameter + globalvar.Inject
				changedUrl := strings.Replace(url, globalvar.Parameter, param, -1)
				request.Req(changedUrl, param)
			} else {
				request.Req(url, globalvar.Parameter)
			}

		}(line)
	}

	wg.Wait()

}
