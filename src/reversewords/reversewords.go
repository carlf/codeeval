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
		words := strings.Fields(line)
		for i := len(words) - 1; i >= 0; i-- {
			fmt.Printf("%s", words[i])
			if i != 0 {
				fmt.Printf(" ")
			}
		}
		fmt.Println("")
	}
}
