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
		n, _ := strconv.Atoi(line)
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		fmt.Println(sum)
	}
}
