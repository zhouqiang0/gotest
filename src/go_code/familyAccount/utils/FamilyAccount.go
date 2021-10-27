package utils

import "fmt"

type Account struct {
	key     string
	balance float64
	money   float64
	note    string
	flag    bool
	detail  string
	loop    bool
}

func NewAccount() *Account {
	return &Account{
		key:     "",
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		detail:  "收支\t账户金额\t收支金额\t说明：",
		loop:    true,
	}
}

func (this *Account) showDetail() {
	fmt.Println("--------------当前收支记录")
	if this.flag {
		fmt.Println(this.detail)
	} else {
		fmt.Println("当前无收支情况")
	}
}
func (this *Account) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.flag = true
	this.detail += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
}
func (this *Account) outcome() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
		return
	}
	this.balance -= this.money
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	this.flag = true
	this.detail += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
}

func (this *Account) quit() {
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
		this.loop = false
	}
}
func (this *Account) MainMenu() {
	for {
		fmt.Println("\n--------------------记账软件------------------")
		fmt.Println("					1.收支明细")
		fmt.Println("					2.登记收入")
		fmt.Println("					3.登记支出")
		fmt.Println("					4.退出软件")
		fmt.Println("请选择（1-4）")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetail()
		case "2":
			this.income()
		case "3":
			this.outcome()
		case "4":
			this.quit()
		default:
			fmt.Println("请输入正确的数字")
		}
		if this.loop == false {
			break
		}
	}
}
