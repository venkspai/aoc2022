package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello world")

	//declaring variables
	var number0, number1, number2, number3 = 1, 1.0, 2.2, 0.0

	var number4 float64 = 3.23
	number4 = float64(number0)

	//short declarations
	strVar := "1000"
	var myString string
	var myBool bool

	//initializing variables
	//number1, number2 = 1, 2

	number3 = number1 + number2
	fmt.Println("Total is ", number3)
	fmt.Println("Number 0 is ", number0)

	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	//Using math and strings packages
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head first go"))

	//Printing a rune
	fmt.Println('A')

	//Printing Booleans and Integers
	fmt.Println(true)
	fmt.Println(23)

	//printing comparisons
	fmt.Println(5 == 8)
	fmt.Println(4+4 == 2)
	//Printing reflections
	fmt.Println(reflect.TypeOf(2.5))

	//Printing starting values
	fmt.Println(myString, myBool)

	//Printing conversions
	fmt.Println("Conversion ", number4)

}
