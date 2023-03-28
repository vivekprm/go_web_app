package main

import (
	"fmt"

	xyz "github.com/vivekprm/go_web_app/add"
)

func main () {
	a := 3
	b := 5
	res := xyz.Add(a, b)
	fmt.Printf("Sum is: %d\n", res)
}