package main

import "github.com/fsgreco/go-bunny-sign/bunnysign"

func main() {
	// Simple usage
	phrases := []string{"Hello, World!", "This is a bunny sign!"}
	bunnysign.Display(phrases, true) // true = persist on screen
}
