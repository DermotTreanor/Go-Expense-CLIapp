package fileshare

import (
	"encoding/csv"
	"os"
	"sort"
	"strconv"
)

func WriteCSVFile(path string, expenseMap map[string]float64) (err error) {
	
	var expenseData []string
	headers := []string{}
	
	//Open the file and a new csv.Writer
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil{
		return err
	}
	defer file.Close()

	//Make a sorted slice of headers.
	for key := range expenseMap{
		headers = append(headers, key)
	}
	sort.Strings(headers)

	//Create the csv writer and defer a Flush call
	writer := csv.NewWriter(file)
	defer writer.Flush()

	//Get a FileInfo object and check if the array already has something in it. If not, write the headers.
	fileInfo, err := os.Stat(path)
	if err != nil{
		return err
	}

	if fileInfo.Size() == 0 {
		writer.Write(headers)
	}

	//Create a string slice of expenses in the same order that we put in our headers.
	for _, header := range headers{
		var singleExpense string = strconv.FormatFloat(expenseMap[header], 'f', 2, 64) 
		expenseData = append(expenseData, singleExpense)

	}
	//Write the slice of strings. The writer will handle all of the delimitation.
	writer.Write(expenseData)
	
	return nil
}


