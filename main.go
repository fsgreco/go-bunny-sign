package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsgreco/go-bunny-sign/bunnysign"
)

type CLIConfig struct {
	Clear bool
}

func parseFlags() CLIConfig {
	clear := flag.Bool("c", false, "Clear the bunny after showing the message")
	flag.Parse()
	return CLIConfig{
		Clear: *clear,
	}
}

func main() {

	config := parseFlags()

	phrases := flag.Args() // os.Args[1:]
	if len(phrases) == 0 {
		fmt.Println("Please provide at least a phrase to be shown by the bunny.")
		fmt.Println("Example: go run main.go 'Hello, world!'")
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("DEBUG - this line should remain")
	fmt.Println()

	options := bunnysign.Options{Persist: !config.Clear}

	bunnysign.Display(phrases, options)

	fmt.Println()
}
