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
		seen := make(map[int]bool)
		happy := 1
		for true {
			if n == 1 {
				happy = 1
				break
			}
			if seen[n] {
				happy = 0
				break
			}

			seen[n] = true
			sum := 0
			for n != 0 {
				digit := n % 10
				sum += digit * digit
				n /= 10
			}
			n = sum
		}
		fmt.Println(happy)
	}
}
