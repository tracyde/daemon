package main

import (
	"fmt"
	"flag"
)

var master string

func init() {
	flag.StringVar(&master, "master", "", "the master node to check into")
}

func main() {
	flag.Parse()
	fmt.Printf("Master: %s\n", master)
}
