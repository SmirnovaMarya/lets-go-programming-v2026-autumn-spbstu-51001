package main

import (
	"errors"
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30
)

func processDepartment(employeeCount int) error {
	lower := minTemp
	upper := maxTemp

	for e := 0; e < employeeCount; e++ {
		var sign string
		var value int

		if _, err := fmt.Scan(&sign, &value); err != nil {
			return errors.New("invalid employee input")
		}

		switch sign {
		case ">=":
			if value > lower {
				lower = value
			}
		case "<=":
			if value < upper {
				upper = value
			}
		default:
			return errors.New("invalid operation")
		}

		if lower > upper {
			fmt.Println(-1)
		} else {
			fmt.Println(lower)
		}
	}

	return nil
}

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
		if err := processDepartment(employeeCount); err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}
