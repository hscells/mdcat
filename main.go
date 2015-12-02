package main

import (
	"bufio"
	"io/ioutil"
	"os"
)

func main() {

	args := os.Args[1:]

	input := ""

	if len(args) == 1 {
		data, err := ioutil.ReadFile(args[0])
		if err != nil {
			panic(err)
		}
		input = string(data)
	} else {
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
	}

	Parse(input)

}
