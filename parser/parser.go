package parser

import (
	"errors"
	"fmt"
	"log-parser/entity"
	"regexp"
)

//go:generate mockgen -destination=mocks/parser.go -package=mocks -source=parser.go
type Parser interface {
	ParserInitGame(input string) bool
	ParserInterval(input string) bool
	ParserClientConnect(input string) bool
	ParserClientUserInfoChanged(input string) bool
	ParserKilled(input string) bool
	GetInformationPlayerUpdate(input string) (entity.UpdatePlayerInformation, error)
	GetInformationPlayerConnected(input string) (entity.PlayerConnectedInformation, error)
	GetInformationKilledPlayer(input string) (entity.PlayerKilledInformation, error)
}

type parser struct{}

func NewParser() Parser {
	return &parser{}
}

func (p *parser) ParserInitGame(input string) bool {
	regex := regexp.MustCompile(`.*?(\d+:\d+) InitGame:`)
	match := regex.FindStringSubmatch(input)

	return len(match) >= 1
}

func (p *parser) ParserInterval(input string) bool {
	regex := regexp.MustCompile(`\s+([-\s]+)$`)
	match := regex.FindStringSubmatch(input)

	return len(match) >= 1
}

func (p *parser) ParserClientConnect(input string) bool {
	regex := regexp.MustCompile(`(\d+:\d+)\s+ClientConnect:\s+(\d+)`)
	match := regex.FindStringSubmatch(input)

	return len(match) >= 1
}

func (p *parser) ParserClientUserInfoChanged(input string) bool {
	regex := regexp.MustCompile(`(\d+:\d+)\s+ClientUserinfoChanged:\s+(\d+)\s+n\\(.*?)\\t\\`)
	match := regex.FindStringSubmatch(input)

	return len(match) >= 1
}

func (p *parser) ParserKilled(input string) bool {
	regex := regexp.MustCompile(`\d+:\d+\s+Kill:\s+\d+\s+\d+\s+\d+:\s+(.*?)\skilled\s(.*?)\sby\s(.*?)$`)

	match := regex.FindStringSubmatch(input)

	return len(match) >= 1
}

// ----------------------------------------------------------------

func (p *parser) GetInformationPlayerUpdate(input string) (entity.UpdatePlayerInformation, error) {

	playerInformation := entity.UpdatePlayerInformation{}
	regex := regexp.MustCompile(`(\d+:\d+)\s+ClientUserinfoChanged:\s+(\d+)\s+n\\(.*?)\\t\\`)
	match := regex.FindStringSubmatch(input)

	if len(match) > 3 {
		playerInformation.PlayerId = match[2]
		playerInformation.Name = match[3]

		return playerInformation, nil
	}
	fmt.Println("debug2")

	return playerInformation, errors.New("invalid input")
}

func (p *parser) GetInformationPlayerConnected(input string) (entity.PlayerConnectedInformation, error) {

	playerConnected := entity.PlayerConnectedInformation{}
	regex := regexp.MustCompile(`(\d+:\d+)\s+ClientConnect:\s+(\d+)`)
	match := regex.FindStringSubmatch(input)

	if len(match) > 2 {
		playerConnected.PlayerId = match[2]

		return playerConnected, nil
	}

	return playerConnected, errors.New("invalid input")
}

func (p *parser) GetInformationKilledPlayer(input string) (entity.PlayerKilledInformation, error) {

	playerKilledInformation := entity.PlayerKilledInformation{}

	regex := regexp.MustCompile(`\d+:\d+\s+Kill:\s+\d+\s+\d+\s+\d+:\s+(.*?)\skilled\s(.*?)\sby\s(.*?)$`)
	match := regex.FindStringSubmatch(input)

	if len(match) > 3 {
		playerKilledInformation.KillerName = match[1]
		playerKilledInformation.DeadName = match[2]
		playerKilledInformation.Reason = match[3]

		return playerKilledInformation, nil
	}

	return playerKilledInformation, errors.New("invalid input")
}
