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
		values := strings.Fields(line)
		a, _ := strconv.Atoi(values[0])
		b, _ := strconv.Atoi(values[1])
		n, _ := strconv.Atoi(values[2])
		for i := 1; i <= n; i++ {
			print_num := true
			if i%a == 0 {
				fmt.Print("F")
				print_num = false
			}
			if i%b == 0 {
				fmt.Print("B")
				print_num = false
			}
			if print_num {
				fmt.Print(i)
			}
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
	}
}
