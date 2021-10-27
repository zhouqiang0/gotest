package main

import "fmt"

type Usb interface {
	start()
	stop()
}

type phone struct {
}

func (p phone) start() {
	fmt.Println("手机接入")
}
func (p phone) stop() {
	fmt.Println("手机断开")
}
func (p phone) call() {
	fmt.Println("打电话")
}

type camera struct {
}

func (p camera) start() {
	fmt.Println("相机接入")
}
func (p camera) stop() {
	fmt.Println("相机断开")
}

type computer struct {
}

func (cp computer) working(usb Usb) {
	usb.start()
	if phone, yes := usb.(phone); yes {
		phone.call()
	}

	usb.stop()
}

func main() {
	var usbarr [3]Usb
	usbarr[0] = phone{}
	usbarr[1] = camera{}
	usbarr[2] = phone{}

	var cp1 computer

	pp := new(phone)
	fmt.Printf("类型1：%T，类型2%T\n", usbarr[0], pp)
	cp1.working(pp)
	fmt.Println()
	for _, v := range usbarr {
		cp1.working(v)
		fmt.Println()
	}
}
