package main

import "fmt"

func main(){
	var name, olderWom string
	var gender, civilStatus, age, womAgeIndex int
	var marrOld30, marrMen, men, women, closingIndex int
	var singleMen, womSingYoung25, sumAgeMarrMen, womMaxAge int
	var aveAgeMarrMen float64

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
	
		fmt.Print("Insert your gender: ")
		fmt.Scan(&gender)

		fmt.Print("Insert your civil status: ")
		fmt.Scan(&civilStatus)

		if gender == 1 {
			men++
			if civilStatus == 3 {
				marrMen++
				sumAgeMarrMen += age
			    if age > 30 {
					marrOld30++
				}
			}
			if civilStatus == 1 {
				singleMen++
			}
		} else {
			women++
			if civilStatus == 1 && age < 25 {
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
		aveAgeMarrMen = float64(sumAgeMarrMen) / float64(marrMen)
		fmt.Printf("The average age of the married men is: %.2f\n", aveAgeMarrMen)
	}

	if closingIndex > 0 {
		fmt.Printf("The amount of married men older than 30 is: %d\n", marrOld30)
		fmt.Printf("The amount of single women younger than 25 is: %d\n", womSingYoung25)
		fmt.Printf("The total amount of men is: %d, and the women is: %d\n", men, women)
		fmt.Printf("The amount of single men is: %d\n", singleMen)
		fmt.Printf("The oldest women is: %s\n", olderWom)	
	}	
}

