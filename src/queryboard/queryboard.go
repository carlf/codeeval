package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Board struct {
	board  []int
	width  int
	height int
}

func (b Board) Set(x, y, value int) {
	b.board[y*b.width+x] = value
}

func (b Board) Get(x, y int) int {
	return b.board[y*b.width+x]
}

func (b Board) Setrow(y, value int) {
	for i := 0; i < b.width; i++ {
		b.Set(i, y, value)
	}
}

func (b Board) Setcol(x, value int) {
	for i := 0; i < b.height; i++ {
		b.Set(x, i, value)
	}
}

func (b Board) Getrow(y int) int {
	sum := 0
	for i := 0; i < b.width; i++ {
		sum += b.Get(i, y)
	}
	return sum
}

func (b Board) Getcol(x int) int {
	sum := 0
	for i := 0; i < b.height; i++ {
		sum += b.Get(x, i)
	}
	return sum
}

func NewBoard(x, y int) Board {
	b := Board{}
	b.board = make([]int, x*y)
	b.width = x
	b.height = y
	return b
}

func main() {
	filename := os.Args[1]

	b := NewBoard(256, 256)

	re := regexp.MustCompile("[a-zA-Z0-9]+")

	bytes, _ := ioutil.ReadFile(filename)
	input := string(bytes[:])
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			break
		}
		operation := re.FindAllString(line, -1)
		command := operation[0]
		arg1, _ := strconv.Atoi(operation[1])
		arg2 := 0
		if len(operation) > 2 {
			arg2, _ = strconv.Atoi(operation[2])
		}

		if command == "SetCol" {
			b.Setcol(arg1, arg2)
		} else if command == "SetRow" {
			b.Setrow(arg1, arg2)
		} else if command == "QueryCol" {
			fmt.Println(b.Getcol(arg1))
		} else if command == "QueryRow" {
			fmt.Println(b.Getrow(arg1))
		} else {
			fmt.Println("Invalid Command")
			os.Exit(1)
		}
	}
}
