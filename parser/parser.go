package parser

import (
	"errors"
	"fmt"
	"log-parser/entity"
	"regexp"
)

func ParserInitGame(input string) bool {
	regex := regexp.MustCompile(`.*?(\d+:\d+) InitGame:`)
	match := regex.FindStringSubmatch(input)

	return len(match) >= 1
}

func ParserInterval(input string) bool {
	regex := regexp.MustCompile(`\s+([-\s]+)$`)
	match := regex.FindStringSubmatch(input)

	return len(match) >= 1
}

func ParserClientConnect(input string) bool {
	regex := regexp.MustCompile(`.*?(\d+:\d+) ClientConnect:`)
	match := regex.FindStringSubmatch(input)

	return len(match) >= 1
}

func ParserClientUserInfoChanged(input string) bool {
	regex := regexp.MustCompile(`(\d+:\d+)\s+ClientUserinfoChanged:\s+(\d+)\s+n\\(.*?)\\t\\`)
	match := regex.FindStringSubmatch(input)

	return len(match) >= 1
}

func ParserKilled(input string) bool {
	regex := regexp.MustCompile(`\d+:\d+\s+Kill:\s+\d+\s+\d+\s+\d+:\s+(.*?)\skilled\s(.*?)\sby\s(.*?)$`)

	match := regex.FindStringSubmatch(input)

	return len(match) >= 1
}

// ----------------------------------------------------------------

func GetInformationPlayerUpdate(input string) (entity.UpdatePlayerInformation, error) {

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

func GetInformationPlayerConnected(input string) (entity.PlayerConnectedInformation, error) {

	playerConnected := entity.PlayerConnectedInformation{}
	regex := regexp.MustCompile(`(\d+:\d+)\s+ClientConnect:\s+(\d+)`)
	match := regex.FindStringSubmatch(input)

	if len(match) > 2 {
		playerConnected.PlayerId = match[2]

		return playerConnected, nil
	}

	return playerConnected, errors.New("invalid input")
}

func GetInformationKilledPlayer(input string) (entity.PlayerKilledInformation, error) {

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
