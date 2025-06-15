package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if reader.Scan() {
			words := cleanInput(reader.Text())
			if len(words) == 0 {
				continue
			}
			inputCommand := words[0]
			cmd, ok := getCommands()[inputCommand]

			if !ok {
				fmt.Println("Unknown command")
				continue
			}

			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}
