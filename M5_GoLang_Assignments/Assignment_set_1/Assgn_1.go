package main

import "fmt"

type employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

var emps []employee

// Add Employee
func AddEmployee(name string, age int, dept string) error {
	if age >= 18 {
		id := len(emps) + 1
		emps = append(emps, employee{id, name, age, dept})
	} else {
		return fmt.Errorf("Age must be greater than 18!")
	}

	return nil
}

// search employee
func SearchEmployee(id int) (employee, error) {
	for _, emp := range emps {
		if emp.ID == id {
			return emp, nil
		}
	}

	return employee{}, fmt.Errorf("No Employee exists!!")
}

// List emps by department
func ListByDept(dept string) ([]string, error) {
	var result []string
	for _, emp := range emps {
		if emp.Department == dept {
			result = append(result, emp.Name)
		}
	}

	if len(result) == 0 {
		return []string{}, fmt.Errorf("No Employee in this Department!")
	} else {
		return result, nil
	}
}

func CountEmployees(dept string) ([]string, int, error) {
	var ans []string
	for _, emp := range emps {
		if emp.Department == dept {
			ans = append(ans, emp.Name)
		}
	}

	if len(ans) == 0 {
		return []string{}, 0, fmt.Errorf("No Employee in dept. %v", dept)
	} else {
		return ans, len(ans), nil
	}
}

func main_1() {
	AddEmployee("Rakesh", 24, "Science")
	AddEmployee("Ramesh", 22, "IT")
	AddEmployee("Rajesh", 23, "Comp. Science")
	AddEmployee("Lokesh", 28, "IT")

	emp, err := SearchEmployee(4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(emp)
	}

	ans, err := ListByDept("IT")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, emps := range ans {
			fmt.Printf("%s\n", emps)
		}
	}

	result, count, err := CountEmployees("T")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("No. of employees in given department: %v\n", count)
		fmt.Println("The Employees are: ")
		for i, emp_name := range result {
			fmt.Printf("%v: %v\n", i+1, emp_name)
		}
	}
}
