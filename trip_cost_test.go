package dungeon

import "testing"

func TestBuildTripCost(t *testing.T) {
	cost := BuildTripCost()

	if cost == nil {
		t.Fatal("Unable to build TripCost")
	}
}
