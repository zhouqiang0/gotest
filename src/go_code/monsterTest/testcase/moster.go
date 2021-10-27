package testcase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type monster struct {
	Name  string
	Age   int
	skill string
}

func (this *monster) Store() bool {
	//先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Printf("序列化err:%v", err)
		return false
	}

	filePath := "C:/1test/monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Printf("保存err:%v", err)
		return false
	}
	return true
}

func (this *monster) ReStore() bool {
	filePath := "C:/1test/monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("读取err:%v", err)
		return false
	}

	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Printf("反序列化err:%v", err)
		return false
	}
	return true
}
