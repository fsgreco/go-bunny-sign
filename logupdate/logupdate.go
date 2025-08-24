package logupdate

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func CreatePlayer(FrameGenerator func(content string) string) func(content string) int {
	previousFrameLines := 0

	return func(content string) int {
		if previousFrameLines > 0 {
			ClearLines(previousFrameLines)
		}
		// print the frame
		frame := FrameGenerator(content)
		fmt.Printf("%s\n", frame)

		previousFrameLines = strings.Count(frame, "\n") + 1
		// wait random to print the new frame
		time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)

		return previousFrameLines
	}
}

// Clears the last n lines from the terminal.
// n is the number of lines to clear.
// Note: This function assumes the terminal supports ANSI escape codes.
func ClearLines(n int) {
	if n > 0 {
		// \033[{n}A moves the cursor up n lines
		// \033[J clears from the cursor to the end of the screen
		fmt.Printf("\033[%dA\033[J", n)
	}
}
