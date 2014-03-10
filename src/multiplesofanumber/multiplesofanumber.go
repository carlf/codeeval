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
		x, _ := strconv.Atoi(values[0])
		n, _ := strconv.Atoi(values[1])

		for i := n; true; i += n {
			if i >= x {
				fmt.Println(i)
				break
			}
		}
	}
}
