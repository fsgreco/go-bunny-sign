package main

import (
	"fmt"
	"os"

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

	bunnysign.Display(phrases)
}
