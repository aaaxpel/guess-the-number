package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var number int
var difficulty string
var guesses int
var guessesLeft int

func chooseDifficulty() {
	fmt.Print("Enter your choice: ")
	fmt.Scan(&difficulty)
	switch difficulty {
	case "1":
		fmt.Printf("Great! You have selected the Easy difficulty level.\nLet's start the game!\n\n")
		guesses = 10
		difficulty = "Easy"
	case "2":
		fmt.Printf("Great! You have selected the Medium difficulty level.\nLet's start the game!\n\n")
		guesses = 5
		difficulty = "Medium"
	case "3":
		fmt.Printf("Great! You have selected the Hard difficulty level.\nLet's start the game!\n\n")
		guesses = 3
		difficulty = "Hard"
	default:
		fmt.Println("Choose between 1 (easy), 2 (medium) and 3 (hard)")
		chooseDifficulty()
	}
	guessesLeft = guesses
}

func play() bool {
	var guess string
	if guessesLeft > 0 {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		guessedNumber, err := strconv.Atoi(guess)
		if err != nil || guessedNumber > 100 || guessedNumber < 1 {
			fmt.Println("Please choose a number between 1 - 100")
			play()
		}

		guessesLeft--

		switch {
		case guessedNumber > number:
			fmt.Printf("Incorrect! The number is less than %v. Guesses left: %v\n\n", guessedNumber, guessesLeft)
			play()
		case guessedNumber < number:
			fmt.Printf("Incorrect! The number is greater than %v. Guesses left: %v\n\n", guessedNumber, guessesLeft)
			play()
		case guessedNumber == number:
			fmt.Printf("Congratulations! You guessed the number in %v attempts.\n\n", guesses-guessesLeft)
			return true
		}
	} else {
		fmt.Printf("Unfortunately you didn't guess %v.\n\n", number)
	}
	return false
}

func main() {
	var playAgain string

	if playAgain == "" {
		fmt.Printf("\nWelcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\n\n")
	}
	number = rand.Intn(100-1) + 1
	number = 1

	fmt.Printf("Select difficulty level:\n1. Easy (10 guesses)\n2. Medium (5 guesses)\n3. Hard (3 guesses)\n\n")

	chooseDifficulty()

	start := time.Now()

	if result := play(); result == true {
		fmt.Printf("It took you %vs to guess the number on %v difficulty!\n\n", time.Now().Sub(start).Seconds(), difficulty)
	}

	fmt.Print("Would you like to play again? (y/n): ")
	fmt.Scan(&playAgain)
	if playAgain == "y" {
		main()
	}
	return
}
