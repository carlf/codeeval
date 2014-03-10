package main

import (
	"bytes"
	"fmt"
)

func main() {
	for i := 1; i <= 12; i++ {
		var line bytes.Buffer
		for j := 1; j <= 12; j++ {
			line.WriteString(fmt.Sprintf("%4d", i*j))
		}
		fmt.Println(line.String())
	}
}
