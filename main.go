package main

import (
	"fmt"
)

func gradeToPoints(grade string) (float64, error) {
	grades := map[string]float64{
		"A+": 4.0, "A": 4.0, "A-": 3.7,
		"B+": 3.3, "B": 3.0, "B-": 2.7,
		"C+": 2.3, "C": 2.0, "C-": 1.7,
		"D+": 1.3, "D": 1.0, "D-": 0.7,
		"F":  0.0,
	}
	points, ok := grades[grade]
	if !ok {
		return 0, fmt.Errorf("invalid grade: %s", grade)
	}
	return points, nil
}

func main() {
	var numCourses int
	fmt.Print("How many courses? ")
	fmt.Scan(&numCourses)

	totalPoints := 0.0
	totalCredits := 0.0

	for i := 1; i <= numCourses; i++ {
		var name, grade string
		var credits float64

		fmt.Printf("\nCourse %d name: ", i)
		fmt.Scan(&name)
		fmt.Printf("Credits: ")
		fmt.Scan(&credits)
		fmt.Printf("Grade (e.g. A, B+, C-): ")
		fmt.Scan(&grade)

		points, err := gradeToPoints(grade)
		if err != nil {
			fmt.Println("Invalid grade, skipping.")
			continue
		}

		totalPoints += points * credits
		totalCredits += credits
	}

	if totalCredits == 0 {
		fmt.Println("No valid courses entered.")
		return
	}

	gpa := totalPoints / totalCredits
	fmt.Printf("\nYour GPA: %.2f\n", gpa)
}
