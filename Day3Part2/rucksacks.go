package main
 
import (
    "bufio"
    "fmt"
    "os"

	
)

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

	var myStringSlice []rune
	var totalPriority int = 0

    for fileScanner.Scan() {
		myStringSlice = ([]rune(fileScanner.Text()))
		totalPriority = totalPriority + calcPriority(myStringSlice)
		//fmt.Println(string(myStringSlice))
    }

	fmt.Println("Total Priority: ", totalPriority)

	
  
    readFile.Close()


}