//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/**
		 * 方法：和函数一样，但是需要指定一个接收者，这个类型的值或者这个类型的指针调用
		 * func (t Type) methodName (parameter list) (return list) {}
	     *
	     * 对比函数
		 * A 意义
		 *  方法：某个类别的行为功能，需要指定的接受者调用
		 *  函数：一段独立的功能的代码，可以直接调用
		 *
		 * B 语法
		 *  方法：方法名可以相同，只要接收者不同
		 *  函数：命名不能冲突
	*/

	// 结构体值调用
	w1 := Worker{name: "HuaR", age: 22, gender: "male"}
	w1.work() /* 调用work方法 $this->work()*/

	fmt.Println("=================================")

	// 结构体指针调用
	w2 := &Worker{"LiY", 21, "male"}
	fmt.Println(w2)
	fmt.Printf("%T \n", w2)
	w2.work()

	w2.rest()

	w1.rest()

	fmt.Println("=================================")

	w2.printInfo()

	// 方法定义的结构体指针调用，但是可以直接用结构体调用，方法自动去取地址
	c1 := Cat{color: "白色", age: 6}
	c1.printInfo()
}

// 定义结构体
type Worker struct {
	name   string
	age    int
	gender string
}

type Cat struct {
	color string
	age   int
}

// 定义方法，只有Worker结构体可以调用
func (w Worker) work() { /* 无参数无返回 */
	fmt.Println(w.name, "在战斗~~~")
}

// 定义方法，结构体指针调用
func (p *Worker) rest() { /* w1结构体调用的时候，方法自动取w1的地址 */
	fmt.Println(p.name, "在嬉戏~~~")
}

// 方法名字相同，调用者不同
func (p *Worker) printInfo() {
	fmt.Printf("工人名字:%s, 工人年龄：%d, 工人性别:%s \n", p.name, p.age, p.gender)
}

// 方法名字相同，调用者不同
func (p *Cat) printInfo() {
	fmt.Printf("猫颜色:%s, 猫年龄:%d \n", p.color, p.age)
}
