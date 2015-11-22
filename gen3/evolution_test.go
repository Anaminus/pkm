package gen3_test

import (
	"github.com/anaminus/pkm/gen3"
	"testing"
)

func TestEvolution(t *testing.T) {
	ver := gen3.OpenROM(ROM(t))
	if ver == nil {
		t.Fatalf("failed to open ROM")
	}

	evo := ver.SpeciesByIndex(1).Evolutions()[0]
	if v := evo.Target(); v != ver.SpeciesByIndex(2) {
		if v == nil {
			t.Errorf("Target: unexpected result <nil>")
		} else {
			t.Errorf("Target: unexpected result %d (%s)", v.Index(), v.Name())
		}
	}
	if v := evo.Method(); v != 4 {
		t.Errorf("Method: unexpected result %d", v)
	}
	if v := evo.Param(); v != 16 {
		t.Errorf("Param: unexpected result %d", v)
	}

	type ev struct {
		species string
		evo     int
		method  string
	}
	expected := []ev{
		{"CHANSEY", 0, "Friendship"},
		{"EEVEE", 3, "Friendship (Day)"},
		{"EEVEE", 4, "Friendship (Night)"},
		{"BULBASAUR", 0, "Level 16"},
		{"HAUNTER", 0, "Trade"},
		{"ONIX", 0, "Trade holding METAL COAT"},
		{"EEVEE", 0, "Use THUNDERSTONE"},
		{"TYROGUE", 1, "Level 20 if ATK > DEF"},
		{"TYROGUE", 2, "Level 20 if ATK = DEF"},
		{"TYROGUE", 0, "Level 20 if ATK < DEF"},
		{"WURMPLE", 0, "Personality[1] (7)"},
		{"WURMPLE", 1, "Personality[2] (7)"},
		{"NINCADA", 0, "Level 20 (Spawns extra)"},
		{"NINCADA", 1, "Level 20 (Spawned)"},
		{"FEEBAS", 0, "Beauty (170)"},
	}
	for _, e := range expected {
		species := ver.SpeciesByName(e.species)
		evo := species.Evolutions()[e.evo]
		if evo.MethodString() != e.method {
			t.Errorf("MethodString: unexpected result: %s (%d) -> \"%s\"", e.species, e.evo, evo.MethodString())
		}
	}
}
