package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fsgreco/go-bunny-sign/bunnysign"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide at least a phrase to be shown by the bunny.")
		fmt.Println("Example: go run main.go 'Hello, world!'")
		os.Exit(1)
	}

	phrases := os.Args[1:]

	fmt.Println()
	fmt.Println("DEBUG - this line should remain")
	fmt.Println()

	for _, phrase := range phrases {
		msg := []string{}
		words := strings.Split(phrase, " ")

		// Track the previous frame's line count
		previousFrameLines := 0

		for i, word := range words {
			msg = append(msg, word)

			frame := bunnysign.Generate(strings.Join(msg, " "))

			if i > 0 {
				clearLines(previousFrameLines)
			}
			fmt.Printf("%s\n", frame)

			// Save the current frame's line count for the next iteration
			previousFrameLines = strings.Count(frame, "\n") + 1

			time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
			if i == len(words)-1 {
				time.Sleep(800 * time.Millisecond)
			}
		}
	}
}

// Clears the last n lines from the terminal.
// n is the number of lines to clear.
// Note: This function assumes the terminal supports ANSI escape codes.
func clearLines(n int) {
	if n > 0 {
		// \033[{n}A moves the cursor up n lines
		// \033[J clears from the cursor to the end of the screen
		fmt.Printf("\033[%dA\033[J", n)
	}
}
