package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"monster_name"` //反射机制
	Age   int    `json:"monster_age"`
	ATK   float64
	Def   float64
	Skill string
}

func testStruct() {
	monster := Monster{
		Name:  "若陀",
		Age:   3800,
		ATK:   1234.4,
		Def:   2334.6,
		Skill: "踩地板",
	}
	//序列化json
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化失败%v", err)
	}
	fmt.Printf("序列化结果：%v\n", string(data))
}

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "公子"
	a["age"] = 24
	a["address"] = "黄金屋"

	data, err := json.Marshal(&a)
	if err != nil {
		fmt.Printf("序列化失败%v", err)
	}
	fmt.Printf("序列化结果：%v\n", string(data))
}

func testSlice() {
	var slice []map[string]interface{}
	var map1 = map[string]interface{}{"name": "女士", "age": 32, "address": []string{"天守阁", "至冬"}}
	var map2 = map[string]interface{}{"name": "特瓦林", "age": 1234, "address": "风龙废墟"}
	slice = append(slice, map1)
	slice = append(slice, map2)

	data, err := json.Marshal(&slice)
	if err != nil {
		fmt.Printf("序列化失败%v", err)
	}
	fmt.Printf("序列化结果：%v\n", string(data))
}

func unmarshalStruct() {
	str := "{\"monster_name\":\"若陀\",\"monster_age\":3800,\"ATK\":1234.4,\"Def\":2334.6,\"Skill\":\"踩地板\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	fmt.Printf("monster=%v\n", monster)
}

func unmarshalMap() {
	str := "{\"address\":\"黄金屋\",\"age\":24,\"name\":\"公子\"}"

	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	fmt.Printf("monster=%v\n", a)
}

func unmarshalSlice() {
	str := "[{\"address\":[\"天守阁\",\"至冬\"],\"age\":32,\"name\":\"女士\"}," +
		"{\"address\":\"风龙废墟\",\"age\":1234,\"name\":\"特瓦林\"}]"

	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	fmt.Printf("monster=%v\n", slice)
}

func main() {
	//结构体，map，切片序列化

	testStruct()

	testMap()

	testSlice()

	//反序列化
	unmarshalStruct()

	unmarshalMap()

	unmarshalSlice()
}
