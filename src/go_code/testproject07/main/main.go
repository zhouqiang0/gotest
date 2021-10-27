package main

import "fmt"

func main() {
	key := ""
	balance := 10000.0
	money := 0.0
	note := ""
	flag := false
	detail := "收支\t账户金额\t收支金额\t说明："
loop:
	for {
		fmt.Println("\n--------------------记账软件------------------")
		fmt.Println("					1.收支明细")
		fmt.Println("					2.登记收入")
		fmt.Println("					3.登记支出")
		fmt.Println("					4.退出软件")
		fmt.Println("请选择（1-4）")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("--------------当前收支记录")
			if flag {
				fmt.Println(detail)
			} else {
				fmt.Println("当前无收支情况")
			}

		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			flag = true
			detail += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			balance -= money
			fmt.Println("本次支出说明:")
			fmt.Scanln(&note)
			flag = true
			detail += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)

		case "4":
			fmt.Println("是否确定退出？y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("输入有误！")
			}
			if choice == "y" {
				break loop
			}
		default:
			fmt.Println("请输入正确的数字")
		}
	}
}
