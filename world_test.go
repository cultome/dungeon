package dungeon

import (
	"fmt"
	"testing"
)

func TestBuildWorld(t *testing.T) {
	world := buildWorld()

	if world == nil {
		t.Fatal("Unable to build world!")
	}
}

func TestWalkTheWorld(t *testing.T) {
	world := buildWorld()
	player := world.Zones[0].Players[0]
	zone0 := world.Zones[0]
	zone1 := world.Zones[1]
	zone2 := world.Zones[2]

	fmt.Printf("[1] Player is in Zone: %+v\n", player.Zone)
	// move to another connected and available zone
	err := world.MovePlayer(player, zone1)
	if err != nil {
		t.Fatalf("0->1 should be a valid movement: %s", err)
	}

	fmt.Printf("[2] Player is in Zone: %+v\n", player.Zone)
	// move back but path not available
	err = world.MovePlayer(player, zone0)
	if err == nil {
		t.Fatal("1->0 should not be a valid movement")
	}

	fmt.Printf("[*] Player is in Zone: %+v\n", player.Zone)
	// move forward
	err = world.MovePlayer(player, zone2)
	if err != nil {
		t.Fatalf("1->2 should be a valid movement: %s", err)
	}

	fmt.Printf("[3] Player is in Zone: %+v\n", player.Zone)
	// close the cycle
	err = world.MovePlayer(player, zone0)
	if err != nil {
		t.Fatalf("2->0 should be a valid movement: %s", err)
	}

	fmt.Printf("[4] Player is in Zone: %+v\n", player.Zone)
}

func buildWorld() *World {
	player := BuildPlayer()

	cost := BuildTripCost()

	zone0 := BuildZone()
	zone1 := BuildZone()
	zone2 := BuildZone()

	path01 := BuildPath("0->1", zone1, cost)
	path02 := BuildPath("0->2", zone1, cost)
	path12 := BuildPath("1->2", zone2, cost)
	path20 := BuildPath("2->0", zone0, cost)
	path21 := BuildPath("2->1", zone1, cost)

	zone0.AddPlayer(player)

	zone0.AddPath(path01)
	zone0.AddPath(path02)
	zone1.AddPath(path12)
	zone2.AddPath(path20)
	zone2.AddPath(path21)

	zones := []*Zone{zone0, zone1, zone2}
	world := BuildWorld(zones)

	return world
}
