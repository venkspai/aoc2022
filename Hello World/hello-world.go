package main

import (
	"reflect"
	"fmt"
    "strconv"
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

  }
