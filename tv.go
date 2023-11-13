package dungeon

type Tv struct {
	Zone *Zone
}

func BuildThingTv() *Tv {
	return &Tv{Zone: nil}
}

func (t *Tv) GetZone() *Zone {
	return t.Zone
}

func (t *Tv) SetZone(zone *Zone) {
	t.Zone = zone
}
