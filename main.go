package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

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

func main(){
	var name, olderWom, gender, civilStatus string
	var age, womAgeIndex, closingIndex int
	var marrOld30, marrMen, men, women int
	var singleMen, womSingYoung25, sumAgeMarrMen, womMaxAge int
	var avgAgeMarrMen float64

	womAgeIndex = 0
	closingIndex = 0

	for {
		fmt.Print("Insert your age (0 to exit): ")
		fmt.Scan(&age)

		if age == 0 {
			fmt.Println("Age inserted was 0, closing app...")
			break
		}

		fmt.Print("Insert your name: ")
		fmt.Scan(&name)

		gender = readValid("Insert your gender: ", "12")

		civilStatus = readValid("Insert your civil status: ", "1234")

		if gender == "1" {
			men++
			if civilStatus == "3" {
				marrMen++
				sumAgeMarrMen += age
			    if age > 30 {
					marrOld30++
				}
			}
			if civilStatus == "1" {
				singleMen++
			}
		} else {
			women++
			if civilStatus == "1" && age < 25 {
				womSingYoung25++
			}
			if womAgeIndex == 0 {
				womMaxAge = age
				olderWom = name
				womAgeIndex++
			} else {
				if womMaxAge < age {
					womMaxAge = age
					olderWom = name
				}
			}
		}
	closingIndex++
}

	if marrMen > 0 {
		avgAgeMarrMen = float64(sumAgeMarrMen) / float64(marrMen)
		fmt.Printf("The avgrage age of the married men is: %.2f\n", aveAgeMarrMen)
	}

	if closingIndex > 0 {
		fmt.Printf("The amount of married men older than 30 is: %d\n", marrOld30)
		fmt.Printf("The amount of single women younger than 25 is: %d\n", womSingYoung25)
		fmt.Printf("The total amount of men is: %d, and the women is: %d\n", men, women)
		fmt.Printf("The amount of single men is: %d\n", singleMen)
		fmt.Printf("The oldest women is: %s\n", olderWom)	
	}	
}
