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
		nums := make([]float64, 0)
		num_strings := strings.Split(line, " ")
		for _, n := range num_strings {
			num, _ := strconv.ParseFloat(n, 64)
			nums = append(nums, num)
		}
		nums = sort(nums)
		num_strings = make([]string, 0)
		for _, n := range nums {
			num_strings = append(num_strings, fmt.Sprintf("%.3f", n))
		}
		fmt.Println(strings.Join(num_strings, " "))
	}
}

func sort(nums []float64) []float64 {
	for i := 1; i < len(nums); i++ {
		x := nums[i]
		j := i
		for j > 0 && nums[j-1] > x {
			nums[j] = nums[j-1]
			j -= 1
		}
		nums[j] = x
	}
	return nums
}
