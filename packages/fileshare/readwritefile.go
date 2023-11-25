package fileshare

import (
	"math"
	"os"
	"encoding/binary"
)



func fromFloat64SliceToByteSlice( floatSlice []float64) []byte{
	var sizeForByteSlice int = len(floatSlice) * 8
	var byteSlice []byte = make([]byte, sizeForByteSlice)

	for index, v := range floatSlice{
		
		currentFloatAsBits := math.Float64bits(v)
		
		var offset int = index * 8
		binary.LittleEndian.PutUint64(byteSlice[offset: offset + 8], currentFloatAsBits)

	}

	return byteSlice

}


// WORK IN PROGRESS!!!!
func WriteFileData(path string, expenseSlice []float64)(err error){

	var dataSlice []byte = fromFloat64SliceToByteSlice(expenseSlice)

	file, err:= os.OpenFile(path, os.O_WRONLY| os.O_CREATE| os.O_APPEND, 0644)
	defer file.Close()

	_, err = file.Write(dataSlice)

	return nil
}


