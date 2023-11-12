package dungeon

import "testing"

func TestBuildZone(t *testing.T) {
	zone := BuildZone()

	if zone == nil {
		t.Fatal("Zone not builded!")
	}
}
