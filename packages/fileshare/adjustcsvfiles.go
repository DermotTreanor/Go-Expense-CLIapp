package fileshare

import (
	"encoding/csv"
	"os"
	"sort"
)

func WriteCSVFile(path string, expenseMap map[string]float64) (err error) {
	
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
		header := []string{}
		for key := range expenseMap{
			header = append(header, key)
		}
		sort.Strings(header)
		writer.Write(header)
	}



	return nil
}


