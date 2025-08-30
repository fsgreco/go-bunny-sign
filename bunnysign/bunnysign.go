// Provides an animated messages for terminal applications displaying a bunny holding a sign.
package bunnysign

import (
	"strings"
	"time"

	"github.com/fsgreco/go-bunny-sign/logupdate"
)

// Take a slice of phrases and displays them word by word with animation.
// The persist parameter controls whether the final message remains on screen.
func Display(phrases []string, persist bool) {
	render := logupdate.CreatePlayer(Generate)

	for indxPhrase, phrase := range phrases {
		isLastPhrase := indxPhrase == len(phrases)-1
		msg := ""
		words := strings.Split(phrase, " ")

		for i, word := range words {
			isLastWord := i == len(words)-1

			if i > 0 {
				msg += " "
			}
			msg += word

			lines := render(msg)

			if isLastWord {
				time.Sleep(800 * time.Millisecond)
			}
			isLastFrame := isLastPhrase && isLastWord
			if isLastFrame && !persist {
				// when using the CLI -c flag will set Persist to false
				logupdate.ClearLines(lines)
			}
		}
	}
}

// Creates a bunny sign ASCII art with the given message.
// Returns a formatted string with the message already wrapped by the sign.
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

// Wraps text to fit within the bunny sign with proper padding.
// It handles line breaks and ensures text fits within the specified maxLength.
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
