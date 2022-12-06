package main

import (
	"fmt"
	"time"
	"reflect"
	"strings"
)

func main() {

	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)
	fmt.Println(reflect.TypeOf(now))

	//Replacer
	replacer := strings.NewReplacer("#", "o")
	broken := "G# R#cks!"
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
	

}