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

func (e FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Full-Time Employee - ID: %d, Name: %s, Salary: %d", e.ID, e.Name, e.Salary)
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (e PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Part-Time Employee - ID: %d, Name: %s, Hourly Rate: %d, Hours Worked: %.2f", e.ID, e.Name, e.HourlyRate, e.HoursWorked)
}

type Company struct {
	Employees map[string]Employee
}

func (c *Company) AddEmployee(id string, emp Employee) {
	if _, exists := c.Employees[id]; exists {
		fmt.Printf("Error: Employee with ID %s already exists!\n", id)
		return
	}
	c.Employees[id] = emp
	fmt.Println("Employee added successfully!")
}

func (c *Company) ListEmployees() {
	fmt.Println("Employee List:")
	for id, emp := range c.Employees {
		fmt.Printf("ID: %s, Details: %s\n", id, emp.GetDetails())
	}
}

func main() {
	company := Company{Employees: make(map[string]Employee)}
	company.AddEmployee("1", FullTimeEmployee{ID: 1, Name: "Yerkhan", Salary: 105000})
	company.AddEmployee("2", PartTimeEmployee{ID: 2, Name: "Nuris", HourlyRate: 170000, HoursWorked: 20.5})
	company.AddEmployee("3", FullTimeEmployee{ID: 3, Name: "Liana", Salary: 100000})
	company.AddEmployee("4", PartTimeEmployee{ID: 4, Name: "Askat", HourlyRate: 105000, HoursWorked: 18.5})
	company.ListEmployees()
}
