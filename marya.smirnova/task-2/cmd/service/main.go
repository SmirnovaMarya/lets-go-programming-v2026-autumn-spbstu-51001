package main

import "fmt"

func main() {
	var departmentCount int

	if _, err := fmt.Scan(&departmentCount); err != nil {
		fmt.Println("Invalid departments count")
		return
	}

	for d := 0; d < departmentCount; d++ {
		var employeeCount int

		if _, err := fmt.Scan(&employeeCount); err != nil {
			fmt.Println("Invalid employees count")
			return
		}
	}
}
