package main

import (
	"fmt"
	"strconv"
)

var SPENDING_CATEGORIES []string = []string{
	"Food",
	"Shopping",
	"Travel",
	"Utility",
	"Other",
}

func setExpenseMap(expenseMap map[string]float64) map[string]float64 {
	if len(expenseMap) == 0{
		for _, value := range SPENDING_CATEGORIES{
			expenseMap[value] = 0
		}
		return expenseMap
	} else{
		return expenseMap
	}
}


func inputExpenses ()[]float64{
	var expenses []float64
	var currentExpenseString string
	var expenseCategoryMap =  map[string]float64{}

	expenseCategoryMap = setExpenseMap(expenseCategoryMap)

	fmt.Println("What expenses have you incurred so far today? [q to quit]")
	for currentExpenseString != "q"{
		fmt.Scanln(&currentExpenseString)
		if currentExpenseString != "q"{
			currentExpenseValue, err := strconv.ParseFloat(currentExpenseString, 64)

			if err != nil {
				fmt.Println("Error: ", err)
				fmt.Println("Please enter the value again or press q to quit:")
				continue
			}
			expenses = append(expenses, currentExpenseValue)
		}

		fmt.Println("----------------------------------------------------------")

	}

	return expenses
}
func sumCosts(spendingData []float64)float64{
	var totalSpending float64 = 0
	for _, value := range spendingData{
		totalSpending += value
	}
	return totalSpending
}

func main() {

	var individualExpenses []float64 = inputExpenses()
	var summedExpenses float64 = sumCosts(individualExpenses) 

	
	fmt.Println("Expenses: ", individualExpenses)
	fmt.Printf("Your expenses for today add up to %.2f.\n", summedExpenses)

}


