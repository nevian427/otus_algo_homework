package main

import (
	"fmt"
	"strings"
)

const (
	widthX = 25
	widthY = 25
)

func main() {
	for y := 0; y < widthY; y++ {
		row := strings.Builder{}
		for x := 0; x <= widthX; x++ {
			col := " ."
			if x*y < 90 {
				col = " #"
			}
			row.WriteString(col)
		}
		fmt.Println(row.String())
	}
}
