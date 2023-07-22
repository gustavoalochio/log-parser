package main

import (
	"bufio"
	"fmt"
	parser "log-parser/parser"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	gameRunning := false

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case parser.ParserInitGame(line):
			gameRunning = true
			//create game
			//create player <world>
			fmt.Println("Game is running!")
		case parser.ParserInterval(line):
			if gameRunning {
				gameRunning = false
				fmt.Println("Game has been finished")
			}
		// case parser.ParserClientConnect(line):
		//create player in game
		// case parser.ParserClientInfoUpdate(line):
		//update player in game
		// case parser.ParserKilled(line):
		//update deaths to player
		//update reasons death hash in game
		default:
			fmt.Println("something")
		}

		// create final message, list players and deaths
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("read err:", err)
	}
}
