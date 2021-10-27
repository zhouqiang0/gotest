package main

import "fmt"

//select 解决管道取数据的阻塞问题

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}
label:
	for {
		select {
		//取不到会自动匹配、配下一个case直到default
		case v := <-intChan:
			fmt.Printf("从intChan读取数据%d\n", v)
		case v := <-strChan:
			fmt.Printf("从strChan读取数据%s\n", v)
		default:
			fmt.Printf("都取不到")
			break label
		}
	}
}
