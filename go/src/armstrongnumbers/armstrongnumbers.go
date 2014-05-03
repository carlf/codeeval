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
		length := int(math.Floor(math.Log10(float64(n)))) + 1
		sum := 0
		for x := n; x > 0; x /= 10 {
			digit := x % 10
			sum += int(math.Pow(float64(digit), float64(length)))
		}
		if sum == n {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}
