package dungeon

import "testing"

func TestBuildPath(t *testing.T) {
	testId := "12345"
	destiny := BuildZone()
	cost := BuildTripCost()
	path := BuildPath(testId, destiny, cost)

	if path == nil {
		t.Fatal("path not builded!")
	}

	if path.ID != testId {
		t.Fatal("Path ID was not setteed correctly")
	}
}
