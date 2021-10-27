package main

import (
	"fmt"
	"test/src/go_code/testproject04/model"
)

type Person struct {
	Name string
	age  int
}

type man struct {
	Person
}

func (p *Person) getSum(i int, j int) int {
	return i + j
}

func main() {
	var tom Person = Person{"tom", 8}
	ss := tom.getSum(1, tom.age)

	//var zhou model.student = model.student{"qiang",17.8}
	var zhou = model.NewStudent("zhou", 18.7)
	fmt.Println(zhou)
	fmt.Printf("分数：%v", zhou.GetScore())

	fmt.Println(ss)

	var xx model.Xiao
	xx.Name = "xiao"
	fmt.Println(xx)
	fmt.Printf("分数：%v\n", xx.GetScore())

	var mm man
	mm.Name = "man"
	mm.age = 123
	fmt.Println(mm)
}
