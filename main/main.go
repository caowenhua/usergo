package main

import (
	"fmt"

	"me.user/network"
)

func main() {
	network.CListen()
	fmt.Println("End")
	network.CClose()
}
