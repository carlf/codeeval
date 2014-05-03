package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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
		result := make([]int, 0)
		aset := make(map[string]bool)
		bset := make(map[string]bool)
		sets := strings.Split(line, ";")
		a := strings.Split(sets[0], ",")
		b := strings.Split(sets[1], ",")
		for _, n := range a {
			aset[n] = true
		}
		for _, n := range b {
			bset[n] = true
		}
		for n, _ := range aset {
			if bset[n] {
				value, _ := strconv.Atoi(n)
				result = append(result, value)
			}
		}
		sort.Sort(sort.IntSlice(result))
		for x, n := range result {
			fmt.Printf("%d", n)
			if x < len(result)-1 {
				fmt.Printf(",")
			}
		}
		fmt.Printf("\n")
	}
}
