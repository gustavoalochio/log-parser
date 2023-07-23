package entity

type Game struct {
	Running       bool
	Players       []*Player
	PlayersByName map[string]*Player
	PlayersById   map[string]*Player
	DeadReasons   map[string]int
}

func NewGame() *Game {
	return &Game{
		Running:       false,
		Players:       []*Player{},
		PlayersByName: make(map[string]*Player),
		PlayersById:   make(map[string]*Player),
		DeadReasons:   make(map[string]int),
	}
}

func (g *Game) Enable() {
	g.Running = true
}

func (g *Game) Disable() {
	g.Running = false
}

func (g *Game) AddPlayer(p *Player) {
	if _, found := g.PlayersById[p.ID]; !found {
		g.Players = append(g.Players, p)
		g.PlayersById[p.ID] = p
		g.PlayersByName[p.Name] = p
	}
}

func (g *Game) AddDeadReason(reason string) {
	if _, found := g.DeadReasons[reason]; found {
		g.DeadReasons[reason] += 1
	} else {
		g.DeadReasons[reason] = 1
	}
}

func InitGame() *Game {
	game := NewGame()
	game.Enable()

	playerWorld := NewPlayer("0")
	playerWorld.UpdatePlayer("<world>")
	game.AddPlayer(playerWorld)

	return game
}
