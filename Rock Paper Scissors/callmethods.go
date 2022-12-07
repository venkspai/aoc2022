package main
 
import (

	"strings"
    "bufio"
    "fmt"
    "os"

	
)

	
	

	func comparePlays (player1 string,  player2 string) int {
		var num1, num2, playScore int = 0, 0, 0
		switch play1 := player1; play1 {
		case "A":
			num1 = 1

		case "B":
			num1 = 2

		case "C":
			num1 = 3
		}

		switch play2 := player2; play2 {
		case "X":
			num2 = 1
			if num1 == 1 {
				playScore = num2 + 3
			} else if num1 == 2 {
				playScore = num2 + 0
			} else if num1 == 3 {
				playScore = num2 + 6
			}

		case "Y":
			num2 = 2
			if num1 == 1 {
				playScore = num2 + 6
			} else if num1 == 2 {
				playScore = num2 + 3
			} else if num1 == 3 {
				playScore = num2 + 0
			}

		case "Z":
			num2 = 3
			if num1 == 1 {
				playScore = num2 + 0
			} else if num1 == 2 {
				playScore = num2 + 6
			} else if num1 == 3 {
				playScore = num2 + 3
			}
		}


		//fmt.Println(playScore)
		return playScore
	}

	func comparePlays2 (player1 string,  player2 string) int {
		var num1, num2, playScore int = 0, 0, 0
		switch play1 := player1; play1 {
		case "A":
			num1 = 1

		case "B":
			num1 = 2

		case "C":
			num1 = 3
		}

		switch play2 := player2; play2 {
		case "X":
			playScore = 0
			if num1 == 1 {
				num2 = 3
			} else if num1 == 2 {
				num2 = 1
			} else if num1 == 3 {
				num2 = 2
			}

		case "Y":
			playScore = 3
			if num1 == 1 {
				num2 = 1
			} else if num1 == 2 {
				num2 = 2
			} else if num1 == 3 {
				num2 = 3
			}

		case "Z":
			playScore = 6
			if num1 == 1 {
				num2 = 2
			} else if num1 == 2 {
				num2 = 3
			} else if num1 == 3 {
				num2 = 1
			}
		}

		playScore = playScore + num2

		//fmt.Println(playScore)
		return playScore
	}
  
 
func main() {
 
	//Declaring variables
	var score int = 0	
    readFile, err := os.Open("data.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

	/*
	Scan each line
	Parse out A and X 
	Send to a function
	Return score 
	Add Score

	Fucntion to calculate 
	num1 num2
	num1 - A (1,2,3)
	num2 - X (1,2,3)
	num1 compare num2
	calculate score
	return score 
	*/
	var mySlice []string

 

    for fileScanner.Scan() {
		mySlice = strings.Split(fileScanner.Text(), " ")
		//fmt.Println("first letter is ", mySlice[0], "second letter is ", mySlice[1])
		//fmt.Println(reflect.TypeOf(mySlice[0]))
		
		//score = score + comparePlays(mySlice[0], mySlice[1])
		//fmt.Println("Old Score: ", score)
		score = score + comparePlays2(mySlice[0], mySlice[1])
		fmt.Println("New Score: ", score)
    }


  
    readFile.Close()




}