//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/**
	 * 字符串是字节的集合
	 */
	s1 := "Yin"
	s2 := ` tian`
	fmt.Println(s1 + s2)
	fmt.Println(s1[2])

	//一个汉字三个字节 utf8
	s3 := "汉中"
	fmt.Println(len(s3))
	fmt.Println(s3[0], s3[1], s3[2])

	fmt.Println("============================")

	//遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}

	fmt.Println("============================")

	for index, value := range s3 {
		fmt.Printf("%d -> %d \n", index, value)
	}

	fmt.Println("============================")

	//转换
	slice1 := []byte{230, 177, 137, 228, 184, 173}
	s4 := string(slice1)
	fmt.Println(s4)

	//字符串里面的字节不能被修改，只能整个字符串被修改
}
