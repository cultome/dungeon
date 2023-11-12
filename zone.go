package dungeon

import "fmt"

type Zone struct {
	Paths   []*Path
	Players []*Player
}

func BuildZone() *Zone {
	return &Zone{Paths: []*Path{}, Players: []*Player{}}
}

func (z *Zone) AddPlayer(player *Player) error {
	found := false

	for _, p := range z.Players {
		if p == player {
			found = true
			break
		}
	}

	if found {
		return fmt.Errorf("Player %v already are in zone", player)
	} else {
		z.Players = append(z.Players, player)
		player.Zone = z

		return nil
	}
}

func (z *Zone) AddPath(path *Path) error {
	found := false

	for _, p := range z.Paths {
		if p == path {
			found = true
			break
		}
	}

	if found {
		return fmt.Errorf("Path %v already exist in zone", path)
	} else {
		z.Paths = append(z.Paths, path)

		return nil
	}
}

func (z *Zone) RemovePlayer(player *Player) error {
	foundIdx := -1

	for idx, p := range z.Players {
		if p == player {
			foundIdx = idx
			break
		}
	}

	if foundIdx >= 0 {
		z.Players[foundIdx].Zone = nil
		z.Players = append(z.Players[:foundIdx], z.Players[foundIdx+1:]...)

		return nil
	} else {
		return fmt.Errorf("Player [%+v] is not in zone [%+v]", player, z)
	}
}
