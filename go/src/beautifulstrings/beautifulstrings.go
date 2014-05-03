package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func main() {
	filename := os.Args[1]

	re := regexp.MustCompile("[^a-zA-Z]")

	bytes, _ := ioutil.ReadFile(filename)
	input := string(bytes[:])
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			break
		}

		letter_counts := make(map[string]int)

		line = re.ReplaceAllString(line, "")
		line = strings.ToLower(line)

		for _, c := range line {
			letter_counts[string(c)] += 1
		}

		sum := 0
		i := 26
		for _, p := range sortMapByValue(letter_counts) {
			sum += i * p.Value
			i--
		}
		fmt.Println(sum)
	}
}
