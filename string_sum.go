package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
		//fmt.Println(err.Error())
		return "", err
	}
	//if len(input) <= 2 {
	//	err := fmt.Errorf("should be at least 2 signs: %w", errorEmptyInput)
	//fmt.Println(input, err.Error())
	//	return "", err
	//}

	//sum
	if len(input) == 1 {
		err := fmt.Errorf("should be at least 2 signs: %w", errorNotTwoOperands)
		fmt.Println(nil, err.Error())
	}

	//signPlus := strings.Contains(input, sign)
	Res1count := strings.Count(input, "+")

	if Res1count >= 2 {
		//fmt.Println(input)
		err := fmt.Errorf("more than 2 argument:: %w", errorEmptyInput)
		fmt.Println(err.Error())

	}

	//sum
	//var sum int
	if Res1count < 2 {
		parts := strings.Split(input, "+")
		part1 := parts[0]
		part2 := parts[1]
		sint1, _ := strconv.Atoi(part1)
		sint2, _ := strconv.Atoi(part2)
		//sum
		sumint := sint1 + sint2
		//fmt.Println(sumint)

		result := strconv.Itoa(sumint)
		//fmt.Printf("%s", output)
		err := fmt.Errorf("sum is not correct: %w", errorNotTwoOperands)
		if err == nil {

			fmt.Println(err.Error())
		}

		fmt.Println(result, nil)

	}
	//end
	return "", nil
}
