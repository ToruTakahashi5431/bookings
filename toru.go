package main

import (
	"fmt"
	"log"
)

func main() {
	logger := log.Default()
	fmt.Printf("%v\n", logger)
	fmt.Printf("%#v\n", logger)
}
