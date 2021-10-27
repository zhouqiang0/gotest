package main

import (
	"fmt"
	"testing"
)

//测试文件命名为xxx_test.go,方法命名为TestXxx()
//Testing框架调用所有上述方法。
func TestAddUpper(t *testing.T) {
	//调用
	res := addUpper(10)
	if res != 10 {
		t.Fatalf("执行错误，期望值=%v 实际值：%v\n", 55, res)
	}
	//如果正确
	t.Logf("addUpper()执行正确...")
}

func TestHello(t *testing.T) {
	fmt.Println("TestHello()被调用..")
}
