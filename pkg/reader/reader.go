package reader

import (
	"bufio"
	"fmt"
	"github.com/jf-Lindberg/lalalint/pkg/rules"
	"log"
	"os"
)

func channelThing(line string, c chan string) {
	line = rules.LintComment(line)
	c <- line
}

func Reader(fileName string) {
	file, err := os.Open("./data/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	c := make(chan string)
	for scanner.Scan() {
		line := scanner.Text()
		go channelThing(line, c)
		line = <-c
		fmt.Printf("%d) %s\n", row, line)
		row++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
