package main

import (
	"fmt"
	"cli_go_program/packages/userinput"
	"cli_go_program/packages/fileshare"
)


func main() {
	var path string = "./program_files/data_from_go.txt"

	//Use the userinput package to get data from the user and return it. 
	var individualExpenses, mapOfExpenseToCategory = userinput.InputExpenses()
	var summedExpenses float64 = userinput.SumCosts(individualExpenses) 


	//Write userinput to the file
	err := fileshare.WriteFileData(path, individualExpenses)
	if err != nil{
		fmt.Println("THERE'S A WRITE ERROR!", err)
	} else{
		fmt.Println("Oh, good! There's no WRITE error!")
	}


	//Read the total file
	var fileData []float64
	fileData, err = fileshare.ReadFileData(path)
	if err != nil{
		fmt.Println("THERE'S A READ ERROR!", err)
	} else{
		fmt.Println("Oh, good! There's no READ error!")
	}

	fmt.Println(fileData)






	fmt.Print("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("Map: ", mapOfExpenseToCategory)
	fmt.Println("Expenses: ", individualExpenses)
	fmt.Printf("Your expenses for today add up to %.2f.", summedExpenses)
	fmt.Print("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")

}
