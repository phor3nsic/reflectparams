package read

import (
	"bufio"
	"log"
	"os"

	"github.com/phor3nsic/reflectparams/pkg/request"
)

func ReadFile(path string, param string) {

	var line string

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		request.Req(line, param)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func ReadStdin(param string) {
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		request.Req(line, param)
	}

}
