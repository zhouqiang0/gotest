package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main() {
	var intChan chan int = make(chan int, 3)
	fmt.Printf("intChan:%v\n", intChan)

	//向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 12
	//intChan <- 13  往管道写数据不能超过容量

	fmt.Printf("管道长度%v 容量%v\n", len(intChan), cap(intChan))

	//从管道读取数据
	var num2 int
	num2 = <-intChan
	num3 := <-intChan

	fmt.Printf("num2:%d num3:%d\n", num2, num3)
	fmt.Printf("管道长度%v 容量%v\n", len(intChan), cap(intChan))

	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "tom"
	allChan <- Cat{"小花猫", 4}

	<-allChan
	<-allChan

	//从interface管道取出cat
	newCat := <-allChan

	fmt.Printf("newCat类型:%T, newCat：%v\n", newCat, newCat)
	//fmt.Printf("猫名字：%v",newCat.Name)  报错
	//使用类型断言
	cat1 := newCat.(Cat)
	fmt.Printf("猫名字：%v\n", cat1.Name)

	//管道的关闭与遍历
	close(intChan)
	//intChan <- 10  报错
	//关闭后只能读取数据，
	int1, ok := <-intChan
	fmt.Printf("int1:：%v 是否取出:%v\n", int1, ok)

	//使用for-range 遍历。如果管道没关闭，出现deadlock错误。
	//若管道以及关闭，则会正常遍历
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	close(intChan2) //注意关闭管道

	for v := range intChan2 {
		fmt.Printf("v=%v\n", v)

	}
}
