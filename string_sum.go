package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
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
	//Check if input string is empty
	if input == "" || strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	//Check for number of operands
	var nSlice []string
	s := strings.Split(strings.TrimFunc(input, func(r rune) bool { return !unicode.IsNumber(r) && !unicode.IsLetter(r) }), "-")
	for _, v := range s {
		vs := strings.Split(strings.TrimFunc(v, func(r rune) bool { return !unicode.IsNumber(r) && !unicode.IsLetter(r) }), "+")
		for _, b := range vs {
			nSlice = append(nSlice, strings.TrimFunc(b, func(r rune) bool { return !unicode.IsNumber(r) && !unicode.IsLetter(r) }))
		}
	}
	var numSlice []string
	for _, v := range nSlice {
		if strings.ContainsAny(v, "0123456789") {
			numSlice = append(numSlice, v)
		}
	}
	if len(numSlice) != 2 {
		fmt.Println(numSlice, len(numSlice))
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	//Check if two numbers are valid
	o1, err := strconv.Atoi(numSlice[0])
	if err != nil {
		return "", fmt.Errorf("first element of slice is not valid: %w", err)
	}
	o2, err := strconv.Atoi(numSlice[1])
	if err != nil {
		return "", fmt.Errorf("second element of slice is not valid: %w", err)
	}

	//Calculate signs for each operand and summarize them
	var so1, so2 = 1, 1
	v := strings.SplitN(input, numSlice[0], 2)
	for _, vv := range v[0] {
		if string(vv) == "-" {
			so1 *= -1
		}
	}
	for _, vv := range v[1] {
		if string(vv) == "-" {
			so2 *= -1
		}
	}

	sum := o1*so1 + o2*so2
	output = strconv.Itoa(sum)

	return output, nil
}
