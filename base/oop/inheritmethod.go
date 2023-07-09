package main

import "fmt"

func main() {
	/**
	 * 继承中的方法，方法的继承
	 *
	 * 1.子类可以直接访问父类的方法和属性
	 * 2.子类可以新增自己的属性和方法
	 * 3.子类可以重写父类的方法(override，将父类已有的方法，重新实现)

	 */

	// 创建父类Person类型
	p1 := Person{name: "刘备", age: 56}
	fmt.Println(p1.name, p1.age) /* 父类对象，访问父类字段属性 */
	p1.eat()                     /* 父类对象，访问父类方法 */

	fmt.Println("================================")

	// 创建子类Student类型
	s1 := Student{Person{"刘禅", 20}, "蜀汉初级中学"}
	fmt.Println(s1.name)   /* s1.Person.name */
	fmt.Println(s1.age)    /* 子类对象，可以直接访问父类字段(提升字段) */
	fmt.Println(s1.school) /* 子类对象，访问自己新增字段 */

	s1.eat()   /* 子类对象访问父类方法，重写后子类对象访问自己重写的方法 */
	s1.study() /* 子类对象访问子类新增方法 */
}

// 定义"父类"
type Person struct {
	name string
	age  int
}

// 定义"子类"
type Student struct {
	Person
	school string
}

// 定义父类方法
func (p Person) eat() {
	fmt.Println("父类的方法,吃窝窝头~")
}

// 定义子类方法
func (s Student) study() {
	fmt.Println("子类新增的方法,好好学习~")
}

// 重写父类方法eat()
func (s Student) eat() {
	fmt.Println("子类重写父类的方法,吃炸鸡喝可乐~")
}
