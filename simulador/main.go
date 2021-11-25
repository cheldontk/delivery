package main

import (
	"fmt"

	"github.com/cheldontk/delivery/route"
)

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}
	err := route.LoadPosition()
	if err != nil {
		panic("Errorrrr")
	}
	positions, _ := route.ExportJsonPositions()

	for _, v := range positions {
		fmt.Printf("%v", v)
	}

}
