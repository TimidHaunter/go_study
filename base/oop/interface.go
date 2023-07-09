//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 接口：一组方法签名，接口定义方法的声明
	 * 类型为接口中所有方法提供定义，实现接口
	 *
	 * Go语言，接口和类型的实现关系，是非嵌入的
	 *
	 * 其他语言，显示定义
	 * class Mouse implements USB{}
	 */

	// 创建Mouse
	m1 := Mouse{"雷蛇小黑"}
	fmt.Println(m1)
	// 创建FlashDisk
	f1 := FlashDisk{"金士顿8G"}
	fmt.Println(f1)

	fmt.Println("=====================================")

	//testInterface需要一个USB类型，m1实现了USB
	testInterface(m1)
	testInterface(f1)

	fmt.Println("=====================================")

	// 当需要接口类型的对象时，可以使用任意实现类对象代替
	var usb USB              /* 定义usb变量，类型是接口USB */
	fmt.Printf("%T \n", usb) /* nil */
	usb = m1
	fmt.Printf("%T \n", usb) /* main.Mouse */
	m1.start()
	m1.end()

	// 接口对象不能访问实现类中的属性
	//m1.name

	fmt.Println("=====================================")

	//m1 可以访问start()方法 end()方法 name属性
	m1.start()
	m1.end()
	fmt.Println(m1.name)

}

// 定义接口
type USB interface {
	start() /* USB开始工作 */
	end()   /* USB停止工作 */
}

// 实现类
type Mouse struct {
	name string
}

type FlashDisk struct {
	name string
}

// 实现接口方法
func (m Mouse) start() {
	fmt.Println(m.name, "鼠标,准备就绪,可以开始工作了~~~")
}

func (m Mouse) end() {
	fmt.Println(m.name, "鼠标,结束工作,可以安全退出~~~")
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "U盘,准备就绪,可以开始工作了~~~")
}

func (f FlashDisk) end() {
	fmt.Println(f.name, "U盘,结束工作,可以安全退出~~~")
}

// 测试
func testInterface(usb USB) {
	usb.start()
	usb.end()
}
