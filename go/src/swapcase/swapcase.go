package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	caseswitch := func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			return (r + 32)
		}
		if r >= 'a' && r <= 'z' {
			return (r - 32)
		}
		return r
	}
	for scanner.Scan() {
		fmt.Println(strings.Map(caseswitch, scanner.Text()))
	}
}
