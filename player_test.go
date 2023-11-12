package dungeon

import "testing"

func TestBuildPlayer(t *testing.T) {
	player := BuildPlayer()

	if player == nil {
		t.Fatal("Player not builded!")
	}
}

func TestRemoveFromZone(t *testing.T) {
	zone := BuildZone()
	player := BuildPlayer()

	zone.AddPlayer(player)

	player.Zone = zone

	err := player.RemoveFromZone()
	if err != nil {
		t.Fatalf("Unable to remove user from zone: %s", err.Error())
	}

	if player.Zone != nil {
		t.Fatalf("Unable to remove user from zone: Player still have a zone")
	}
}
