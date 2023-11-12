package dungeon

import "testing"

func TestBuildPlayer(t *testing.T) {
	player := BuildPlayer()

	if player == nil {
		t.Fatal("Player not builded!")
	}
}
