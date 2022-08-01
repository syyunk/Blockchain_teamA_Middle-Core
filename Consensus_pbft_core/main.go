package main

import (
	"os"
)

func main() {
	nodeID := os.Args[1] /// cmd] pbft 5000 [enter]
	server := NewServer(nodeID)

	server.Start()
}
