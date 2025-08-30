package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsgreco/go-bunny-sign/bunnysign"
)

type CLIConfig struct {
	Clear bool
	Debug bool
}

func parseFlags() CLIConfig {
	clear := flag.Bool("c", false, "Clear the bunny after showing the message")
	debug := flag.Bool("debug", false, "Debug CLI features")
	flag.Parse()
	return CLIConfig{
		Clear: *clear,
		Debug: *debug,
	}
}

func main() {

	config := parseFlags()

	phrases := flag.Args() // os.Args[1:]
	if len(phrases) == 0 {
		fmt.Println("Please provide at least a phrase to be shown by the bunny.")
		fmt.Println("Example: bunnysign 'Hello, world!'")
		os.Exit(1)
	}

	if config.Debug {
		fmt.Println()
		fmt.Println("DEBUG - this line should remain")
	}
	fmt.Println()

	bunnysign.Display(phrases, !config.Clear)
}
