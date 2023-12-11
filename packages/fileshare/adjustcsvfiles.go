package fileshare

import (
	"encoding/csv"
	"os"
	"sort"
	"strconv"
)

func WriteCSVFile(path string, expenseMap map[string]float64) (err error) {
	
	var expenseData []string
	//Create a sorted slice of the headers
	headers := []string{}
		for key := range expenseMap{
			headers = append(headers, key)
		}
		sort.Strings(headers)
	
	//Open the file and a new csv.Writer
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil{
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	fileInfo, err := os.Stat(path)
	if err != nil{
		return err
	}

	if fileInfo.Size() == 0 {
		writer.Write(headers)
	}

	//Create the string slices of expenses in the same order as we put in our headers
	for _, header := range headers{
		var singleExpense string = strconv.FormatFloat(expenseMap[header], 'f', 2, 64) 
		expenseData = append(expenseData, singleExpense)

	}
	writer.Write(expenseData)
	
	
	
	
	
	return nil
}


