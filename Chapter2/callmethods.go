package main
 
import (
    "bufio"
    "fmt"
    "os"
	"strconv"
	
)
 
func main() {
 
	//Declaring variables
	var tempTotal, number2, max = 0, 0, 0
	//var err string

    readFile, err := os.Open("data.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
    for fileScanner.Scan() {
   		intVar, _ := strconv.Atoi(fileScanner.Text())
		if intVar != 0 {
			number2 = intVar
			tempTotal = tempTotal + number2 
		} else	{
			if tempTotal > max {
				max = tempTotal
			}
			tempTotal = 0
			number2 = 0
		}
		
		fmt.Println("Max calories: ", max)

    }
  
    readFile.Close()


/*
import (
	"fmt"
	"time"
	"reflect"
	"strings"
	"bufio"
	"os"
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

	//Reading values from the standard input

	fmt.Print("Enter your grade: ")
	reader := bufio.NewReader(os.Stdin)
	os.Re
	input, _ := reader.ReadString('\n')
	fmt.Println("Your grade is ", input)
*/	

}