package main

import (
	"fmt"
	"cli_go_program/packages/userinput"
	"cli_go_program/packages/fileshare"
)


func main() {
	var path string
	path = "./data_from_go.txt"

	//Use the userinput package to get data from the user and return it. 
	var individualExpenses, mapOfExpenseToCategory = userinput.InputExpenses()
	var summedExpenses float64 = userinput.SumCosts(individualExpenses) 


	err := fileshare.WriteFileData(path, individualExpenses)
	if err != nil{
		fmt.Println("THERE'S AN ERROR!", err)
	} else{
		fmt.Println("Oh, good! There's no error!")
	}
	fmt.Print("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("Map: ", mapOfExpenseToCategory)
	fmt.Println("Expenses: ", individualExpenses)
	fmt.Printf("Your expenses for today add up to %.2f.", summedExpenses)
	fmt.Print("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")

}


