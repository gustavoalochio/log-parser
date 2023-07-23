package engine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log-parser/entity"
	"log-parser/parser"
	"os"
)

//go:generate mockgen -destination=mocks/engine.go -package=mocks -source=engine.go
type Engine interface {
	InitGame() *engine
	Start()
	BuildDeadReasons(game *entity.Game, gameCounter int) string
	BuildResultGame(result entity.ResultGameInformation, gameCounter int) string
	EndGame(game *entity.Game) entity.ResultGameInformation
	Kill(game *entity.Game, playerKilledInformation entity.PlayerKilledInformation)
}

type engine struct {
	Parser parser.Parser
	Game   *entity.Game
}

func NewEngine(parser parser.Parser) Engine {
	return &engine{
		Parser: parser,
		Game:   nil,
	}
}

func (e *engine) InitGame() *engine {
	e.Game = entity.NewGame()
	e.Game.Enable()

	playerWorld := entity.NewPlayer("0")
	playerWorld.UpdatePlayer(e.Game.DefaultPlayerName)
	e.Game.AddPlayer(playerWorld)

	return e
}

func (e *engine) Start() {
	scanner := bufio.NewScanner(os.Stdin)
	gameCounter := 0
	e.Game = entity.NewGame()

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case e.Parser.ParserInitGame(line):
			gameCounter += 1
			e.InitGame()

		case e.Parser.ParserInterval(line):
			if e.Game.Running {
				result := e.EndGame(e.Game)
				fmt.Print(e.BuildResultGame(result, gameCounter))
				fmt.Print(e.BuildDeadReasons(e.Game, gameCounter))
			}
		case e.Parser.ParserClientConnect(line):
			playerConnectedInformation, err := e.Parser.GetInformationPlayerConnected(line)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			newPlayer := entity.NewPlayer(playerConnectedInformation.PlayerId)
			e.Game.AddPlayer(newPlayer)

		case e.Parser.ParserClientUserInfoChanged(line):

			playerUpdatedInformation, err := e.Parser.GetInformationPlayerUpdate(line)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			e.Game.PlayersById[playerUpdatedInformation.PlayerId].UpdatePlayer(playerUpdatedInformation.Name)
			e.Game.PlayersByName[playerUpdatedInformation.Name] = e.Game.PlayersById[playerUpdatedInformation.PlayerId]

		case e.Parser.ParserKilled(line):

			playerKilledInformation, err := e.Parser.GetInformationKilledPlayer(line)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			e.Kill(e.Game, playerKilledInformation)
			e.Game.AddDeadReason(playerKilledInformation.Reason)

		default:
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("read err:", err)
	}
}

func (e *engine) BuildResultGame(result entity.ResultGameInformation, gameCounter int) string {

	gameInstance := fmt.Sprintf("game_%d", gameCounter)

	jsonResult, err := json.Marshal(map[string]entity.ResultGameInformation{gameInstance: result})
	if err != nil {
		fmt.Println("converter error - JSON:", err)
		return ""
	}

	return string(jsonResult) + "\n"
}

func (e *engine) BuildDeadReasons(game *entity.Game, gameCounter int) string {
	gameInstance := fmt.Sprintf("game-%d", gameCounter)

	if len(game.DeadReasons) == 0 {
		return ""
	}

	jsonResult, err := json.Marshal(map[string]map[string]int{gameInstance: game.DeadReasons})
	if err != nil {
		fmt.Println("converter error - JSON:", err)
		return ""
	}
	return string(jsonResult) + "\n"
}

func (e *engine) EndGame(game *entity.Game) entity.ResultGameInformation {
	game.Disable()

	resultGameInformation := entity.ResultGameInformation{
		Kills: make(map[string]int),
	}

	for _, player := range game.Players {
		resultGameInformation.TotalKills += player.Kills

		if player.Name != game.DefaultPlayerName {
			resultGameInformation.Players = append(resultGameInformation.Players, player.Name)
			resultGameInformation.Kills[player.Name] = player.Kills
		}
	}

	return resultGameInformation
}

func (e *engine) Kill(game *entity.Game, playerKilledInformation entity.PlayerKilledInformation) {
	if playerKilledInformation.KillerName == playerKilledInformation.DeadName {
		return
	} else if playerKilledInformation.KillerName == game.DefaultPlayerName {
		game.PlayersByName[playerKilledInformation.DeadName].AddKill(-1)
	} else {
		game.PlayersByName[playerKilledInformation.KillerName].AddKill(1)
	}
}
