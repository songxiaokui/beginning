package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "he xx he xx he xxx"
	fmt.Println(strings.Replace(a, "he", "oo", -1))
	fmt.Println(strings.Replace(a, "he", "oo", 0))
	fmt.Println(strings.Replace(a, "he", "oo", 1))
	fmt.Println(strings.Replace(a, "he", "oo", 2))
}
