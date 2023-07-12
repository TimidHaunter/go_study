//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"math"
)

func main() {
	/**
	 * 接口断言
	 *
	 * 要知道USB到底是鼠标还是U盘
	 */
	var t1 Triangle = Triangle{3, 4, 5} /* 实例化一个结构体 */
	fmt.Println(t1.peri())
	fmt.Println(t1.area())
	fmt.Println(t1.a, t1.b, t1.c)

	fmt.Println("=============================")

	var c1 Circle = Circle{4}
	fmt.Println(c1.peri())
	fmt.Println(c1.area())
	fmt.Println(c1.radius)

	fmt.Println("=============================")

	var s1 Shape = t1 /* 声明一个接口类型，t1赋值给s1 */
	fmt.Println(s1.peri())
	fmt.Println(s1.area())

	//fmt.Println(s1.a, s1.b, s1.c) /* s1就没有三条变长 接口不能访问实现类字段 */

	fmt.Println("=============================")

	var s2 Shape = c1
	fmt.Println(s2.peri())
	fmt.Println(s2.area())

	fmt.Println("=============================")

	// 调用函数
	testShape(t1)
	testShape(c1)
	testShape(s2)

	fmt.Println("=============================")

	var s3 Shape
	getType(t1)
	getType(c1)
	getType(s1)
	getType(s2)
	getType(s3)

	fmt.Println("=============================")

	var t2 *Triangle = &Triangle{2, 3, 4}
	fmt.Printf("t2:%T,%p,%p \n", t2, &t2, t2) /* 结构体指针t2,接口类型s,ins 三个存储的都是一个地址 */
	getType(t2)                               /* 值传递，将t2值赋值给s */

	fmt.Println("=============================")

	getType2(t1)
	getType2(c1)
	getType2(s1)
	getType2(s2)
	getType2(s3)

	getType2(t2)
}

// 定义接口
type Shape interface {
	peri() float64 /* 求周长方法 */
	area() float64 /* 求面积方法 */
}

// 定义接口的实现类
type Triangle struct {
	// a float64
	// b float64
	// c float64
	a, b, c float64
}

// 三角形求周长方法
func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := t.peri() / 2
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

// 定义圆形实现类
type Circle struct {
	radius float64
}

func (c Circle) peri() float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

// 接口的使用
// 定义函数，参数是接口
func testShape(s Shape) {
	fmt.Printf("周长:%.2f,面积:%.2f \n", s.peri(), s.area())
}

// s既可以是接口类型s也可以是实现类型t c
func getType(s Shape) {
	// 但是具体是啥类型需要断言来确定 [if]
	if ins, ok := s.(Triangle); ok { /* 判断s是不是Triangle */
		//如果s是，ok=true
		fmt.Println("是三角形，三边是：", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("是圆形，半径是：", ins.radius)
	} else if ins, ok := s.(*Triangle); ok {
		fmt.Printf("ins:%T,%p,%p \n", ins, &ins, ins)
		fmt.Printf("s:%T,%p,%p \n", s, &s, s)
	} else {
		fmt.Println("我啥都不是~")
	}
}

// switch断言
func getType2(s Shape) {
	switch ins := s.(type) {
	case Triangle:
		fmt.Println("三角形,边长:", ins.a, ins.b, ins.c)
	case Circle:
		fmt.Println("圆形,半径:", ins.radius)
	case *Triangle:
		fmt.Println("三角形结构体指针,边长:", ins.a, ins.b, ins.c)
	default:
		fmt.Println("我啥都不是~")
	}
}
