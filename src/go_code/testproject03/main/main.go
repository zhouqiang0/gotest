package main

import (
	"fmt"
)

func Bub(arr *[]float64) {
	for i := len(*arr) - 1; i >= 1; i-- {
		for j := 0; j < i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				//temp := (*arr)[j]
				//(*arr)[j] = (*arr)[j+1]
				//(*arr)[j+1] = temp
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}
func main() {
	var asd = []float64{24.0, 34.5, 12.5, 56.5, 7.5}
	Bub(&asd)
	fmt.Println(asd)

	var slice1 = make([]map[string]int, 2)
	slice1[0] = map[string]int{
		"asd": 1,
		"qwe": 2,
	}
	slice1[1] = map[string]int{
		"asd1": 1,
		"qwe1": 2,
	}
	//slice1[2] = map[string]int{
	//	"asd2" : 1,
	//	"qwe2" : 2,
	//}
	ssda := map[string]int{
		"asd2": 1,
		"qwe2": 2,
	}
	slice1 = append(slice1, ssda)
	fmt.Println(slice1)

	//const(
	//	a = iota
	//	b = 6
	//	c = iota
	//	d = iota
	//	e
	//	f
	//)
	//fmt.Println(a,b,c,d,e,f)
}
