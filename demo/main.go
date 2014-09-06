package main

import (
	"fmt"

	"github.com/levicook/cycle"
)

func main() {
	colors := cycle.Strings("#fff", "#ccc")

	for i := 0; i < 10; i++ {
		fmt.Printf("Cycle %v: %q\n", i, colors.Next())
	}
}
