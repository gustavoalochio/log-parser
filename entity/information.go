package entity

type UpdatePlayerInformation struct {
	PlayerId string
	Name     string
}

type PlayerConnectedInformation struct {
	PlayerId string
}

type PlayerKilledInformation struct {
	KillerName string
	DeadName   string
	Reason     string
}

type ResultGameInformation struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}
