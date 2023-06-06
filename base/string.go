//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	var s1 string = "Hanunt"
	fmt.Printf("%T, %s\n", s1, s1)

	var s2 string = `Hanunt`
	fmt.Printf("%T, %s\n", s2, s2)

	//单引号是int型
	v1 := 'A' //int32 65 ASCII码
	v2 := "B"
	fmt.Printf("%T, %d\n", v1, v1)
	fmt.Printf("%T, %s\n", v2, v2)

	//单个字符可以用单引号，字符串必须用双引号
	//illegal rune literal
	// var1 := 'yintian'
	// fmt.Printf("%T, %d\n", var1, var1)

	//转义\"
}
