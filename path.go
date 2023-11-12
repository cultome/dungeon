package dungeon

type Path struct {
	ID      string
	Destiny *Zone
	Cost    *TripCost
}

func BuildPath(id string, destiny *Zone, cost *TripCost) *Path {
	return &Path{ID: id, Destiny: destiny, Cost: cost}
}
