package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]

	bytes, _ := ioutil.ReadFile(filename)
	input := string(bytes[:])
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			break
		}
		values := strings.Split(line, ",")
		s := values[0]
		c := values[1]
		index := strings.LastIndex(s, c)
		fmt.Println(index)
	}
}
