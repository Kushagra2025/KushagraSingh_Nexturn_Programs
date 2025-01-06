package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Define the structure for each question
type Question struct {
	Question string
	Options  [4]string
	Answer   int
}

func main() {
	// Define a slice of questions for the quiz
	questionBank := []Question{
		{"What is 2 + 2?", [4]string{"1", "2", "3", "4"}, 4},
		{"What is the capital of France?", [4]string{"Berlin", "Madrid", "Paris", "Rome"}, 3},
		{"Which programming language is this quiz written in?", [4]string{"Go", "Java", "C", "Python"}, 1},
		{"What is 10 / 2?", [4]string{"2", "3", "5", "7"}, 3},
	}

	// Variable to track score
	var score int

	// Start the quiz
	fmt.Println("Welcome to the Online Examination System!")
	fmt.Println("Type 'exit' at any time to quit the quiz.")
	fmt.Println()

	// Loop through each question
	for i, q := range questionBank {
		// Display question
		fmt.Printf("Question %d: %s\n", i+1, q.Question)
		for j, option := range q.Options {
			fmt.Printf("%d: %s\n", j+1, option)
		}

		// Timer for each question (10 seconds per question)
		timer := time.NewTimer(10 * time.Second)

		// Channel to capture user input
		answerChan := make(chan int)

		go func() {
			// Get user input
			var userAnswer string
			fmt.Print("Your answer (1-4): ")

			// Reading user input
			_, err := fmt.Scan(&userAnswer)
			if err != nil {
				// Handle invalid input
				fmt.Println("Invalid input. Please enter a valid option.")
				answerChan <- -1
				return
			}

			// If the user wants to exit the quiz
			if strings.ToLower(userAnswer) == "exit" {
				answerChan <- -2
				return
			}

			// Try converting the answer to an integer
			userAnswerInt, err := strconv.Atoi(userAnswer)
			if err != nil || userAnswerInt < 1 || userAnswerInt > 4 {
				// Handle invalid integer input
				fmt.Println("Invalid input. Please select an option between 1 and 4.")
				answerChan <- -1
				return
			}
			// Send the valid answer to the channel
			answerChan <- userAnswerInt
		}()

		select {
		case answer := <-answerChan:
			if answer == -2 {
				// User requested to exit the quiz early
				fmt.Println("Exiting the quiz...")
				return
			}
			if answer == -1 {
				// Invalid input, prompt the user again (use continue to skip the invalid question)
				fmt.Println("Invalid input, skipping to next question.")
				continue
			}

			// Check if the answer is correct
			if answer == q.Answer {
				score++
				fmt.Println("Correct!")
			} else {
				fmt.Println("Incorrect.")
			}

		case <-timer.C:
			// If time runs out, skip the question and proceed
			fmt.Println("\nTime's up! Moving to next question...\n")
		}
	}

	// Show the final score and performance classification
	fmt.Println("Quiz completed.")
	fmt.Printf("Your score: %d/%d\n", score, len(questionBank))
	if score == len(questionBank) {
		fmt.Println("Excellent!")
	} else if score > len(questionBank)/2 {
		fmt.Println("Good!")
	} else {
		fmt.Println("Needs Improvement!")
	}
}
