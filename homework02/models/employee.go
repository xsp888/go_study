package models

import "fmt"

// 员工
type Employee struct {
	// 基本信息
	Person Person
	// 工号
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工的姓名：%s,年龄：%d,工号：%d\n",
		e.Person.Name, e.Person.Age, e.EmployeeID)
}
