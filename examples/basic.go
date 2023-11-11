package main

import (
	"fmt"

	"github.com/AlinBorcea/mill"
)

func main() {
	fmt.Println("hei")
	mill := mill.NewMill()
	mill.PlaceMan(0, 0)
	mill.PlaceMan(0, 1)
}
