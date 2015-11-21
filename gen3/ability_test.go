package gen3_test

import (
	"github.com/anaminus/pkm/gen3"
	"testing"
)

func TestAbility(t *testing.T) {
	ver := gen3.OpenROM(ROM(t))
	if ver == nil {
		t.Fatalf("failed to open ROM")
	}

	ability := ver.AbilityByIndex(1)
	if v := ability.Index(); v != 1 {
		t.Errorf("Index: unexpected index %d", v)
	}
	if v := ability.Name(); v != "STENCH" {
		t.Errorf("Name: unexpected name %s", v)
	}
	if v := ability.Description(); v != "Helps repel wild POKéMON." {
		t.Errorf("Description: unexpected description %s", v)
	}
}
