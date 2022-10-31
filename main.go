package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type variables struct {
	date string
	time string
}

func main() {
	http.HandleFunc("/", addPersonal)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func addPersonal(w http.ResponseWriter, r *http.Request) {

	now := time.Now()
	HomePageVars := variables{
		date: now.Format("02-01-2006"),
		time: now.Format("15:04:05"),
	}

	t, error := template.ParseFiles("./views/index.html")
	if error != nil {
		log.Print("template parsing error: ", error)
	}
	error = t.Execute(w, HomePageVars)
	if error != nil {
		log.Print("template executing error: ", error)
	}
}

var totalCalories int

func addFood(calories int) int {
	totalCalories = calories + totalCalories
	return totalCalories
}

func removeFood(calories int) int {
	totalCalories = calories - totalCalories
	return totalCalories
}

func calcTotal(calories int) {

}

func addPhysical(height string, weight string) {
	var strArray [2]string
	strArray[0] = height
	strArray[1] = weight
}

func logExercise(exercise string) string {
	return exercise
}
