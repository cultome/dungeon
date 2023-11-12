package dungeon

type Player struct {
	Zone *Zone
}

func BuildPlayer() *Player {
	return &Player{Zone: nil}
}

func (p *Player) RemoveFromZone() error {
	return p.Zone.RemovePlayer(p)
}
