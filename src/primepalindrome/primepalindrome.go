package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	s := seive(1000)

	for i := 1000; i >= 10; i-- {
		if !s[i] && is_palindrome(i) {
			fmt.Println(i)
			break
		}
	}
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

func is_palindrome(n int) bool {
	s := strconv.Itoa(n)
	len := len(s)
	s2 := make([]byte, len)
	for i := 0; i < len; i++ {
		s2[len-i-1] = s[i]
	}
	if string(s2) == s {
		return true
	}

	return false
}
