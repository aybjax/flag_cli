package main

import (
	"flag"
	"fmt"
)

var name *string = flag.String("name", "default", "help text")
var num1 *int = flag.Int("num1", 0, "help text")
var num2 int

func main() {
	flag.Parse()
	fmt.Printf("Hello %s, whasuuuuuup! Here is num1: %d and num2: %d\n", *name, *num1, num2)
}

func init() {
	flag.IntVar(&num2, "num2", 0, "Another way of creation")
}