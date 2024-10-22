package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name     string `json:"name"` // Capitalized for JSON export
	Age      int    `json:"age"`
	IsRemote bool   `json:"isRemote"`
	Address
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

func (a Address) printAddress() {
	fmt.Printf("Address: %s, %s\n", a.Street, a.City)
}

func (e *Employee) updateName(newName string) { // Use pointer receiver to update the name
	e.Name = newName
	fmt.Println("Updated Name:", e.Name)
}

func main() {
	address := Address{
		Street: "123 Main Street",
		City:   "New York",
	}

	employee1 := Employee{
		Name:     "LEK",
		Age:      30,
		IsRemote: true,
		Address:  address,
	}

	employee1.printAddress()    // Prints the employee's address
	employee1.updateName("Bob") // Updates the employee's name

	fmt.Println("Employee Name:", employee1.Name)
	fmt.Println("Employee Age:", employee1.Age)

	job := struct {
		Title  string
		Salary int
	}{
		Title:  "Software Engineer",
		Salary: 10000,
	}
	fmt.Println("Job Title:", job.Title)

	// Marshal the employee struct into JSON
	jsonData, err := json.Marshal(employee1)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
	} else {
		fmt.Println("Employee JSON:", string(jsonData))
	}
}
