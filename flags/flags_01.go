package main

import (
	"flag"
	"fmt"
)

func main() {
	background := flag.String("bg", "000000", "background color")
	flag.Parse()
	fmt.Printf("The background color is %s\n", *background)
}
