//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.txt") /* 返回值(*File, error) 文件指针对象，错误 */
	if err != nil {               /* 有错 */
		fmt.Println(err)
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println("1.Op:", ins.Op)
			fmt.Println("2.Path:", ins.Path)
			fmt.Println("3.Err:", ins.Err)
		}
		return
	}
	fmt.Println(f.Name(), "打开文件成功~")
}
