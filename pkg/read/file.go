package read

import (
	"bufio"
	"log"
	"os"
	"sync"

	"github.com/phor3nsic/reflectparams/pkg/request"
)

func ReadFile(path string, param string) {
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
			request.Req(url, param)
		}(line)

	}

	wg.Wait()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func ReadStdin(param string) {
	var wg sync.WaitGroup
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			request.Req(url, param)
		}(line)
	}

	wg.Wait()

}
