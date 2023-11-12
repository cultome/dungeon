package dungeon

import (
	"fmt"
)

type World struct {
	Zones []*Zone
}

func BuildWorld(zones []*Zone) *World {
	return &World{Zones: zones}
}

func (w *World) MovePlayer(player *Player, zone *Zone) error {
	cost := calculateTripCost(player.Zone, zone)

	if cost.Valid {
		player.RemoveFromZone()
		zone.AddPlayer(player)
		return nil
	} else {
		return fmt.Errorf("Unable to move [%+v] from [%+v] to [%+v]", player, player.Zone, zone)
	}
}

func calculateTripCost(zoneA, zoneB *Zone) *TripCost {
	for _, path := range zoneA.Paths {
		if path.Destiny == zoneB {
			return path.Cost
		}
	}

	return &TripCost{Valid: false}
}
