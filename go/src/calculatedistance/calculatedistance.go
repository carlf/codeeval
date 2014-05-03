package main

import "fmt"
import "log"
import "bufio"
import "os"
import "regexp"
import "strconv"
import "math"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	re, _ := regexp.Compile(`\((-?[0-9]{1,3}), (-?[0-9]{1,3})\) \((-?[0-9]{1,3}), (-?[0-9]{1,3})\)`)
	for scanner.Scan() {
		substrings := re.FindStringSubmatch(scanner.Text())
		x1, _ := strconv.Atoi(substrings[1])
		y1, _ := strconv.Atoi(substrings[2])
		x2, _ := strconv.Atoi(substrings[3])
		y2, _ := strconv.Atoi(substrings[4])
		xd := x2 - x1
		yd := y2 - y1
		distance := math.Sqrt(float64(xd*xd + yd*yd))
		fmt.Println(distance)
	}
}
