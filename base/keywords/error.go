//go:build ignore
// +build ignore

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	/**
	 * 错误和异常
	 */

	// 1.这些包错误信息都已经写好了，我们只有会处理报错就行
	f, err := os.Open("test.txt") /* 返回值(*File, error) 文件指针对象，错误 */
	if err != nil {               /* 有错 */
		// log.Fatal(err)
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "打开文件成功~")

	fmt.Println("===============================")

	// 2.自定义的函数怎么返回错误信息

	/* error接口定义
	 * type error interface {
	 *     Error() string
	 * }
	 */
	err1 := errors.New("自定义错误")
	fmt.Println(err1)
	fmt.Printf("%T \n", err1) /* *errors.errorString */

	fmt.Println("===============================")

	err2 := checkAge(-1)
	if err2 != nil {
		// 有错误
		fmt.Println(err2)
		return
	}
	fmt.Println("没问题，程序继续")
}

func checkAge(age int) error {
	if age < 0 {
		return errors.New("年龄不正常")
	}
	fmt.Println("年龄正常，是：", age)
	return nil
}
