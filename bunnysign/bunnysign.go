package bunnysign

import (
	"strings"
	"time"

	"github.com/fsgreco/go-bunny-sign/logupdate"
)

func Display(phrases []string) {
	render := logupdate.CreatePlayer(Generate)

	for _, phrase := range phrases {
		msg := ""
		words := strings.Split(phrase, " ")

		for i, word := range words {

			if i > 0 {
				msg += " "
			}
			msg += word

			render(msg)

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
