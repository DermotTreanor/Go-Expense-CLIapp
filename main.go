package main

import (
	"fmt"
	"cli_go_program/packages/userinput"
)


func main() {

	var individualExpenses, mapOfExpenseToCategory = userinput.InputExpenses()
	var summedExpenses float64 = userinput.SumCosts(individualExpenses) 

	fmt.Print("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("Map: ", mapOfExpenseToCategory)
	fmt.Println("Expenses: ", individualExpenses)
	fmt.Printf("Your expenses for today add up to %.2f.", summedExpenses)
	fmt.Print("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")

}


