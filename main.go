package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Input Binary: ")
	var binaryInput string
	fmt.Scanln(&binaryInput)
	err := checkInput(binaryInput)

	if err == nil {
		decimalTotal := 0
		decimalTotal = getDecimalTotal(decimalTotal, binaryInput)

		fmt.Printf("Decimal Output : %d", decimalTotal)
	} else {
		fmt.Println(err)
	}
}

func multiplyByPowerTwo(number int, pos int) int {
	for i := 0; i < pos; i++ {
		number = number * 2
	}
	return number
}

func getDecimalTotal(decimalTotal int, binaryInput string) int {
	for pos, char := range binaryInput {
		intChar := int(char - '0')
		decimalTotal += multiplyByPowerTwo(intChar, len(binaryInput)-pos-1)
	}
	return decimalTotal
}

func checkInput(byteInput string) error {

	if len(byteInput) > 8 {
		return errors.New("more than 8 binary numbers... exceeds a byte")
	}

	for _, char := range byteInput {
		if char != '0' && char != '1' {
			return errors.New("input contains characters other than 0 or 1")
		}
	}

	return nil
}
