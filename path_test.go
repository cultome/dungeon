package dungeon

import "testing"

func TestBuildPath(t *testing.T) {
	testId := "12345"
	path := BuildPath(testId)

	if path.ID != testId {
		t.Fatal("Path ID was not setteed correctly")
	}
}
