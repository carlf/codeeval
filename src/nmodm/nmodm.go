package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
		n, _ := strconv.Atoi(values[0])
		m, _ := strconv.Atoi(values[1])
		mod := n - (n/m)*m
		fmt.Println(mod)
	}
}
