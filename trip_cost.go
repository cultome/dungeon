package dungeon

type TripCost struct {
	Valid bool
}

func BuildTripCost() *TripCost {
	return &TripCost{Valid: true}
}
