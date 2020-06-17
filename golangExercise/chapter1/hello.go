package main

import (
	"fmt"
)

func main() {
	var v3 *int
	v4 := new(int)

	fmt.Println(v3)
	fmt.Println(&v3)
	fmt.Println(v4)
	fmt.Println(&v4)
}
