package main

import "../network"

func main() {
	network.Listen()
	network.Close()
}
