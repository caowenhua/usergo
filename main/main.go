package main

import (
	"fmt"

	"../network"
)

func main() {
	network.CListen()
	fmt.Println("End")
	network.CClose()
}
