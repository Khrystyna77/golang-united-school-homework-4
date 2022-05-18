package string_sum

import (
	"errors"
	"fmt"
	"strconv"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace

	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	if len(input) == 0 || input == "" || input == " " {
		err := fmt.Errorf("input is empty: %w", errorEmptyInput)
		//fmt.Println(input, err.Error())
		return "", err
	}
	//if len(input) <= 2 {
	//	err := fmt.Errorf("should be at least 2 signs: %w", errorEmptyInput)
	//fmt.Println(input, err.Error())
	//	return "", err
	//}
	if len(input) == 0 || input == "" || input == " " {
		err := fmt.Errorf("input is empty: %w", errorEmptyInput)
		fmt.Println(input, err.Error())
		//return "", err
	}
	if len(input) < 2 {
		err := fmt.Errorf("should be at least 2 signs: %w", errorNotTwoOperands)
		fmt.Println(input, err.Error())
	}

	if len(input) >= 2 {
		var sum int

		for _, x := range input {
			xstr := string(x) //x is rune converted to string
			xint, _ := strconv.Atoi(xstr)
			sum += xint

		}

		output := strconv.Itoa(sum)
		//fmt.Printf("%s", output)
		err := fmt.Errorf("sum is not correct: %w", errorNotTwoOperands)
		if err == nil {

			fmt.Println(err.Error())
		}

		fmt.Println(output, nil)

	}
	return "", nil
	//end
	//	return "", nil
}
