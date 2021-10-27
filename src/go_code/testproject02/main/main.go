package main

import (
	f "fmt"
	b "test/src/go_code/testproject01/apk"
	a "test/src/go_code/testproject02"
)

func main() {
	f.Println("这是main函数")
	Fff()
	a.CFF()
	name := b.ReName("zhou")
	f.Println(name)

	//var qwe, asd int = 10, 100
	//f.Println("输入qwe：" )
	//f.Scanln(&qwe)
	//f.Println("输入asd：" )
	//f.Scanln(&asd)
	//f.Println(qwe)
	//f.Println(1)
	//f.Println(asd)
}
