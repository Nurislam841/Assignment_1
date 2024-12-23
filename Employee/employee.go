package main

import (
	"fmt"
)

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("FullTimeEmployee [ID: %d, Name: %s, Salary: %d Tenge]", f.ID, f.Name, f.Salary)
}

func (p PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("PartTimeEmployee [ID: %d, Name: %s, Hourly Rate: %d Tenge, Hours Worked: %.2f]", p.ID, p.Name, p.HourlyRate, p.HoursWorked)
}

type Company struct {
	employees map[string]Employee
}

func (c *Company) AddEmployee(id string, emp Employee) {
	if c.employees == nil {
		c.employees = make(map[string]Employee)
	}
	c.employees[id] = emp
	fmt.Println("Employee added successfully.")
}

func (c *Company) ListEmployee() {
	for _, emp := range c.employees {
		fmt.Println(emp.GetDetails())
	}
}

func main() {
	company := Company{employees: make(map[string]Employee)}
	company.AddEmployee("1", FullTimeEmployee{ID: 1, Name: "Yerkhan", Salary: 105000})
	company.AddEmployee("2", PartTimeEmployee{ID: 2, Name: "Nuris", HourlyRate: 170000, HoursWorked: 20.5})
	company.AddEmployee("1", FullTimeEmployee{ID: 1, Name: "Liana", Salary: 100000})
	company.AddEmployee("2", PartTimeEmployee{ID: 2, Name: "Askat", HourlyRate: 105000, HoursWorked: 18.5})
	company.ListEmployee()
}
