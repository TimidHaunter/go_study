//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * iota特殊的常量，可以被编译器自动修改的常量
	 *
	 * 每定义一个const，iota初始值为0
	 * 每定义一个常量，iota值+1
	 * 直到下一个const出现，清零
	 */
	const (
		A = iota
		B = iota
		C = iota
	)
	fmt.Println(A, B, C)

	const (
		MEAL        = iota
		FEMALE      = iota
		AGENDER     = iota
		ANDROGYNE   = iota
		ANDROGYNOUS = iota
		BIGENDER    = iota
		CIS         = iota
		//NEUTROIS 没有赋值，和上一行一致，就是iota
		NEUTROIS
		NEITHER
		YU          = "iota" //iota=9
		TRANS                //iota=10
		TRANSGENDER = iota
		TRANSSEXUAL
	)
	fmt.Println(MEAL, FEMALE, AGENDER, ANDROGYNE, ANDROGYNOUS, BIGENDER, CIS, NEUTROIS, NEITHER, YU, TRANS, TRANSGENDER, TRANSSEXUAL)
}
