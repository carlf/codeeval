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
		fmt.Println(fibonacci(n))
	}
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		a := 0
		b := 1
		for i := 1; i < n; i++ {
			a, b = b, a+b
		}
		return b
	}
}
