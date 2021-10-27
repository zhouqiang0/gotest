package main

import (
	"fmt"
	"test/src/go_code/customerManager/model"
	"test/src/go_code/customerManager/service"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

//调用service里的getlist方法获取客户列表
func (this *customerView) customerList() {
	customers := this.customerService.GetList()
	fmt.Println("----------------客户列表----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].ShowInfo())
	}
	fmt.Printf("\n---------------客户列表完成---------------\n")
}

//调用service里的add方法添加客户
func (this *customerView) add() {
	fmt.Println("-------------添加客户--------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("-------------添加完成--------------")
	} else {
		fmt.Println("-------------添加失败--------------")
	}
}

//调用service里的方法删除客户
func (this *customerView) delete() {
	fmt.Println("-----------------删除客户-------------")
	fmt.Println("请输入待删除客户的编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("是否确认删除（y/n）：")
	choice := ""
	fmt.Scanln(&choice)
	for {
		if choice == "y" {
			if this.customerService.Delete(id) {
				fmt.Println("---------删除成功---------")
			} else {
				fmt.Println("---------输入的id号不存在--------")
			}
			break
		} else if choice == "n" {
			return
		} else {
			fmt.Println("请输入正确的字符")
		}
		this.customerService.Delete(id)

	}
}

//退出软件
func (this *customerView) exit() {
	fmt.Println("是否退出（y/n）:")
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "n" {
			break
		}
		fmt.Println("输入有误")
	}
	if this.key == "y" {
		this.loop = false
	}
}

//调用service里的方法修改客户信息
func (this *customerView) update() {
	fmt.Println("----------------------修改客户-------------")
	fmt.Println("请输入待修改客户的编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("-----------------输入回车表示不修改本项-------------")
	fmt.Println("姓名:")
	customer := this.customerService.GetList()[this.customerService.FindById(id)]
	name := customer.Name
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := customer.Gender
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := customer.Age
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := customer.Phone
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := customer.Email
	fmt.Scanln(&email)
	this.customerService.Update(id, name, gender, age, phone, email)
	fmt.Println("-----------------修改完成-------------")

}

func (this *customerView) mainMenu() {
	for {
		fmt.Println("\n----------------客户信息管理软件------------------")
		fmt.Println("					1.添加客户")
		fmt.Println("					2.修改客户")
		fmt.Println("					3.删除客户")
		fmt.Println("					4.客户列表")
		fmt.Println("					5.退出")
		fmt.Println("请选择（1-5）:")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.customerList()
		case "5":
			this.exit()
		default:
			fmt.Println("请输入正确的数字")
		}
		if !this.loop {
			break
		}
	}
}

func main() {
	fmt.Println("ok")
	cv1 := customerView{"", true, service.NewCustomerService()}
	cv1.mainMenu()
}
