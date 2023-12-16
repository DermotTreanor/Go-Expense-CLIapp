package fileshare

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
)

func fromFloat64SliceToByteSlice(floatSlice []float64) []byte {
	var sizeForByteSlice int = len(floatSlice) * 8
	var byteSlice []byte = make([]byte, sizeForByteSlice)

	for index, v := range floatSlice {

		currentFloatAsBits := math.Float64bits(v)

		var offset int = index * 8
		binary.LittleEndian.PutUint64(byteSlice[offset:offset+8], currentFloatAsBits)

	}

	return byteSlice
}
func fromByteSliceToFloat64Slice(byteSlice []byte) (floatSlice []float64){
	if len(byteSlice)%8 != 0 {
		fmt.Println("The slice of bytes is invalid. Did not divide by 8 evenly.")
		return nil
	}
	//Make appropriately sized slice of floats (one-eight the slice of bytes)
	floatSlice = make([]float64, len(byteSlice)/8)

	var offset int = 0;
	for i := 0; i < len(floatSlice); i++{
		floatAsByteSlice := byteSlice[offset: offset + 8]

		floatAsUint64 := binary.LittleEndian.Uint64(floatAsByteSlice)
		currentFloat := math.Float64frombits(floatAsUint64)
		floatSlice[i] = currentFloat
		offset+=8
		
		
	}


	return floatSlice;
}





func WritePlainFileFloatData(path string, expenseSlice []float64) (err error) {

	var dataSlice []byte = fromFloat64SliceToByteSlice(expenseSlice)
	fmt.Println(string(dataSlice))

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	_, err = file.Write(dataSlice)

	return nil
}





func ReadPlainFileFloatData(path string) (floatSlice []float64, err error) {

	//When this dataBuffer is set to have a length of 0 with a capacity of 1024 I think we get infinite loop
	var dataBuffer []byte = make([]byte, 10)
	var data []byte


	//Open the file
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	//Read the file to a slice of bytes
	for {
		bytesRead, err := file.Read(dataBuffer)

		data = append(data, dataBuffer[:bytesRead]...)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
				return nil, err
			}
			break
		}

	}

	floatSlice = fromByteSliceToFloat64Slice(data)

	return floatSlice, err
}
