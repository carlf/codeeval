package main

import (
	"fmt"
	"math"
)

func main() {
	s := seive(10000)
	sum := 0
	count := 0

	for i := 2; i <= 10000; i++ {
		if !s[i] {
			sum += i
			count += 1
			if count == 1000 {
				break
			}
		}
	}

	fmt.Println(sum)
}

func seive(size int) []bool {
	s := make([]bool, size+1)

	s[0], s[1] = true, true

	for num, composite := range s[:int(math.Sqrt(float64(size+1)))] {
		if !composite {
			for i := num * num; i <= size; i += num {
				s[i] = true
			}
		}
	}
	return s
}
