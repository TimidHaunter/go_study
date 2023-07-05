package main

import "fmt"

func main() {
	/*
	 * 数组指针：一个指针，数组的指针，里面存储的事数组地址
	 * *[4]int | *[5]string | *[6]float32
	 *
	 * 指针数组：一个数组，数组存储的是一组指针
	 * [4]*int
	 *
	 *  *[5]float64  指针，存储5个浮点类型数据数组的指针
	 *  *[3]string   指针，存储3个字符串类型数据数组的指针
	 *  [3]*string   数组，存储3个字符串指针地址的数组
	 *  [6]*float64  数组，存储6个浮点指针地址的数组
	 *  *[5]*float64 指针，存储5个浮点指针地址的数组的指针
	 *  *[3]*string  指针，存储5个字符串指针地址的数组的指针
	 *  **[4]string  指针，*[4]string的指针，[4]string的指针，存储4个字符串类型数据数组的指针的指针
	 *  **[4]*string 指针，*[4]*string的指针，[4]*string的指针，存储4个字符串指针地址的数组的指针的指针
	 *
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

	p1[1] = 200     /* 通过数组指针操作数组元素 上面的简写(*p1)[0] */
	fmt.Println(p1) /* 数组指针 */
	fmt.Println(arr1)

	fmt.Println("=========== 指针数组 ===========") /* 指针数组，数组里面存储的元素都是指针*/

	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr2) /* [1 2 3 4] */
	fmt.Println(arr3) /* [0xc000016140 0xc000016148 0xc000016150 0xc000016158] */

	fmt.Println("================================")

	arr2[0] = 10 /* 操作数组元素 */
	fmt.Println(arr2)
	fmt.Println(a)

	*arr3[0] = 200 /* 通过数组指针去操作数组元素 */
	fmt.Println(arr3)
	fmt.Println(a)

	fmt.Println("================================")

	b = 1000
	fmt.Println(arr2) /* [10 2 3 4] */ //arr2取的是b的数值，改变b，arr2元素不改变
	fmt.Println(arr3) //arr3 存储的是变量地址，b改变，arr3元素也改变

	for i := 0; i < len(arr3); i++ {
		fmt.Println(*arr3[i])
	}
}
