package main

import (
	"fmt"
)

// 定义接口
type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

//计算长期员工工资
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//计算合同员工工资
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

/*
totalExpense 可以扩展新的员工类型，而不需要修改任何代码
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

// 接口的妙用1-
func main() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

}
