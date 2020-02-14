package player

type Player struct {
	Name string
}

func NewPlayer(name string) (p *Player, err error) {
	p = &Player{}
	p.Name = name

	return p, nil
}
