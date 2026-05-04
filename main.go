package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Person struct {
	name string
	gender string
	civilStatus string
	age int
}

func readValid(message string, validDigits string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
			fmt.Print(message)
			input, _ := reader.ReadString('\n')
			inserted := strings.TrimSpace(input)

			valid := true
			for _, char := range inserted {
					if !strings.ContainsRune(validDigits, char) {
						valid = false
						break
					}
			}

			if valid && inserted != "" {
				return inserted
			}
			fmt.Println("Invalid input, try again.")
	}
}

func readLine(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readAge(message string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(message)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		var age int
		_, err := fmt.Sscan(input, &age)
		if err == nil && age >= 0 {
			return age
		}
		fmt.Println("Invalid age, try again.")
	}
}

func main(){
	var olderWom string
	var womAgeIndex, closingIndex int
	var marrOld30, marrMen, men, women int
	var singleMen, womSingYoung25, sumAgeMarrMen, womMaxAge int
	var avgAgeMarrMen float64

	for {
		person := Person{}

		person.age = readAge("Insert your age (0 to exit): ")

		if person.age == 0 {
			fmt.Println("Age inserted was 0, closing app...")
			break
		}

		person.name = readLine("Insert your name: ")

		person.gender = readValid("Insert your gender: ", "12")

		person.civilStatus = readValid("Insert your civil status: ", "1234")

		if person.gender == "1" {
			men++
			if person.civilStatus == "3" {
				marrMen++
				sumAgeMarrMen += person.age
			    if person.age > 30 {
					marrOld30++
				}
			}
			if person.civilStatus == "1" {
				singleMen++
			}
		} else {
			women++
			if person.civilStatus == "1" && person.age < 25 {
				womSingYoung25++
			}
			if womAgeIndex == 0 {
				womMaxAge = person.age
				olderWom = person.name
				womAgeIndex++
			} else {
				if womMaxAge < person.age {
					womMaxAge = person.age
					olderWom = person.name
				}
			}
		}
	closingIndex++
}

	if marrMen > 0 {
		avgAgeMarrMen = float64(sumAgeMarrMen) / float64(marrMen)
		fmt.Printf("The average age of the married men is: %.2f\n", avgAgeMarrMen)
	}

	if closingIndex > 0 {
		fmt.Printf("The amount of married men older than 30 is: %d\n", marrOld30)
		fmt.Printf("The amount of single women younger than 25 is: %d\n", womSingYoung25)
		fmt.Printf("The total amount of men is: %d, and the women is: %d\n", men, women)
		fmt.Printf("The amount of single men is: %d\n", singleMen)
		fmt.Printf("The oldest women is: %s\n", olderWom)	
	}	
}
