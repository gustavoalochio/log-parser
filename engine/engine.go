package engine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log-parser/entity"
	"log-parser/parser"
	"os"
)

const WorldPlayerName = "<world>"

func ParsingLog() {
	scanner := bufio.NewScanner(os.Stdin)
	gameCounter := 0
	game := entity.NewGame()

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case parser.ParserInitGame(line):
			gameCounter += 1
			game = entity.InitGame()

		case parser.ParserInterval(line):
			if game.Running {
				result := endGame(game)
				printResultGame(result, gameCounter)
				printDeadReasons(game, gameCounter)
			}
		case parser.ParserClientConnect(line):
			playerConnectedInformation, err := parser.GetInformationPlayerConnected(line)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			newPlayer := entity.NewPlayer(playerConnectedInformation.PlayerId)
			game.AddPlayer(newPlayer)

		case parser.ParserClientUserInfoChanged(line):

			playerUpdatedInformation, err := parser.GetInformationPlayerUpdate(line)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			game.PlayersById[playerUpdatedInformation.PlayerId].UpdatePlayer(playerUpdatedInformation.Name)
			game.PlayersByName[playerUpdatedInformation.Name] = game.PlayersById[playerUpdatedInformation.PlayerId]

		case parser.ParserKilled(line):

			playerKilledInformation, err := parser.GetInformationKilledPlayer(line)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			kill(game, playerKilledInformation)
			game.AddDeadReason(playerKilledInformation.Reason)

		default:
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("read err:", err)
	}
}

func endGame(game *entity.Game) entity.ResultGameInformation {
	game.Disable()

	resultGameInformation := entity.ResultGameInformation{
		Kills: make(map[string]int),
	}

	for _, player := range game.Players {
		resultGameInformation.TotalKills += player.Kills

		if player.Name != WorldPlayerName {
			resultGameInformation.Players = append(resultGameInformation.Players, player.Name)
			resultGameInformation.Kills[player.Name] = player.Kills
		}
	}

	return resultGameInformation
}

func printResultGame(result entity.ResultGameInformation, gameCounter int) {

	gameInstance := fmt.Sprintf("game_%d", gameCounter)

	jsonResult, err := json.Marshal(map[string]entity.ResultGameInformation{gameInstance: result})
	if err != nil {
		fmt.Println("converter error - JSON:", err)
		return
	}

	fmt.Println(string(jsonResult))
}

func printDeadReasons(game *entity.Game, gameCounter int) {
	gameInstance := fmt.Sprintf("game-%d", gameCounter)

	if len(game.DeadReasons) == 0 {
		return
	}

	jsonResult, err := json.Marshal(map[string]map[string]int{gameInstance: game.DeadReasons})
	if err != nil {
		fmt.Println("converter error - JSON:", err)
		return
	}
	fmt.Println(string(jsonResult))
}

func kill(game *entity.Game, playerKilledInformation entity.PlayerKilledInformation) {
	if playerKilledInformation.KillerName == playerKilledInformation.DeadName {
		return
	} else if playerKilledInformation.KillerName == WorldPlayerName {
		game.PlayersByName[playerKilledInformation.DeadName].AddKill(-1)
	} else {
		game.PlayersByName[playerKilledInformation.KillerName].AddKill(1)
	}
}
