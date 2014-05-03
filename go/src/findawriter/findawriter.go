package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, "| ")
		letters := strings.Split(data[0], "")
		indices := strings.Split(data[1], " ")
		for _, n := range indices {
			x, _ := strconv.Atoi(n)
			fmt.Printf("%s", letters[x-1])
		}
		fmt.Printf("\n")
	}
}
