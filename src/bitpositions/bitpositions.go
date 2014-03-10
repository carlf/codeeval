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
		var n, p1, p2 uint64
		n, _ = strconv.ParseUint(values[0], 10, 64)
		p1, _ = strconv.ParseUint(values[1], 10, 64)
		p2, _ = strconv.ParseUint(values[2], 10, 64)
		bits := strings.Split(fmt.Sprintf("%032b", n), "")
		if bits[len(bits)-int(p1)] == bits[len(bits)-int(p2)] {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}
