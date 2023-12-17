package main


import (
	"fmt"
	"cli_go_program/packages/userinput"
	"cli_go_program/packages/fileshare"
)

func main() {
	var path string = "./program_files/csv_data_from_app.csv"

	//Use the userinput package to get data from the user and return it. 
	var individualExpenses, mapOfExpenseToCategory = userinput.InputExpenses()
	var summedExpenses float64 = userinput.SumCosts(individualExpenses) 


	//Write userinput to the file
	err := fileshare.WriteCSVFile(path, mapOfExpenseToCategory)
	if err != nil{
		fmt.Println("There was a WRITE error to the CSV file.\n", err)
	}

	

	fmt.Print("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("Map: ", mapOfExpenseToCategory)
	fmt.Println("Expenses: ", individualExpenses)
	fmt.Printf("Your expenses for today add up to %.2f.", summedExpenses)
	fmt.Print("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")

}
