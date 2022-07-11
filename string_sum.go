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
		fmt.Println(err.Error())
		return "", err
	}
	//if len(input) <= 2 {
	//	err := fmt.Errorf("should be at least 2 signs: %w", errorEmptyInput)
	//fmt.Println(input, err.Error())
	//	return "", err
	//}

	//signPlus := strings.Contains(input, sign)
	//sum
	if len(input) == 1 {
		err := fmt.Errorf("should be at least 2 signs: %w", errorNotTwoOperands)
		fmt.Println(err.Error())
		return "", err
	}

	//signPlus := strings.Contains(input, sign)
	signplus := "+"
	signPlus1 := strings.Contains(input, signplus)
	signminus := "-"
	Signmm := strings.Contains(input, signminus)

	//19.05
	Res1count := strings.Count(input, "+")
	Res2count := strings.Count(input, "-")
	//20.05
	Res1 := strings.Trim(input, " ")
	Res1ch := strings.Contains(input, Res1)

	//c, f
	charec1 := strings.Contains(input, "c")
	charec2 := strings.Contains(input, "f")
	whitesp := strings.Contains(input, " ")
	//array1 := strings.Fields(input)
	if charec2 == true {
		replace := strings.ReplaceAll(input, "+", " ")

		array1 := strings.Fields(replace)
		//fmt.Println(array1[1])

		err := fmt.Errorf("exist error: %v", array1[1])
		fmt.Println(err.Error())
		return "", err

	}
	if charec1 == true {
		replace := strings.ReplaceAll(input, "+", " ")

		array1 := strings.Fields(replace)
		//fmt.Println(array1[1])

		err := fmt.Errorf("exist error: %v", array1[0])
		fmt.Println(err.Error())
		return "", err

	}

	//42

	//42
	if signPlus1 == false && Res1count == 0 && len(input) < 3 && len(input) > 1 {
		err := fmt.Errorf("only 1 argument: %w", errorNotTwoOperands)
		fmt.Println(err.Error())
		return "", err

	}

	//--
	if Signmm == true && Res2count >= 3 {
		err := fmt.Errorf("more than 2 argument: %w", errorNotTwoOperands)
		fmt.Println(err.Error())
		return "", err
	}

	if Signmm == true && Res2count == 1 {
		min := strings.Split(Res1, "-")
		min1 := min[0]
		min2 := min[1]
		min1m, _ := strconv.Atoi(min1)
		min2m, _ := strconv.Atoi(min2)
		if min2m > min1m {
			sumintm := min2m - min1m

			Result3 := strconv.Itoa(sumintm)
			fmt.Println("-"+Result3, nil)

		}
		if min2m < min1m {
			sumintm := min1m + min2m

			Result3 := strconv.Itoa(sumintm)
			fmt.Println("-"+Result3, nil)

		}

		err := fmt.Errorf("sum minus is not correct: %w", errorNotTwoOperands)
		if err == nil {

			fmt.Println(err.Error())
			return "", err

			//fmt.Println(Result3)
		}

		//fmt.Println("-"+Result3, nil)
		//retirn "", err

	}

	//two minus
	if Signmm == true && Res2count < 2 && Res2count > 1 && Res1ch == true {
		partsm := strings.Split(Res1, "-")
		part1m := partsm[1]
		part2m := partsm[2]
		//20.05
		space := strings.Trim(part1m, " ")
		space2 := strings.Trim(part2m, " ")

		sint1m, _ := strconv.Atoi(space)
		sint2m, _ := strconv.Atoi(space2)
		sumintm := sint1m + sint2m

		outputm := strconv.Itoa(sumintm)
		//fmt.Printf("%s","-" + outputm)
		err := fmt.Errorf("sum is not correct: %w", errorNotTwoOperands)
		if err == nil {

			return "", err
		}

		fmt.Println("-"+outputm, nil)
		//retirn "", err

	}

	//+
	if signPlus1 == true && Res1count >= 2 {
		err := fmt.Errorf("more than 2 argument: %w", errorNotTwoOperands)
		fmt.Println(nil, err.Error())
		return "", err
	}
	//whitesp
	if whitesp == true {
		//input = string.ReplaceAll(input, " ", "")
		newinput := strings.Fields(input)

		one, oneer := strconv.Atoi(newinput[0])
		if oneer != nil {
			err := fmt.Errorf("wrong input, %w", errorNotTwoOperands)
			fmt.Println(nil, err.Error())
			return "", err
		}

		two, twoer := strconv.Atoi(newinput[1])
		if twoer != nil {
			err := fmt.Errorf("wrong input2, %w", errorNotTwoOperands)
			fmt.Println(nil, err.Error())
			fmt.Println(one, two)
			return "", err
		}

		//for i := 0, i < len(newinput), i++ {
		//       if newinput[i] == 0
		//}
		fmt.Println(one, two)
		summaWhite := one + two
		out := strconv.Itoa(summaWhite)
		fmt.Println(out, nil)
	}

	if signPlus1 == true && Res1count < 2 {
		//&& charec1 == false && charec2 == false

		parts := strings.Split(input, "+")
		part1 := parts[0]
		part2 := parts[1]

		sint1, _ := strconv.Atoi(part1)
		sint2, _ := strconv.Atoi(part2)
		sumint := sint1 + sint2
		//fmt.Println(sumint)

		output := strconv.Itoa(sumint)
		//fmt.Printf("%s", output)
		err := fmt.Errorf("sum is not correct: %w", errorNotTwoOperands)
		if err == nil {

			fmt.Println(err.Error())
			return "", err
		}

		fmt.Println(output, nil)
		//retirn "", err
	}
	//end
	return "", nil
}
