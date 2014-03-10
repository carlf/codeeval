package main

import (
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	fi, _ := os.Stat(filename)
	fmt.Println(fi.Size())
}
