package main

import (
	"fmt"
	"hw_1/employee"
	"hw_1/employee/developer"
)

func main() {
	var dev employee.Employee = &developer.Developer{}
	fmt.Println(dev.GetSalary())
	dev.SetSalary(800)
	fmt.Println(dev.GetSalary())
}
