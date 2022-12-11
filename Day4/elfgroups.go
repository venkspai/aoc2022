package main
 
import (
	"strconv"
	"strings"
	"log"
    "bufio"
    "fmt"
    "os"

	
)


func calcOverlap(firstElf []string, secondElf []string) int {

	//many assumptions about the format
	//no error handling 

	//7-7 8-42
	//8-42 7-7

	var overlap int = 0
	var elfFirstOne, elfFirstTwo, elfSecondOne, elfSecondTwo int = 0, 0, 0, 0
	var err error

	elfFirstOne, err = strconv.Atoi(firstElf[0])
	if err != nil {
		log.Fatal(err)
	}

	elfFirstTwo, err = strconv.Atoi(firstElf[1])
	if err != nil {
		log.Fatal(err)
	}

	elfSecondOne, err = strconv.Atoi(secondElf[0])
	if err != nil {
		log.Fatal(err)
	}

	elfSecondTwo, err = strconv.Atoi(secondElf[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Numbers:", elfFirstOne, elfFirstTwo, elfSecondOne, elfSecondTwo)

//	5 - 7 .. 7 - 11
// 7 - 7 .. 7 - 87
// 7 - 11 .. 5 .. 7

	if elfFirstOne <= elfSecondOne {
		if elfFirstTwo >= elfSecondOne {
			overlap = 1
			fmt.Println("One")
		}
	} 

	if elfSecondOne <= elfFirstOne {
			if elfSecondTwo >= elfFirstOne {
				overlap = 1
				fmt.Println("Two")
			}
		}
	
	fmt.Println(firstElf, secondElf, overlap)
	return overlap

}

func calcFullyContained(firstElf []string, secondElf []string) int {

	//many assumptions about the format
	//no error handling 

	//7-7 8-42
	//8-42 7-7

	var overlap int = 0
	var elfFirstOne, elfFirstTwo, elfSecondOne, elfSecondTwo int = 0, 0, 0, 0
	var err error

	elfFirstOne, err = strconv.Atoi(firstElf[0])
	if err != nil {
		log.Fatal(err)
	}

	elfFirstTwo, err = strconv.Atoi(firstElf[1])
	if err != nil {
		log.Fatal(err)
	}

	elfSecondOne, err = strconv.Atoi(secondElf[0])
	if err != nil {
		log.Fatal(err)
	}

	elfSecondTwo, err = strconv.Atoi(secondElf[1])
	if err != nil {
		log.Fatal(err)
	}

//	fmt.Println("Numbers:", elfFirstOne, elfFirstTwo, elfSecondOne, elfSecondTwo)

	if elfFirstOne <= elfSecondOne {
		if elfFirstTwo >= elfSecondTwo {
			overlap = 1
//			fmt.Println("One")
		}
	} 

	if elfSecondOne <= elfFirstOne {
			if elfSecondTwo >= elfFirstTwo {
				overlap = 1
//				fmt.Println("Two")
			}
		}
	
//	fmt.Println(firstElf, secondElf, overlap)
	return overlap

}




func calcElfGroups (elfGroup string) ([]string, []string) {
	//fmt.Println(elfGroup)
	res1 := strings.Split(elfGroup, ",")
	//fmt.Println("First split ", reflect.TypeOf(res1[0]))
	//fmt.Println("Second split", res1[1])

	firstElf := strings.Split(res1[0], "-")
	secondElf := strings.Split(res1[1], "-")

	return firstElf, secondElf	
}

 
func main() {

	//declare variable
	var elfGroup int = 0
	var elfGroup2 int = 0


	readFile, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening the file: ", err)
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		elfGroup += calcFullyContained(calcElfGroups(string(fileScanner.Text())))
		elfGroup2 += calcOverlap(calcElfGroups(string(fileScanner.Text())))
	}

	fmt.Println("Total number of elf groups: ", elfGroup)
	fmt.Println("Total number of elf overlaps: ", elfGroup2)

	readFile.Close()

}