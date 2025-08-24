package bunnysign

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Display(phrases []string) {

	for _, phrase := range phrases {
		msg := []string{}
		words := strings.Split(phrase, " ")

		// Track the previous frame's line count
		previousFrameLines := 0

		for i, word := range words {
			msg = append(msg, word)

			frame := Generate(strings.Join(msg, " "))

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

func Generate(message string) string {
	padding := strings.Repeat(" ", 4)

	sign := []string{
		" |￣￣￣￣￣￣￣￣￣￣￣|",
		wrapText(message, padding, 22),
		" |＿＿＿＿＿＿＿＿＿＿＿|",
		"   (\\__/) ||",
		"   (•ㅅ•) ||",
		"   /    づ",
	}

	return strings.Join(sign, "\n")
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

func wrapText(phrase string, padding string, maxLength int) string {
	if maxLength == 0 {
		maxLength = 18
	}
	if padding == "" {
		padding = strings.Repeat(" ", 4)
	}

	words := strings.Split(phrase, " ")
	words[0] = padding + words[0] // Add padding to the first word

	lineCount := 0
	for i, word := range words {
		lineCount += len(word) + 1 // +1 for the space
		if lineCount > maxLength {
			lineCount = len(padding) + len(word)
			words[i] = "\n" + padding + words[i]
		}
	}

	wrappedText := []string{strings.Join(words, " ")}

	return strings.Join(wrappedText, "\n")
}
