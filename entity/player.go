package entity

type Player struct {
	ID    string
	Name  string
	Kills int
}

func NewPlayer(id string) *Player {
	return &Player{
		ID:    id,
		Kills: 0,
	}
}

func (p *Player) UpdatePlayer(name string) {
	p.Name = name
}

func (p *Player) AddKill(value int) {
	p.Kills += value
}
