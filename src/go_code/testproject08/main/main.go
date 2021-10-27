package main

//MPG模式 M：操作系统主线程。P：协程执行的上下文。G:协程
import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	lock  sync.Mutex
)

//计算阶乘
func test2(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test hello" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {
	////开启协程goroutine(轻量级线程)
	//go test()
	//
	////主线程。物理线程，作用于cpu,重量级
	////
	//for i := 1; i <= 10 ; i++ {
	//	fmt.Println("main hello" + strconv.Itoa(i))
	//	time.Sleep(time.Second)
	//}
	////主线程退出了，即使协程还未执行完毕，也会退出
	//cpuNum := runtime.NumCPU()
	//fmt.Println(cpuNum)

	//开启多个协助程执行计算阶乘
	for i := 1; i <= 200; i++ {
		go test2(i)
	}
	//休眠10秒
	time.Sleep(time.Second * 20)

	//遍历结果
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}

}
