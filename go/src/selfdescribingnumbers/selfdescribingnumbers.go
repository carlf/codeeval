package main

import (
	"fmt"
	"io/ioutil"
	"math"
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
		var digit_counts [10]int
		var position_values [10]int
		for i := int(math.Log10(float64(n))); i >= 0; i-- {
			x := n % 10
			position_values[i] = x
			digit_counts[x] += 1
			n /= 10
		}
		match := 1
		for i := range digit_counts {
			if digit_counts[i] != position_values[i] {
				match = 0
				break
			}
		}
		fmt.Println(match)
	}
}
