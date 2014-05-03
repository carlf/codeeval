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
		last := "."
		output := ""
		for _, n := range values {
			if n != last {
				output += fmt.Sprintf("%s,", n)
				last = n
			}
		}
		output = strings.TrimRight(output, ",")
		fmt.Println(output)
	}
}
