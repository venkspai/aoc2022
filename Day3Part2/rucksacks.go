package main
 
import (
    "bufio"
    "fmt"
    "os"

	
)

func calcGroupPriority(elf1 []rune, elf2 []rune, elf3 []rune) int {

	elf1len := len(elf1);
	elf2len := len(elf2);
	elf3len := len(elf3);
/*
Start with first char first group
match it with chars in second group 
until you find the match
hold that
match it with chars in third group
*/

var priority rune
var priorityInt int
out:
	for i := 0; i < elf1len; i++ {
		for t := 0; t < elf2len; t++ {
			if (elf1[i] == elf2[t]) {
				for s := 0; s < elf3len; s++ {
					if (elf1[i] == elf3[s]) {
						priority = elf1[i]
						break out
					} 
				}
			}
		}
	}

	fmt.Println("the group badge is", string(priority))

	if (priority > 96) {
		priorityInt = int(priority - 96)
	} else {
		priorityInt = int(priority - 38)
	}

	return priorityInt
}
/*
func calcPriority(ruckSack []rune) int {
	fmt.Println(string(ruckSack))
	ruckSackLen := len(ruckSack)
	fmt.Println(ruckSackLen)

	var priority rune
	var priorityInt int

	for i :=0; i < ruckSackLen/2; i++ {
		for t:=(ruckSackLen/2); t < (ruckSackLen); t++ {
			if (ruckSack[i] == ruckSack[t]) {
				priority = ruckSack[i]
				break
			}
		}
	}

	if (priority > 96) {
		priorityInt = int(priority - 96)
	} else {
		priorityInt = int(priority - 38)
	}
	
	fmt.Println("Priority rune is", priority)
	fmt.Println("Priority for", string(priority), "is", priorityInt)
	fmt.Println("")
	return priorityInt
	
}
*/

func main() {
 
	//var err string
	readFile, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
  
	/*
	Read the file
	Scan the string
	read the halves into separate arrays
	compare arrays to find the common letter
	capture letter
	Break
	Find the priority value associated with that letter
	total up priority values
	*/

	var myStringSlice1 []rune
	var myStringSlice2 []rune
	var myStringSlice3 []rune
	var totalPriority int = 0

    for fileScanner.Scan() {
		myStringSlice1 = ([]rune(fileScanner.Text()))
		fileScanner.Scan()
		myStringSlice2 = ([]rune(fileScanner.Text()))
		fileScanner.Scan()
		myStringSlice3 = ([]rune(fileScanner.Text()))

		totalPriority = totalPriority + calcGroupPriority(myStringSlice1, myStringSlice2, myStringSlice3)
		fmt.Println(string(myStringSlice1))
		fmt.Println(string(myStringSlice2))
		fmt.Println(string(myStringSlice3))
		
    }

	fmt.Println("Total Priority: ", totalPriority)

	
  
    readFile.Close()


}