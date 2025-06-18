package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	extraParam := ""
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
			if len(words) > 1 {
				extraParam = words[1]
			}
			err := cmd.callback(extraParam)
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}
