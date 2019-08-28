package main

import (
	"fmt"

	"github.com/sathishkumar64/adv_golang/type_exp/old"
	"github.com/sathishkumar64/adv_golang/type_exp/utils"
)

func main() {

	p := old.Person{
		Name: "Sathish",
		Age:  34,
	}
	fmt.Println(p.Name, "is major?", utils.IsMajor(p))
}
