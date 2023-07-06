//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"strings"
)

func main() {
	//https://pkg.go.dev/strings#example-Contains
	s1 := "Yin"
	s2 := strings.ToLower(s1)
	fmt.Println(s2)

	fmt.Println(strings.Contains(s1, "y"))
	fmt.Println(strings.Contains(s1, "Y"))
}
