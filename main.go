package main

import "fmt"

func main() {

	fmt.Println("Welcome to Calorie Tracker")

	fmt.Println("For Adding/Removing Food Press 1")
	fmt.Println("For Calculate Total Calories Press 2")
	fmt.Println("For Adding Personal Physical Information Press 3")

	var caseVal int
	_, err := fmt.Scanf("%d", &caseVal)
	if err != nil {
		fmt.Println("case value entered is wrong")
	}
	switch {
	case caseVal == 1:
		fmt.Println("Calling functionality for Adding/Removing Food")
	case caseVal == 2:
		fmt.Println("Calling functionality for Calculating Total Calories")
	case caseVal == 3:
		fmt.Println("Calling functionality for Adding Personal Physical Information")
	}
}
