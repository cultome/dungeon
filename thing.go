package dungeon

type Thing interface {
	GetZone() *Zone
	SetZone(*Zone)
}
