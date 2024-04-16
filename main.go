package main

import (
	"fmt"

	"github.com/jerberlin/generics/slice"
)

func main() {
	a := [...]int{1, 2, 3}
	s := a[0:]

	f := func(e int) int {
		return e + 1
	}

	fmt.Printf("s = %v,\nf(s) = %v\n", s, slice.MapInt(s, f))
	fmt.Printf("s = %v,\nf(s) = %v\n", s, slice.Map(s, f))
}
