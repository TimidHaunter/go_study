//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 变量A存储156，地址0xc0000160a8
	 * 变量B存储0xc0000160a8，变量A的地址，就说【变量B是指针，指向变量A的指针】
	 */
	a := 156
	fmt.Printf("A数值是：%d \n", a)
	fmt.Printf("A类型是：%T \n", a)
	fmt.Printf("A地址是：%p \n", &a)

	fmt.Println("===========================================")

	// 声明指针
	var ip *int     /* 指向整型 */
	var fp *float32 /* 指向浮点型 */
	ip = &a         /* ip取的A的地址 */

	fmt.Println("ip数值是：", ip)
	fmt.Printf("ip自己的地址是：%p \n", &ip)
	fmt.Println("ip的数值是A的地址，A地址存储的数据是：", *ip) /* 获取指针存储地址，指向的变量数值 */
	fmt.Println("===========================================")
	fmt.Printf("fp数值是：%p \n", fp) /* 空指针 */

	fmt.Println("===========================================")

	//改变变量的值，不会修改其地址
	a = 1560
	fmt.Printf("A数值是：%d \n", a)
	fmt.Printf("A地址是：%p \n", &a)

	fmt.Println("===========================================")

	//通过指针修改变量的值
	*ip = 2340
	fmt.Println(a)

	fmt.Println("===========================================")

	//指针的指针
	var ip2 **int
	fmt.Println(ip2)
	//ip2是一个指针，存储一个int类型的地址
	ip2 = &ip

	fmt.Printf("%T,%T,%T \n", a, ip, ip2) /* int,*int,**int */

	fmt.Println("ip2数值：", ip2)      /* ip2数值是ip的地址 */
	fmt.Println("ip2地址：", &ip2)     /* ip2地址是0xc00000a038 */
	fmt.Println("*ip2数值是：", *ip2)   /* *ip2数值是ip的数值 */
	fmt.Println("**ip2数值是：", **ip2) /* **ip2数值是ip数值对应的值 */
}
