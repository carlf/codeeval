package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]

	values := map[string]int{
		"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5,
		"6": 6, "7": 7, "8": 8, "9": 9, "a": 10, "b": 11,
		"c": 12, "d": 13, "e": 14, "f": 15,
	}

	bytes, _ := ioutil.ReadFile(filename)
	input := string(bytes[:])
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			break
		}
		value := 0
		place_value := int(math.Pow(16, float64(len(line)-1)))
		for _, n := range line {
			value += values[string(n)] * place_value
			place_value /= 16
		}
		fmt.Println(value)
	}
}
