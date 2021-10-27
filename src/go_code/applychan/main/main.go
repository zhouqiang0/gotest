package main

//读写 var aChan chan int
//只写 var aChan chan<- int
//只读 var aChan <-chan int
import (
	"fmt"
)

//两个协程，writeData,readData
func writeData(intChan chan int) {
	//使用defer + recover防止协程错误
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("发生错误%v", err)
		}
	}()

	for i := 1; i <= 2000; i++ {
		intChan <- i
		fmt.Printf("写入数据%v\n", i)
		//time.Sleep(time.Millisecond*20)
	}
	close(intChan) //写50个数据后关闭管道
}

func readData(intChan chan int, resChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		res := 0
		for i := 1; i <= v; i++ {
			res += i
		}
		resChan <- res
		fmt.Printf("读到数据：%v\n", v)
	}
	exitChan <- true //用于main函数协程序,防止快速退出
	close(exitChan)
}

func main() {
	intChan := make(chan int, 500)
	resChan := make(chan int, 2000)

	exitChan := make(chan bool, 1)

	go writeData(intChan)

	go readData(intChan, resChan, exitChan)
	go readData(intChan, resChan, exitChan)
	go readData(intChan, resChan, exitChan)
	go readData(intChan, resChan, exitChan)
	go readData(intChan, resChan, exitChan)
	go readData(intChan, resChan, exitChan)
	go readData(intChan, resChan, exitChan)
	go readData(intChan, resChan, exitChan)

	//time.Sleep(time.Second*5)//传统方法

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}
