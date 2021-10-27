package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

func (s Monster) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.NumField() //获取该结构体有几个字段
	fmt.Printf("struct has %d fields\n", num)
	//遍历struct的字段
	for i := 0; i < num; i++ {
		fmt.Printf("field %d: 值为=%v\n", i, val.Field(i)) //val.Field获取值
		//获取struct标签，通过reflect.Type来获取tag的标签值
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	numMethod := val.NumMethod()
	fmt.Printf("struct has %d method\n", numMethod)

	//调用struct的方法
	val.Method(1).Call(nil) //获取第二个方法并调用（函数名字母顺序a~z）
	val.MethodByName("Print").Call(nil)

	var params []reflect.Value //声明 []reflect.value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //传入的参数为[]reflect.value，返回的也是[]reflect.value
	fmt.Println("res=", res[0].Int()) //返回结果

}

func main() {
	var a Monster = Monster{
		Name:  "北风狼",
		Age:   2000,
		Score: 89.9,
	}
	TestStruct(a)
}
