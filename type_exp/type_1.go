package main

import "fmt"

type myName string

func (n myName) Name() string {
	return string(n)
}

func main() {

	name := myName("Sathish")
	fmt.Printf("%v %T\n", name, name)

	n := myName.Name("Sathish")
	fmt.Printf("%v %T\n", n, n)
}
