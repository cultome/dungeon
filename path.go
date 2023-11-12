package dungeon

type Path struct {
	ID string
}

func BuildPath(id string) *Path {
	return &Path{ID: id}
}
