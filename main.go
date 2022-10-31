package main

import (
	"fmt"
	"os"
)

type Food struct {
	name        string
	calories    int
	description string
}

func addFood() {
	fmt.Println("For Adding Food Press 1")
	fmt.Println("For Deleting Food Press 2")
	var val int
	_, err := fmt.Scanf("%d", &val)
	if err != nil {
		fmt.Println("Error while entering food option value")
	}
}

func main() {

	fmt.Println("Welcome to Calorie Tracker")

	fmt.Println("For Adding/Removing Food Press 1")
	fmt.Println("For Calculate Total Calories Press 2")
	fmt.Println("For Adding Personal Physical Information Press 3")

	var caseVal int
	_, err := fmt.Scanf("%d", &caseVal)
	if err != nil {
		fmt.Println("Error while entering option value")
	}
	if caseVal < 1 || caseVal > 3 {
		fmt.Println("Invalid case value")
		os.Exit(0)
	}
	switch {
	case caseVal == 1:
		fmt.Println("Calling functionality for Adding/Removing Food")
		addFood()
	case caseVal == 2:
		fmt.Println("Calling functionality for Calculating Total Calories")
	case caseVal == 3:
		fmt.Println("Calling functionality for Adding Personal Physical Information")
	}
}
