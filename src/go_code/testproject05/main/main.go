package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type HeroSlice []Hero

type Hero struct {
	Name string
	Age  int
}

func (h HeroSlice) Len() int {
	return len(h)
}

func (h HeroSlice) Less(i, j int) bool {
	return h[i].Age < h[j].Age
}

func (h HeroSlice) Swap(i, j int) {
	tamp := h[i].Age
	h[i].Age = h[j].Age
	h[j].Age = tamp
}

func main() {
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(90)),
			Age:  rand.Intn(90),
		}
		heros = append(heros, hero)
	}
	fmt.Println("--------------排序前---------------")
	for _, info := range heros {
		fmt.Println(info)
	}
	sort.Sort(heros)
	fmt.Println("--------------排序后---------------")
	for _, info := range heros {
		fmt.Println(info)
	}
}
