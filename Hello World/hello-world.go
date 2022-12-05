package main

import (
	"reflect"
	"fmt"
    "strconv"
    "math"
    "strings"
)
      
func main() {
    fmt.Println("hello world")

    //declaring variables
    var number1, number2, number3 int
    strVar := "1000"

    //initializing variables
    number1 = 1
    number2 = 3

    number3 = number1 + number2
    fmt.Println("Total is ", number3)

    
    intVar, err := strconv.Atoi(strVar)
    fmt.Println(intVar, err, reflect.TypeOf(intVar))


    //Using math and strings packages 
    fmt.Println(math.Floor(2.75))
    fmt.Println(strings.Title("head first go"))

    //Printing a rune
    fmt.Println('A')

    //Printing Booleans and Integers
    fmt.Println(true);
    fmt.Println(23)

  }
