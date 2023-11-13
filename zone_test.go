package dungeon

import (
	"testing"
)

func TestBuildZone(t *testing.T) {
	zone := BuildZone()

	if zone == nil {
		t.Fatal("Zone not builded!")
	}
}

func TestAddPlayer(t *testing.T) {
	player := BuildPlayer()
	zone := BuildZone()

	if len(zone.Players) != 0 {
		t.Fatal("Pre-conditions not met: Players in zone")
	}

	err := zone.AddPlayer(player)
	if err != nil {
		t.Fatalf("Unable to add player to zone: %s", err.Error())
	}

	if len(zone.Players) != 1 {
		t.Fatalf("Post-conditions not met: New players not in zone")
	}

	// try to add existing player
	err = zone.AddPlayer(player)
	if err == nil {
		t.Fatalf("Fail to verify existing player in zone")
	}
}

func TestRemovePlayer(t *testing.T) {
	player := BuildPlayer()
	zone := BuildZone()

	err := zone.AddPlayer(player)

	if len(zone.Players) != 1 {
		t.Fatalf("Pre-conditions not met: Player not in zone")
	}

	err = zone.RemovePlayer(player)
	if err != nil {
		t.Fatalf("Unable to remove player from zone: %s", err.Error())
	}

	if len(zone.Players) != 0 {
		t.Fatalf("Post-conditions not met: Player still in zone")
	}

	// try to remove unexistent player
	err = zone.RemovePlayer(player)
	if err == nil {
		t.Fatalf("Failed to verify non-existent player in zone")
	}
}

func TestAddPath(t *testing.T) {
	path := BuildPath("12345", nil, nil)
	zone := BuildZone()

	if len(zone.Paths) != 0 {
		t.Fatalf("Pre-conditions not met: Path in zone")
	}

	err := zone.AddPath(path)
	if err != nil {
		t.Fatalf("Unable to add path to zone: %s", err.Error())
	}

	if len(zone.Paths) != 1 {
		t.Fatalf("Post-conditions not met: Path not in zone")
	}

	// try to add existing path
	err = zone.AddPath(path)
	if err == nil {
		t.Fatalf("Unable to detect existing path in zone")
	}
}

func TestAddThing(t *testing.T) {
	zone := BuildZone()
	thing := BuildThingTv()

	if len(zone.Things) != 0 {
		t.Fatalf("Pre-conditions not met: Things in zone")
	}

	zone.AddThing(thing)

	if len(zone.Things) != 1 {
		t.Fatalf("Post-conditions not met: Things in zone")
	}

	// try to add thing to "another" zone without desassociate first
	err := zone.AddThing(thing)

	if err == nil {
		t.Fatalf("Unable to detect thing in another zone")
	}
}

func TestRemoveThing(t *testing.T) {
	zone := BuildZone()
	thing := BuildThingTv()
	zone.AddThing(thing)

	if len(zone.Things) != 1 {
		t.Fatalf("Pre-conditions not met: Things not in zone")
	}

	zone.RemoveThing(thing)

	if len(zone.Things) != 0 {
		t.Fatalf("Post-conditions not met: Things still in zone")
	}

	// try to remove non-existent thing
	err := zone.RemoveThing(thing)

	if err == nil {
		t.Fatalf("Failed to detect non-existing thing in zone")
	}
}
