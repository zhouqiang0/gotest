package main

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"test/src/go_code/quicksort/utils"
	"time"
)

var (
	start time.Time
	end   time.Duration
	id    int
)

func QuickSortInt(arr *[]int, lleft int, rright int) {

	var left, right, temp = lleft, rright, 0
	if left <= right {
		temp = (*arr)[left]
		for {
			if left != right {
				for { //从右扫描到第一个比temp小的数
					if right > left && (*arr)[right] >= temp {
						right--
					} else {
						break
					}
				}
				(*arr)[left] = (*arr)[right] //交换
				for {                        //从左扫描到第一个比temp大的数
					if left < right && (*arr)[left] <= temp {
						left++
					} else {
						break
					}
				}
				(*arr)[right] = (*arr)[left] //交换
			} else {
				break
			}
		}
		(*arr)[right] = temp //将temp换到它最终的位置
		//递归调用
		QuickSortInt(arr, lleft, left-1)
		QuickSortInt(arr, right+1, rright)
	}
}

func QuickSort(arr *[]float64, lleft int, rright int) {

	var left, right, temp = lleft, rright, 0.0
	if left <= right {
		temp = (*arr)[left]
		for {
			if left != right {
				for { //从右扫描到第一个比temp小的数
					if right > left && (*arr)[right] >= temp {
						right--
					} else {
						break
					}
				}
				(*arr)[left] = (*arr)[right] //交换
				for {                        //从左扫描到第一个比temp大的数
					if left < right && (*arr)[left] <= temp {
						left++
					} else {
						break
					}
				}
				(*arr)[right] = (*arr)[left] //交换
			} else {
				break
			}
		}
		(*arr)[right] = temp //将temp换到它最终的位置
		//递归调用
		QuickSort(arr, lleft, left-1)
		QuickSort(arr, right+1, rright)
	}
}

func toFloatList(s string) (arr []float64) {
	var num float64
	sss := strings.Fields(s)
	for i := 0; i < len(sss); i++ {

		num, _ = strconv.ParseFloat(sss[i], 64)
		arr = append(arr, num)
	}
	return arr
}

func endWithWord(e string) {
	var array1 []float64
	fmt.Println("输入排序序列(输入a结束)：")
	var num string
	for {
		fmt.Scanln(&num)
		if num == e {
			break
		}
		shu, _ := strconv.ParseFloat(num, 64)
		array1 = append(array1, shu)
	}
	//
	//array2 := array1
	//array1 = append(array1, array2...)
	fmt.Println("排序前：", array1)

	QuickSort(&array1, 0, len(array1)-1)
	fmt.Println("排序后：", array1)
}

func fileBased(src string, dest string) {
	var numString, subString string
	//打开文件
	file, err := os.Open(src)
	if err != nil {
		fmt.Println("open file err : ", err)
	}
	defer file.Close()

	//读取文件数据，进行排序
	reader := bufio.NewReader(file)
	for { //循环读每行数字，组合字符串
		string1, err := reader.ReadString('\n')
		subString = strings.Replace(string1, "\r\n", " ", -1)
		numString += subString
		if err == io.EOF {
			break
		}
	}
	nums := strings.TrimSpace(numString)
	//fmt.Printf("nums类型%T, nums : %v\n",nums, nums)

	//var array1 []int
	//var num int
	//sss := strings.Fields(nums)
	//for i := 0; i < len(sss); i++{
	//
	//	num, _ = strconv.Atoi(sss[i])
	//	array1 = append(array1, num)
	//	//fmt.Printf("nums类型%T, nums : %v\n",num, num)
	//}
	array1 := toFloatList(nums)

	start = time.Now()
	//fmt.Println("排序前：",array1)
	QuickSort(&array1, 0, len(array1)-1)
	//fmt.Println("排序后：",array1)
	end = time.Since(start)

	//写入文件
	destFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("writer err:%v", err)
	}
	defer destFile.Close()

	writer := bufio.NewWriter(destFile)

	for i, _ := range array1 {
		writer.WriteString(fmt.Sprintf("%.5f", array1[i]) + " ")
	}
	writer.Flush()

}

//随机生成float64，存入rand.txt。min-max指示范围，n指示数量
func randSort(dest string, min, max float64, n int) {
	//打开文件存放随机数
	destFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("writer err:%v", err)
	}
	defer destFile.Close()

	writer := bufio.NewWriter(destFile)
	var num float64
	for i := 0; i < n; i++ {
		num = min + rand.Float64()*(max-min)
		writer.WriteString(fmt.Sprintf("%.5f", num) + " ")
	}
	writer.Flush()

}

//生成随机数到DB
func randToDB(min, max float64, n int) {
	var num float64
	for i := 0; i < n; i++ {
		num = min + rand.Float64()*(max-min)

		num, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", num), 64)
		addRandNum(id, num)
		id++
	}

}

//从数据库读数据计算再存入数据库
func sortToDB() {
	sqlStr := "select sort_before from sortnum"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer func() {
		rows.Close() // 会释放数据库连接
	}()
	// 循环读取数据
	var nums []float64
	var num float64
	for rows.Next() {
		err := rows.Scan(&num)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		nums = append(nums, num)
	}

	QuickSort(&nums, 0, len(nums)-1)

	for id, num = range nums {
		update(id+1, num)
	}
}

//往数据库添加一个float64
func addRandNum(id int, num float64) (err error) {
	//1.写sql语句
	sqlStr := "insert into sortnum(id, sort_before,sort_after) value(?,?,0.0)"
	//2.预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常 :", err)
		return err
	}
	//3.执行
	_, err = inStmt.Exec(id, num)
	if err != nil {
		fmt.Println("执行异常 :", err)
		return err
	}
	return nil
}

//修改排序后的值
func update(id int, num float64) (err error) {
	sqlStr := "update sortnum set sort_after=? where id =?"
	//2.预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常 :", err)
		return err
	}
	//3.执行
	_, err = inStmt.Exec(num, id)
	if err != nil {
		fmt.Println("执行异常 :", err)
		return err
	}
	return nil
}

//设置监测函数，查询cpu使用率
func cpuMonitor() {
	for {
		cpuInfo, _ := cpu.Percent(time.Duration(time.Millisecond*100), false)
		memInfo, _ := mem.VirtualMemory()
		fmt.Printf("CPU使用率：%v\n 内存已使用：%v\n", cpuInfo, memInfo.UsedPercent)
	}
}

func main() {
	//启动一个协程用于监测CPU使用率
	go cpuMonitor()

	//一个一个输入，以一个参数字符结束
	//endWithWord("a")

	//以文件进行输入
	//fileBased("C:/1test/nums.txt", "C:/1test/sortNums.txt")

	////随机生成n个数(指定范围)到文件rand.txt
	//rand.Seed(time.Now().Unix())
	//randSort("C:/1test/rand.txt", -1000, 1000, 20000000)
	//
	//fileBased("C:/1test/rand.txt", "C:/1test/randsort.txt")
	////end2 := time.Since(start)
	////fmt.Println("写完文件：", end2)
	//
	//
	//fmt.Println("总共耗时： ", end)

	id = 1
	randToDB(-1000, 1000, 10)

	sortToDB()

}
