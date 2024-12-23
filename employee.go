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
	for {
		fmt.Println("\n1. Add Full-Time Employee")
		fmt.Println("2. Add Part-Time Employee")
		fmt.Println("3. List Employees")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id uint64
			var name string
			var salary uint32
			fmt.Print("Enter ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Salary: ")
			fmt.Scan(&salary)
			company.AddEmployee(fmt.Sprintf("%d", id), FullTimeEmployee{ID: id, Name: name, Salary: salary})
		case 2:
			var id uint64
			var name string
			var hourlyRate uint64
			var hoursWorked float32
			fmt.Print("Enter ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Hourly Rate: ")
			fmt.Scan(&hourlyRate)
			fmt.Print("Enter Hours Worked: ")
			fmt.Scan(&hoursWorked)
			company.AddEmployee(fmt.Sprintf("%d", id), PartTimeEmployee{ID: id, Name: name, HourlyRate: hourlyRate, HoursWorked: hoursWorked})
		case 3:
			company.ListEmployees()
		case 4:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
