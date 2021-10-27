package service

import "test/src/go_code/customerManager/model"

type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 10,
		"112", "zs@zzu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) GetList() []model.Customer {
	return this.customers
}

//添加客户
func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

//查找客户
func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}

//删除客户
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

//修改客户
func (this *CustomerService) Update(id int, name string, gender string, age int, phone string, email string) {
	index := this.FindById(id)
	this.customers[index].Name = name
	this.customers[index].Gender = gender
	this.customers[index].Age = age
	this.customers[index].Phone = phone
	this.customers[index].Email = email
}
