//go:build ignore
// +build ignore

package main

import "fmt"

type Person struct {
	name string
}

func (p Person) show() {
	fmt.Println("Person =>", p.name)
}

// 类型别名
type People = Person

type Student struct {
	Person
	People
}

func main() {
	var s Student
	// s.name = "刘秀"
	s.Person.name = "刘秀"
	s.Person.show()

}
