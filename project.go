package main //The package main tells the Go compiler that the package should compile as an executable program instead of a shared library.

import "fmt" //fmt is a Go package that is used to format basic strings, values, inputs, and outputs.

//main function
func main() {
	fmt.Println("Welcome to the Quiz Game!")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name) //scans the user input

	fmt.Printf("Hello, %v, Welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay! You can now play the game!")
	} else {
		fmt.Println("You can't play the game, invalid age!")
		return
	}
	//variables for score and number of questions, := automatically detects the data type
	score := 0
	num_questions := 2

	fmt.Printf("Which is bigger, an Elephant or an Airplane? ")
	var question1 string
	fmt.Scan(&question1)
	//conditional statements in golang
	if question1 == "Airplane" || question1 == "airplane" || question1 == "Aeroplane" || question1 == "aeroplane" {
		fmt.Println("Correct!")
		score++

	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("What animal meows? ")
	var question2 string
	fmt.Scan(&question2)

	if question2 == "Cat" || question2 == "cat" || question2 == "Cheetah" {
		fmt.Println("Correct!")
		score++

	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You scored %v out of %v questions!", score, num_questions)
	// arithmetic logic to get the percentage score of the game
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf(" You scored: %v%%.", percent)
}
