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

type Exercise struct {
	Type        string
	duration    string
	description string
}

type Physical struct {
	height   string
	weight   string
	age      string
	activity string
}

var allFood []Food

func getFoodHelper() Food {
	var name string
	var cal int
	var description string

	fmt.Println("Enter food name")
	fmt.Scanln(&name)
	fmt.Println("Enter Calories")
	fmt.Scanf("%d", &cal)
	fmt.Println("Enter Food description")
	fmt.Scanln(&description)

	fo := Food{name, cal, description}
	return fo
}

func printFood() {
	for i := 0; i < len(allFood); i++ {
		fmt.Println(allFood[i])
	}
}

func addFood() {
	fmt.Println("For Adding Food Press 1")
	fmt.Println("For Deleting Food Press 2")
	var val int
	_, err := fmt.Scanf("%d", &val)
	if err != nil {
		fmt.Println("Error while entering food option value")
	}
	switch {
	case val == 1:
		fmt.Println("Adding Food")
		foo := getFoodHelper()
		allFood = append(allFood, foo)
	case val == 2:
		fmt.Println("Deleting Food")
		var fname string
		flag := false
		fmt.Println("Enter the food to delete")
		fmt.Scanln(&fname)
		for i := 0; i < len(allFood); i++ {
			if allFood[i].name == fname {
				fmt.Println("Removing Food item")
				allFood[i] = allFood[len(allFood)-1]
				allFood = allFood[:len(allFood)-1]
				flag = true
				break
			}
		}
		if !flag {
			fmt.Println("Food item not present in system")
		}
	}
}
func logExercise() Exercise {
	var exerciseType string
	var exerciseDuration string
	var caloriesBurned string

	fmt.Println("Enter the type of Exercise (ex: Cardio or weight training)")
	fmt.Scanln(&exerciseType)
	fmt.Println("Enter the exercise duration (ex: number of sets or time training")
	fmt.Scanln(&exerciseDuration)
	fmt.Println("Enter the number of calories burned (Approximate or exact)")
	fmt.Scanln(&caloriesBurned)

	fo := Exercise{exerciseType, exerciseDuration, caloriesBurned}
	return fo

}

func addPhysicalSpecs() Physical {
	var height string
	var weight string
	var age string
	var activity string

	fmt.Println("Enter your height (Feet'Inches)")
	fmt.Scanln(&height)
	fmt.Println("Enter your weight (lbs pounds)")
	fmt.Scanln(&weight)
	fmt.Println("Enter your age?")
	fmt.Scanln(&age)
	fmt.Println("How active are you?")
	fmt.Scanln(&activity)

	fo := Physical{height, weight, age, activity}
	return fo
}

func main() {

	for {
		fmt.Println("Welcome to Calorie Tracker")

		fmt.Println("For Adding/Removing Food Press 1")
		fmt.Println("For Calculate Total Calories Press 2")
		fmt.Println("For Adding Personal Physical Information Press 3")
		fmt.Println("For listing all food items Press 4 ")
		fmt.Println("For logging exercise Press 5 ")
		fmt.Println("To Exit Press 6")

		var caseVal int
		_, err := fmt.Scanf("%d", &caseVal)
		if err != nil {
			fmt.Println("Error while entering option value")
		}
		if caseVal < 1 || caseVal > 5 {
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
			addPhysicalSpecs()
		case caseVal == 4:
			fmt.Println("listing all food items")
			printFood()
		case caseVal == 5:
			fmt.Println("Calling functionality for Logging Exercise")
			logExercise()
		case caseVal == 6:
			fmt.Println("Exiting the system")
			os.Exit(0)
		}
	}
}
