package main

import (
	"testing"
)

//测试文件命名为xxx_test.go,方法命名为TestXxx()
//Testing框架调用所有上述方法。
func TestGetSub(t *testing.T) {
	//调用
	res := getSub(10, 3)
	if res != 7 {
		t.Fatalf("执行错误，期望值=%v 实际值：%v\n", 55, res)
	}
	//如果正确
	t.Logf("getSub()执行正确...")
}
