package userinput

import(
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
func setExpenseMap(expenseMap map[string]float64, categoryIndex int, categoryIncrement float64) map[string]float64 {
	if len(expenseMap) == 0{
		for _, value := range SPENDING_CATEGORIES{
			expenseMap[value] = 0
		}
		return expenseMap
	} else{
		expenseMap[SPENDING_CATEGORIES[categoryIndex]] = expenseMap[SPENDING_CATEGORIES[categoryIndex]] + categoryIncrement
		return expenseMap
	}
}

func InputExpenses()([]float64, map[string]float64){
	var expenses []float64
	var currentExpenseString string
	var expenseCategoryMap =  map[string]float64{}
	var spendingCategoryReferenceString string
	var spendingCategoryReferenceValue int

	//Make the expenseCategoryMap with a 0 entry for each possible catefory in SPENDING_CATEGORIES
	expenseCategoryMap = setExpenseMap(expenseCategoryMap, -1, 0)

	fmt.Println("——————————————————————————————————————————————————————————")
input_loop:
	for{
		fmt.Print("|Enter an expense to record or press 'q' to quit:\n|Expense: ")
		fmt.Scanln(&currentExpenseString)
		if currentExpenseString == "q"{
			break
		}
		currentExpenseValue, err := strconv.ParseFloat(currentExpenseString, 64)
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println("Please enter the value again or press q to quit:")
			continue
		}
		expenses = append(expenses, currentExpenseValue)


		//Print prompt and get the expense category
		fmt.Println("|Select the number that corresponds to your expense category: ")
			for index, value := range SPENDING_CATEGORIES {
				indexIncremented := index + 1
				fmt.Print("|\t(", indexIncremented, ")", value)
				fmt.Print("\n")
			}
			

		for{
			fmt.Print("|Cateogry: ")
			fmt.Scanln(&spendingCategoryReferenceString)
			if spendingCategoryReferenceString == "q"{
				expenses = expenses[0: len(expenses) - 1]
				fmt.Println("——————————————————————————————————————————————————————————\n")
				break input_loop
			}

			spendingCategoryReferenceValue, err = strconv.Atoi(spendingCategoryReferenceString)
			if err != nil {
				fmt.Println("!\tCategory Selection Error:", err)
				fmt.Println("!\tPlease try to select a category again...")
				continue
			}

			spendingCategoryIndex := spendingCategoryReferenceValue - 1
			expenseCategoryMap = setExpenseMap(expenseCategoryMap, spendingCategoryIndex, currentExpenseValue)
			fmt.Println("——————————————————————————————————————————————————————————\n")
			break

		}

		// The upper line for the next expense:


		fmt.Println("——————————————————————————————————————————————————————————")
	}

	return expenses, expenseCategoryMap
}

func SumCosts(spendingData []float64)float64{
	var totalSpending float64 = 0
	for _, value := range spendingData{
		totalSpending += value
	}
	return totalSpending
}

