package main

import (
	"fmt"
	"time"
)

//两个协程，writeData,readData 获取8000以内的素数
func writeData(intChan chan int) {
	for i := 2; i <= 8000; i++ {
		intChan <- i
		fmt.Printf("写入数据%v\n", i)
		//time.Sleep(time.Millisecond*20)
	}
	close(intChan) //写50个数据后关闭管道
}

func readData(intChan chan int, resChan chan int, exitChan chan bool) {
	var flag bool
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			resChan <- v
			fmt.Printf("%v是素数\n", v)
		}
	}
	exitChan <- true //用于main函数协程序,防止快速退出
	//close(exitChan)  //这里不关闭管道，有可能其他协程还在执行(存入4个true)
}

func main() {
	intChan := make(chan int, 1000)
	resChan := make(chan int, 2000)

	exitChan := make(chan bool, 4)

	start := time.Now()
	go writeData(intChan)

	go readData(intChan, resChan, exitChan)
	go readData(intChan, resChan, exitChan)
	go readData(intChan, resChan, exitChan)
	go readData(intChan, resChan, exitChan)

	//time.Sleep(time.Second*5)//传统方法

	for i := 0; i < 4; i++ {
		<-exitChan
	}
	fmt.Println(time.Since(start)) //计算耗时
	close(resChan)                 //读取到4个true可以关闭resChan

	//for v := range resChan{
	//	fmt.Println(v)
	//}
}
