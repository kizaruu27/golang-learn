package main

import "fmt"

func main() {
	salary := 5000000
	yearsOfWork := 5
	performance := "A"

	salaryWithBonus := func(salary, bonus int) (int, int) {
		bonusSalary := salary * bonus / 100
		return salary + bonusSalary, bonusSalary
	}

	if yearsOfWork >= 5 && performance == "A" {
		finalSalary, bonusSalary := salaryWithBonus(salary, 30)
		fmt.Println("Bonus 30%:", finalSalary, "Bonus:", bonusSalary)
	} else if yearsOfWork >= 3 && performance == "B" {
		finalSalary, bonusSalary := salaryWithBonus(salary, 20)
		fmt.Println("Bonus 20%:", finalSalary, "Bonus:", bonusSalary)
	} else if yearsOfWork >= 1 && performance == "C" {
		finalSalary, bonusSalary := salaryWithBonus(salary, 10)
		fmt.Println("Bonus 10%:", finalSalary, "Bonus:", bonusSalary)
	} else {
		fmt.Println("Tidak mendapatkan bonus")
	}
}
