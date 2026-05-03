package main

import "fmt"
import "os"

func main(){
	var name, olderWom string
	var gender, civilStatus, age, ageIndex int
	var marrOld30, marrMen, men, women int
	var singleMen, womSingYoung25, sumAgeMarrMen, womMaxAge int
	var womAgeIndex int
	var aveAgeMarrMen float64

	womAgeIndex = 0
	ageIndex = 0

	fmt.Print("Insert your age: ")
	fmt.Scan(&age)
	if age == 0 {
		fmt.Println("Age inserted was 0, closing app..")
		os.Exit(0)
	}

	for age > 0 {
		if ageIndex >= 1 {
			fmt.Print("Insert your age: ")
			fmt.Scan(&age)
		}

		if age == 0 {
			fmt.Println("Age inserted was 0, closing app...")
			break
		}

		fmt.Print("Insert your name: ")
		fmt.Scan(&name)
	
		fmt.Print("Insert your gender: ")
		fmt.Scan(&gender)

		fmt.Print("Insert your civil status: ")
		fmt.Scan(&civilStatus)

		if gender == 1 {
			men += 1
			if civilStatus == 3 {
				marrMen += 1
				sumAgeMarrMen += age
			    if age > 30 {
					marrOld30 += 1
				}
			}
			if civilStatus == 1 {
				singleMen += 1
			}
		} else {
			women += 1
			if civilStatus == 1 && age < 25 {
				womSingYoung25 += 1
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
	ageIndex++
}

	if marrMen > 0 {
		aveAgeMarrMen = float64(sumAgeMarrMen) / float64(marrMen)
		fmt.Printf("The average age of the married men is: %.2f\n", aveAgeMarrMen)
	}
	
	fmt.Printf("The amount of married men older than 30 is: %d\n", marrOld30)
	fmt.Printf("The amount of single women younger than 25 is: %d\n", womSingYoung25)
	fmt.Printf("The total amount of men is: %d, and the women is: %d\n", men, women)
	fmt.Printf("The amount of single men is: %d\n", singleMen)
	fmt.Printf("The oldest women is: %s\n", olderWom)
}

