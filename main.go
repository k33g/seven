package main

import (
	"flag"
	"fmt"
	"os"
	"seven/cmd"
)

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("😡 invalid command")
		os.Exit(1)
	}

	command := flag.Args()[0]

	errCmd := cmd.Parse(command, flag.Args()[1:])

	if errCmd != nil {
		fmt.Println("😡", errCmd)
		os.Exit(1)
	}
}
