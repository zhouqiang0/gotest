package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("C:/1test/matrix.txt")
	if err != nil {
		fmt.Println("..")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var row, col = 0, 0
	first, _ := reader.ReadString('\n')
	s := strings.TrimSpace(first)
	sarray := strings.Fields(s)
	row, _ = strconv.Atoi(sarray[0])
	col, _ = strconv.Atoi(sarray[1])
	num, _ := strconv.Atoi(sarray[2])

	matrix := make([][]int, row)
	for i := range matrix {
		matrix[i] = make([]int, col)
	}

	for i := 0; i < num; i++ {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		s = strings.TrimSpace(str)
		sarray = strings.Fields(s)
		row, _ = strconv.Atoi(sarray[0])
		col, _ = strconv.Atoi(sarray[1])
		num, _ = strconv.Atoi(sarray[2])

		matrix[row][col] = num
	}
	for _, v := range matrix {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
