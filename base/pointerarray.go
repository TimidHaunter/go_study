package main

import "fmt"

func main() {
	/*
	 * 数组指针：一个指针，数组的地址
	 *
	 * 指针数组：一个数组，数组存储的是一组指针
	 */

	//1.创建数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)

	//2.创建指针，存储数组地址，数组指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)                 /* &[1 2 3 4] */
	fmt.Printf("p1数据数值：%p \n", p1)  //数组arr1的地址
	fmt.Printf("p1数据地址：%p \n", &p1) //p1指针自己的地址
	fmt.Printf("p1数据类型：%T \n", p1)

	fmt.Println("================================")

	//3.根据数组指针操作数组
	(*p1)[0] = 10   /* (指针) 需要用括号 */
	fmt.Println(p1) /* 数组指针 */
	fmt.Println(arr1)

	fmt.Println("================================")

	p1[1] = 200     /* 通过指针操作数组元素 */
	fmt.Println(p1) /* 数组指针 */
	fmt.Println(arr1)
}
