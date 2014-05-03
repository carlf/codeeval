package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

func main() {
	nums := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), ";")
		for _, num := range words {
			fmt.Printf("%s", nums[num])
		}
		fmt.Printf("\n")
	}
}
