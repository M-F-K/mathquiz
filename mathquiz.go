package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Stats struct {
	correct   int
	incorrect int
	total     int
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	stats := Stats{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Math Quiz!")
	fmt.Println("Answer multiplication questions or type 'q' to quit.")
	fmt.Println()

	for {
		// Generate random numbers between 1 and 10
		x := rand.Intn(10) + 1
		y := rand.Intn(10) + 1
		correctAnswer := x * y

		fmt.Printf("What is %d * %d? ", x, y)

		// Read user input
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		// Check if user wants to quit
		if strings.ToLower(input) == "q" {
			break
		}

		// Try to parse the input as an integer
		answer, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number or 'q' to quit.")
			fmt.Println()
			continue
		}

		// Check the answer
		stats.total++
		if answer == correctAnswer {
			stats.correct++
			fmt.Println("Correct! 🎉")
		} else {
			stats.incorrect++
			fmt.Printf("Incorrect. The correct answer is %d.\n", correctAnswer)
		}
		fmt.Println()
	}

	// Print statistics
	fmt.Println("\n=== Quiz Statistics ===")
	fmt.Printf("Total questions asked: %d\n", stats.total)
	fmt.Printf("Correct answers: %d\n", stats.correct)
	fmt.Printf("Incorrect answers: %d\n", stats.incorrect)
	
	if stats.total > 0 {
		percentage := float64(stats.correct) / float64(stats.total) * 100
		fmt.Printf("Success rate: %.1f%%\n", percentage)
	}
	
	fmt.Println("\nThanks for playing!")
}
